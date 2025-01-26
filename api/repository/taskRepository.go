package repository

import (
	"context"
	"fmt"
	"strings"
	"task-tracker/api/models"
	"task-tracker/api/utils"
	"task-tracker/pkg/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// TaskRepository is the interface for the task repository
type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Task) (string, error)
	GetTaskByID(ctx context.Context, id uuid.UUID) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task, userID uuid.UUID) (*models.Task, error)
	DeleteTask(ctx context.Context, taskID uuid.UUID, userID uuid.UUID) error
	GetAllTasksByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Task, error)
}

type taskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) TaskRepository {
	return &taskRepository{db: db}
}

// CreateTask creates a new task in the database
func (r *taskRepository) CreateTask(ctx context.Context, task *models.Task) (string, error) {
	//Query to insert a new task into the database
	query := `INSERT INTO tasks (title, description, user_id)
			  VALUES (@Title, NULLIF(@Description,NULL), @UserID)
			  RETURNING id`

	//Create a map of arguments to be passed to the query
	args := pgx.NamedArgs{
		"Title":       task.Title,
		"Description": task.Description,
		"UserID":      task.UserID,
	}

	//Execute the query and scan the result into the id variable
	var id string
	err := r.db.QueryRow(ctx, query, args).Scan(&id)

	if err != nil {
		logger.Error("Failed to create task",
			zap.Error(err),
			zap.String("user_id", task.UserID.String()),
		)
		return "", utils.HandleDBError(err)
	}
	return id, nil

}

// GetTaskByID retrieves a task by its ID from the database
func (r *taskRepository) GetTaskByID(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	//Query to retrieve a task by its ID
	query := `SELECT id, title, description, status, user_id
			  FROM tasks
			  WHERE id = @ID`
	args := pgx.NamedArgs{
		"ID": id,
	}

	//Execute the query and scan the result into the task variable
	var task models.Task
	err := r.db.QueryRow(ctx, query, args).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.UserID)

	if err != nil {
		logger.Error("Failed to fetch task",
			zap.Error(err),
			zap.String("task_id", id.String()),
		)
		return nil, utils.HandleDBError(err)
	}
	return &task, nil
}

// UpdateTask updates an existing task in the database
func (r *taskRepository) UpdateTask(ctx context.Context, task *models.Task, userID uuid.UUID) (*models.Task, error) {
	//Check the required fields to update
	var updates []string
	args := pgx.NamedArgs{
		"ID":     task.ID,
		"UserID": userID,
	}

	if task.Title != "" {
		updates = append(updates, "title = @Title")
		args["Title"] = task.Title
	}
	if task.Description != "" {
		updates = append(updates, "description = @Description")
		args["Description"] = task.Description
	}
	if task.Status != "" {
		updates = append(updates, "status = @Status")
		args["Status"] = task.Status
	}

	if len(updates) == 0 {
		return nil, utils.NewErrorStruct(400, "No fields to update") // No fields to update
	}

	//Query to update the task
	query := fmt.Sprintf(`
		UPDATE tasks 
		SET %s , updated_at = NOW()
		WHERE id = @ID AND user_id = @UserID
		RETURNING  id, title, description, status, user_id, created_at, updated_at`, strings.Join(updates, ", "))

	var updatedTask models.Task
	err := r.db.QueryRow(ctx, query, args).Scan(
		&updatedTask.ID,
		&updatedTask.Title,
		&updatedTask.Description,
		&updatedTask.Status,
		&updatedTask.UserID,
		&updatedTask.CreatedAt,
		&updatedTask.UpdatedAt)
	if err != nil {
		logger.Error("Failed to update task",
			zap.Error(err),
			zap.String("task_id", task.ID.String()),
			zap.String("user_id", userID.String()),
		)
		return nil, utils.HandleDBError(err)
	}

	return &updatedTask, nil
}

// DeleteTask deletes a task from the database
func (r *taskRepository) DeleteTask(ctx context.Context, taskID uuid.UUID, userID uuid.UUID) error {
	//Query to delete a task from the database
	query := `DELETE FROM tasks WHERE id = @ID AND user_id = @UserID Returning true`
	args := pgx.NamedArgs{
		"ID":     taskID,
		"UserID": userID,
	}

	//Execute the query and scan the result into the isDeleted variable
	isDeleted := false
	err := r.db.QueryRow(ctx, query, args).Scan(&isDeleted)

	if !isDeleted {
		logger.Error("Failed to delete task",
			zap.Error(err),
			zap.String("task_id", taskID.String()),
			zap.String("user_id", userID.String()),
		)
		return utils.NewErrorStruct(404, "Task not found")
	}

	if err != nil {
		logger.Error("Failed to delete task",
			zap.Error(err),
			zap.String("task_id", taskID.String()),
			zap.String("user_id", userID.String()),
		)
		return utils.HandleDBError(err)
	}
	return nil

}

// GetAllTasksByUserID retrieves all tasks for a user from the database
func (r *taskRepository) GetAllTasksByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Task, error) {
	//Query to retrieve all tasks for a user
	query := `SELECT id, title, description, status, user_id, created_at, updated_at
			  FROM tasks
			  WHERE user_id = @UserID`
	args := pgx.NamedArgs{
		"UserID": userID,
	}

	//Execute the query and scan the results into the tasks variable
	rows, err := r.db.Query(ctx, query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan the results into the tasks variable
	var tasks []*models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.UserID,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			logger.Error("Failed to fetch user tasks",
				zap.Error(err),
				zap.String("user_id", userID.String()),
			)
			return nil, utils.HandleDBError(err)
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
