package usecases

import (
	"context"
	"task-tracker/api/models"
	"task-tracker/api/repository"
	"task-tracker/api/schemas"
	"task-tracker/api/utils"
	"time"

	"github.com/google/uuid"
)

type AuthUseCase struct {
	authRepo repository.AuthRepository
}

func NewAuthUseCase(authRepo repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{authRepo: authRepo}
}

// SignUp creates a new user
func (auth *AuthUseCase) SignUp(ctx context.Context, req *schemas.SignUpRequest) (*schemas.SignUpResponse, error) {

	var user models.User
	var userID uuid.UUID
	var err error

	user.Username = req.Username

	// Hash the password
	if user.Password, err = utils.HashPassword(req.Password); err != nil {
		return nil, err
	}

	if userID, err = auth.authRepo.CreateUser(ctx, &user); err != nil {
		return nil, err
	}
	return &schemas.SignUpResponse{
		UserID: userID,
	}, nil

}

// GetUserByID retrieves a user by their ID
func (auth *AuthUseCase) SignIn(ctx context.Context, username, password string) (*schemas.SignInResponse, error) {

	// Retrieve the user from the database
	user, err := auth.authRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, schemas.ErrInvalidCredentials
	}

	// Verify the password
	if err := utils.VerifyPassword(password, user.Password); err != nil {
		return nil, schemas.ErrInvalidCredentials
	}

	// Generate a JWT token
	claims := &utils.Claims{
		UserID:    user.ID.String(),
		UserName:  user.Username,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	token, err := utils.GenerateAcessToken(claims)
	if err != nil {
		return nil, schemas.ErrInternalServer
	}

	return &schemas.SignInResponse{
		Token: token,
	}, nil
}
