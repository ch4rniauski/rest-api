package handlers

import (
	"net/http"
	"rest-api/internal/database/repositories"
	"rest-api/internal/models"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type TaskHandler struct {
	repo *repositories.TaskRepo
}

func NewTaskHandler(repo *repositories.TaskRepo) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func RegisterRoutes(e *echo.Echo, repo *repositories.TaskRepo) {
	h := NewTaskHandler(repo)

	e.GET("/tasks", h.GetAll)
	e.POST("/tasks", h.Create)

	e.GET("/tasks/:id", h.GetById)
	e.PUT("/tasks/:id", h.Update)
	e.PATCH("/tasks/:id", h.Update)
}

func (h *TaskHandler) GetAll(c *echo.Context) error {
	tasks, err := h.repo.GetAll(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve tasks",
		})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetById(c *echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid task ID",
		})
	}

	task, err := h.repo.GetById(c.Request().Context(), id)

	if task == nil && err == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Task was not found",
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve task",
		})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Create(c *echo.Context) error {
	var input models.CreateTask

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if strings.TrimSpace(input.Description) == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Description can not be empty",
		})
	}

	if len(input.Description) > 200 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Description length can not be longer than 200 symbols",
		})
	}

	task, err := h.repo.Create(c.Request().Context(), input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error occured while creating task",
		})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) Update(c *echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid task ID",
		})
	}

	var input models.UpdateTask

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	task, err := h.repo.Update(c.Request().Context(), id, input)

	if task == nil && err == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Task was not found",
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error occured while updating task",
		})
	}

	return c.JSON(http.StatusOK, task)
}
