package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/anomalyco/task-management-api/internal/application"
	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/application/ports"
)

type contextKey string

const (
	authInfoKey contextKey = "auth_info"
)

type AuthInfo struct {
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	SessionID string `json:"session_id"`
}

func SetAuthInfo(ctx context.Context, claims *ports.TokenClaims) context.Context {
	info := &AuthInfo{
		UserID:    claims.UserID,
		Role:      string(claims.Role),
		SessionID: claims.SessionID,
	}
	return context.WithValue(ctx, authInfoKey, info)
}

func GetAuthInfo(ctx context.Context) (*AuthInfo, bool) {
	info, ok := ctx.Value(authInfoKey).(*AuthInfo)
	return info, ok
}

type AuthHandler struct {
	authUC *application.AuthUseCase
}

func NewAuthHandler(uc *application.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUC: uc}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo de solicitud invalido"})
		return
	}

	output, err := h.authUC.Login(r.Context(), input)
	if err != nil {
		if err == application.ErrInvalidCredentials {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "credenciales invalidas"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error interno"})
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	authInfo, ok := GetAuthInfo(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "no autenticado"})
		return
	}

	if err := h.authUC.Logout(r.Context(), authInfo.SessionID); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al cerrar sesion"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "sesion cerrada"})
}

func (h *AuthHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	authInfo, ok := GetAuthInfo(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "no autenticado"})
		return
	}

	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "cuerpo de solicitud invalido"})
		return
	}

	if err := h.authUC.ChangePassword(r.Context(), authInfo.UserID, input.OldPassword, input.NewPassword); err != nil {
		if err == application.ErrUserNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "usuario no encontrado"})
			return
		}
		if err == application.ErrInvalidCredentials {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "contrasena actual incorrecta"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error al cambiar contrasena"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
