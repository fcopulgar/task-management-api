package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/anomalyco/task-management-api/internal/application"
	"github.com/anomalyco/task-management-api/internal/application/dto"
)

type TaskHandler struct {
	taskUC *application.TaskUseCase
}

func NewTaskHandler(uc *application.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUC: uc}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	authInfo, _ := GetAuthInfo(r.Context())

	var input dto.CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo de solicitud invalido"})
		return
	}

	output, err := h.taskUC.CreateTask(r.Context(), input, authInfo.UserID)
	if err != nil {
		handleTaskError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskUC.ListTasks(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al listar tareas"})
		return
	}

	if tasks == nil {
		tasks = []dto.TaskOutput{}
	}

	writeJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, err := h.taskUC.GetTask(r.Context(), id)
	if err != nil {
		if err == application.ErrTaskNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "tarea no encontrada"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al obtener tarea"})
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var input dto.UpdateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo de solicitud invalido"})
		return
	}
	input.ID = id

	output, err := h.taskUC.UpdateTask(r.Context(), input)
	if err != nil {
		handleTaskError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.taskUC.DeleteTask(r.Context(), id); err != nil {
		handleTaskError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleTaskError(w http.ResponseWriter, err error) {
	switch err {
	case application.ErrTaskNotFound:
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "tarea no encontrada"})
	case application.ErrTaskNotModifiable:
		writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
	default:
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error interno"})
	}
}
