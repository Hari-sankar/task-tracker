package controllers

import (
	"net/http"
	"task-tracker/api/models"
	"task-tracker/api/schemas"
	"task-tracker/api/usecases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
	taskUseCase usecases.TaskUseCase
}

func NewTaskController(taskUseCase usecases.TaskUseCase) *TaskController {
	return &TaskController{taskUseCase: taskUseCase}
}

// Create task godoc
// @Summary Create a new task
// @Description Create a new task
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body schemas.CreateTaskRequest true "Create Task request"
// @Success 200 {object} schemas.CreateTaskResponse
// @Failure 401 {object} utils.ErrorStruct "Invalid credentials"
// @Failure 500 {object} utils.ErrorStruct "Internal server error"
// @Router /task [post]
func (ctrl *TaskController) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()

	var req schemas.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(schemas.ErrInvalidInput)
		return
	}
	userID := c.GetString("userID")
	var task models.Task
	task.Title = req.Title
	task.Description = req.Description
	task.UserID = uuid.MustParse(userID)

	response, err := ctrl.taskUseCase.CreateTask(ctx, &task)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, &schemas.CreateTaskResponse{ID: response})

}

// Get a tasks godoc
// @Summary Get a task
// @Description Get a task by id
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param taskID path string true "taskID" Example(123e4567-e89b-12d3-a456-426614174000)
// @Success 200 {object} schemas.GetTaskByIDResponse "Tasks retrieved successfully"
// @Failure 500 {object} utils.ErrorStruct "Internal server error"
// @Router /task/{taskID} [get]
func (ctrl *TaskController) GetTaskByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		c.Error(schemas.ErrInvalidInput)
		return
	}

	task, err := ctrl.taskUseCase.GetTaskByID(ctx, id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, task)
}

// Update task godoc
// @Summary Update task
// @Description Update task
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body schemas.UpdateTaskRequest true "Update Task request"
// @Success 200 {object} schemas.UpdateTaskResponse "Task updated successfully"
// @Failure 400 {object} utils.ErrorStruct "Invalid input"
// @Failure 500 {object} utils.ErrorStruct "Internal server error"
// @Router /task [patch]
func (ctrl *TaskController) UpdateTask(c *gin.Context) {
	// Bind the request body to the UpdateTaskRequest struct
	ctx := c.Request.Context()
	var req schemas.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(schemas.ErrInvalidInput)
		return
	}

	var task models.Task
	task.ID = req.TaskID
	task.Title = req.Title
	task.Description = req.Description
	task.Status = req.Status
	userID := uuid.MustParse(c.GetString("userID"))

	updatedTask, err := ctrl.taskUseCase.UpdateTask(ctx, &task, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, schemas.UpdateTaskResponse{UpdatedTask: *updatedTask})
}

// Delete task godoc
// @Summary Delete task
// @Description Delete task
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param taskID path string true "Task ID" Example(123e4567-e89b-12d3-a456-426614174000)
// @Success 200 {object} schemas.DeleteTaskResponse "Task deleted successfully"
// @Failure 400 {object} utils.ErrorStruct "Invalid input"
// @Failure 500 {object} utils.ErrorStruct "Internal server error"
// @Router /task/{taskID} [delete]
func (ctrl *TaskController) DeleteTask(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := uuid.Parse(c.GetString("userID"))
	if err != nil {
		c.Error(schemas.ErrInvalidInput)
		return
	}

	task_id, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		c.Error(schemas.ErrInvalidInput)
		return
	}

	err = ctrl.taskUseCase.DeleteTask(ctx, task_id, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// Get all tasks godoc
// @Summary Get all tasks
// @Description Get all tasks
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} schemas.GetAllTasksResponse "Tasks retrieved successfully"
// @Failure 500 {object} utils.ErrorStruct "Internal server error"
// @Router /task [get]
func (ctrl *TaskController) GetAllTasks(c *gin.Context) {
	ctx := c.Request.Context()
	userID := uuid.MustParse(c.GetString("userID"))
	tasks, err := ctrl.taskUseCase.GetAllTasksByUserID(ctx, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}
