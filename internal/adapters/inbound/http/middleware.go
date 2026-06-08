package http

import (
	"net/http"
	"strings"
	"time"

	"github.com/anomalyco/task-management-api/internal/application"
	"github.com/anomalyco/task-management-api/internal/application/ports"
)

type AuthMiddleware struct {
	userRepo    ports.UserRepository
	sessionRepo ports.SessionRepository
	tokenSvc    ports.TokenService
}

func NewAuthMiddleware(
	userRepo ports.UserRepository,
	sessionRepo ports.SessionRepository,
	tokenSvc ports.TokenService,
) *AuthMiddleware {
	return &AuthMiddleware{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		tokenSvc:    tokenSvc,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "token no proporcionado"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "formato de token invalido"})
			return
		}

		tokenString := parts[1]

		claims, err := m.tokenSvc.Validate(r.Context(), tokenString)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "token invalido"})
			return
		}

		session, err := m.sessionRepo.FindByID(r.Context(), claims.SessionID)
		if err != nil || session == nil {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "sesion no encontrada"})
			return
		}

		if !session.IsValid(time.Now()) {
			if session.IsRevoked() {
				writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "sesion revocada"})
				return
			}
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "sesion expirada"})
			return
		}

		user, err := m.userRepo.FindByID(r.Context(), claims.UserID)
		if err != nil || user == nil {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "usuario no encontrado"})
			return
		}

		if !user.IsActive() {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "usuario inactivo"})
			return
		}

		ctx := SetAuthInfo(r.Context(), claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) RequirePasswordNotTemporary(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authInfo, ok := GetAuthInfo(r.Context())
		if !ok {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "no autenticado"})
			return
		}

		user, err := m.userRepo.FindByID(r.Context(), authInfo.UserID)
		if err != nil || user == nil {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "usuario no encontrado"})
			return
		}

		if user.MustChangePassword {
			if !isPasswordChangeRoute(r.URL.Path) {
				writeJSON(w, http.StatusForbidden, map[string]string{"error": "debe cambiar su contrasena antes de continuar"})
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func isPasswordChangeRoute(path string) bool {
	return path == "/auth/password" || path == "/auth/logout" || path == "/health"
}

func RequireAuth(authMW *AuthMiddleware) func(http.Handler) http.Handler {
	return authMW.Authenticate
}

func (m *AuthMiddleware) ChangePasswordError(w http.ResponseWriter, r *http.Request, err error) {
	if err == application.ErrInvalidCredentials {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "credenciales invalidas"})
		return
	}
	writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "error interno"})
}
