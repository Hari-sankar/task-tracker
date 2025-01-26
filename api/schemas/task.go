package schemas

import (
	"task-tracker/api/models"

	"github.com/google/uuid"
)

// CreateTaskRequest is the request body for creating a new task
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
}

// CreateTaskResponse is the response body for creating a new task
type CreateTaskResponse struct {
	ID string `json:"id"`
}

// TaskResponse is the response body for a task
type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
	UserID      string `json:"user_id"`
}

// GetAllTasksResponse is the response body for getting all tasks
type GetAllTasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

// UpdateTaskRequest is the request body for updating a task
type UpdateTaskRequest struct {
	TaskID      uuid.UUID         `json:"id" binding:"required"`
	Title       string            `json:"title" binding:"omitempty"`
	Description string            `json:"description" binding:"omitempty"`
	Status      models.TaskStatus `json:"status" binding:"omitempty,oneof=PENDING IN_PROGRESS COMPLETED"`
}

// UpdateTaskResponse is the response body for updating a task
type UpdateTaskResponse struct {
	UpdatedTask models.Task `json:"task"`
}

// GetTaskByIDResponse is the response body for getting a task by ID
type GetTaskByIDResponse struct {
	Task TaskResponse `json:"task"`
}

// DeleteTaskResponse is the response body for deleting a task
type DeleteTaskResponse struct {
	ID string `json:"id"`
}
