package usecases

import (
	"context"
	"task-tracker/api/models"
	"task-tracker/api/repository"

	"github.com/google/uuid"
)

// TaskUseCase defines the interface for the task use case
type TaskUseCase interface {
	CreateTask(ctx context.Context, task *models.Task) (string, error)
	GetTaskByID(ctx context.Context, id uuid.UUID) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task, userID uuid.UUID) (*models.Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
	GetAllTasksByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Task, error)
}

// taskUseCase implements the TaskUseCase interface
type taskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUseCase(taskRepo repository.TaskRepository) TaskUseCase {
	return &taskUseCase{taskRepo: taskRepo}
}

// CreateTask creates a new task
func (uc *taskUseCase) CreateTask(ctx context.Context, task *models.Task) (string, error) {
	return uc.taskRepo.CreateTask(ctx, task)
}

// GetTaskByID retrieves a task by its ID
func (uc *taskUseCase) GetTaskByID(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	return uc.taskRepo.GetTaskByID(ctx, id)
}

// UpdateTask updates an existing task
func (uc *taskUseCase) UpdateTask(ctx context.Context, task *models.Task, userID uuid.UUID) (*models.Task, error) {
	return uc.taskRepo.UpdateTask(ctx, task, userID)
}

// DeleteTask deletes a task by its ID
func (uc *taskUseCase) DeleteTask(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	return uc.taskRepo.DeleteTask(ctx, id, userID)
}

// GetAllTasksByUserID retrieves all tasks for a user
func (uc *taskUseCase) GetAllTasksByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Task, error) {
	return uc.taskRepo.GetAllTasksByUserID(ctx, userID)
}
