package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/internal/database/repositories"
	"rest-api/internal/models"
	"strings"

	"github.com/google/uuid"
)

type TaskHandler struct {
	repo *repositories.TaskRepo
}

func NewTaskHandler(repo *repositories.TaskRepo) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJson(w, statusCode, map[string]string{"error": message})
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	tasks, err := h.repo.GetAll(req.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve tasks")
		return
	}

	respondWithJson(w, http.StatusOK, tasks)
}

func (h *TaskHandler) GetById(w http.ResponseWriter, req *http.Request) {
	pathParts := strings.Split(strings.TrimPrefix(req.URL.RawPath, "/tasks/"), "/")
	idStr := pathParts[0]

	id, err := uuid.Parse(idStr)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := h.repo.GetById(req.Context(), id)

	if task == nil && err == nil {
		respondWithError(w, http.StatusNotFound, "Task was not found")
		return
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve task")
		return
	}

	respondWithJson(w, http.StatusOK, task)
}

func (h *TaskHandler) Create(w http.ResponseWriter, req *http.Request) {
	var input models.CreateTask

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error occured while decoding body")
		return
	}

	if strings.TrimSpace(input.Description) == "" {
		respondWithError(w, http.StatusBadRequest, "Description can not be empty")
		return
	}

	if len(input.Description) > 200 {
		respondWithError(w, http.StatusBadRequest, "Description length can not be longer than 200 symbols")
		return
	}

	task, err := h.repo.Create(req.Context(), input)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error occured while creating task")
		return
	}

	respondWithJson(w, http.StatusOK, task)
}

func (h *TaskHandler) Update(w http.ResponseWriter, req *http.Request) {
	pathParts := strings.Split(strings.TrimPrefix(req.URL.RawPath, "/tasks/"), "/")
	idStr := pathParts[0]

	id, err := uuid.Parse(idStr)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	var input models.UpdateTask

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error occured while decoding body")
		return
	}

	task, err := h.repo.Update(req.Context(), id, input)

	if task == nil && err != nil {
		respondWithError(w, http.StatusNotFound, "Task was not found")
		return
	}

	if err == nil {
		respondWithError(w, http.StatusInternalServerError, "Error occured while updating task")
		return
	}

	respondWithJson(w, http.StatusOK, task)
}
