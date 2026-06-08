package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/anomalyco/task-management-api/internal/application"
	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/domain"
)

type ExecutorHandler struct {
	execUC *application.ExecutorUseCase
}

func NewExecutorHandler(uc *application.ExecutorUseCase) *ExecutorHandler {
	return &ExecutorHandler{execUC: uc}
}

func (h *ExecutorHandler) ListMyTasks(w http.ResponseWriter, r *http.Request) {
	authInfo, _ := GetAuthInfo(r.Context())

	tasks, err := h.execUC.ListMyTasks(r.Context(), authInfo.UserID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al listar tareas"})
		return
	}

	if tasks == nil {
		tasks = []dto.TaskOutput{}
	}

	writeJSON(w, http.StatusOK, tasks)
}

func (h *ExecutorHandler) GetMyTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authInfo, _ := GetAuthInfo(r.Context())

	task, err := h.execUC.GetMyTask(r.Context(), id, authInfo.UserID)
	if err != nil {
		if err == application.ErrTaskNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "tarea no encontrada"})
			return
		}
		writeJSON(w, http.StatusForbidden, map[string]string{"error": "acceso denegado"})
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *ExecutorHandler) TransitionStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authInfo, _ := GetAuthInfo(r.Context())

	var input dto.TransitionTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo de solicitud invalido"})
		return
	}

	output, err := h.execUC.TransitionTask(r.Context(), id, authInfo.UserID, input.NewStatus)
	if err != nil {
		handleExecutorError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *ExecutorHandler) CommentOnTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authInfo, _ := GetAuthInfo(r.Context())

	var input dto.CreateCommentInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo de solicitud invalido"})
		return
	}

	output, err := h.execUC.CommentOnTask(r.Context(), id, authInfo.UserID, input.Comment)
	if err != nil {
		handleExecutorError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func handleExecutorError(w http.ResponseWriter, err error) {
	switch {
	case err == application.ErrTaskNotFound:
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "tarea no encontrada"})
	case err == application.ErrTaskOverdue:
		writeJSON(w, http.StatusConflict, map[string]string{"error": "la tarea esta vencida"})
	case err == domain.ErrNotTaskOwner:
		writeJSON(w, http.StatusForbidden, map[string]string{"error": "la tarea no pertenece al usuario"})
	case err == domain.ErrInvalidTransition:
		writeJSON(w, http.StatusConflict, map[string]string{"error": "transicion de estado no permitida"})
	case err == domain.ErrTaskTerminal:
		writeJSON(w, http.StatusConflict, map[string]string{"error": "la tarea ya esta en estado terminal"})
	case err == domain.ErrCommentOnNonOverdue:
		writeJSON(w, http.StatusConflict, map[string]string{"error": "solo se puede comentar tareas vencidas"})
	default:
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error interno"})
	}
}
