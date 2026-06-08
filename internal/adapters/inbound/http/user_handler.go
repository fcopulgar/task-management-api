package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/anomalyco/task-management-api/internal/application"
	"github.com/anomalyco/task-management-api/internal/application/dto"
)

type UserHandler struct {
	userUC *application.UserUseCase
}

func NewUserHandler(uc *application.UserUseCase) *UserHandler {
	return &UserHandler{userUC: uc}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo de solicitud invalido"})
		return
	}

	output, err := h.userUC.CreateUser(r.Context(), input)
	if err != nil {
		switch err {
		case application.ErrCannotCreateAdmin:
			writeJSON(w, http.StatusForbidden, map[string]string{"error": err.Error()})
		case application.ErrEmailAlreadyExists:
			writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		default:
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al crear usuario"})
		}
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUC.ListUsers(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al listar usuarios"})
		return
	}

	if users == nil {
		users = []dto.UserOutput{}
	}

	writeJSON(w, http.StatusOK, users)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.userUC.GetUser(r.Context(), id)
	if err != nil {
		if err == application.ErrUserNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "usuario no encontrado"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al obtener usuario"})
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var input dto.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo de solicitud invalido"})
		return
	}
	input.ID = id

	output, err := h.userUC.UpdateUser(r.Context(), input)
	if err != nil {
		switch err {
		case application.ErrUserNotFound:
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "usuario no encontrado"})
		case application.ErrEmailAlreadyExists:
			writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		default:
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al actualizar usuario"})
		}
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.userUC.DeactivateUser(r.Context(), id); err != nil {
		if err == application.ErrUserNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "usuario no encontrado"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al desactivar usuario"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
