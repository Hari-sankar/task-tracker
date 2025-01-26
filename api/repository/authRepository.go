package repository

import (
	"context"
	"fmt"
	"task-tracker/api/models"
	"task-tracker/api/schemas"
	"task-tracker/api/utils"
	"task-tracker/pkg/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (uuid.UUID, error)
}

type authRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(ctx context.Context, user *models.User) (uuid.UUID, error) {
	query := `INSERT INTO users (username, password)
			  VALUES (@Username, @Password)
			  RETURNING id`

	args := pgx.NamedArgs{
		"Username": user.Username,
		"Password": user.Password,
	}

	var id string
	err := r.db.QueryRow(ctx, query, args).Scan(&id)
	if err != nil {
		logger.Error("Failed to Create user",
			zap.Error(err),
			zap.Any("user", user),
		)

		// Check for unique constraint violation
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" { // PostgreSQL unique violation code
				return uuid.Nil, utils.NewErrorStruct(400, fmt.Sprintf("username %s already exists", user.Username))
			}
		}
		return uuid.Nil, utils.NewErrorStruct(500, fmt.Sprintf("unable to create user: %v", err))
	}

	return uuid.MustParse(id), nil
}

func (r *authRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT id, username, password FROM users WHERE username = $1`

	var user models.User
	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		logger.Error("Failed to fetch user",
			zap.Error(err),
			zap.String("username", username),
		)
		if err == pgx.ErrNoRows {
			return nil, schemas.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}
