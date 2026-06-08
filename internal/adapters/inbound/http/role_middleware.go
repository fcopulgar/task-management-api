package http

import (
	"net/http"

	"github.com/anomalyco/task-management-api/internal/domain"
)

func RequireRole(role domain.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authInfo, ok := GetAuthInfo(r.Context())
			if !ok {
				writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "no autenticado"})
				return
			}

			if domain.Role(authInfo.Role) != role {
				writeJSON(w, http.StatusForbidden, map[string]string{"error": "acceso no autorizado para este perfil"})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RequireAnyRole(roles ...domain.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authInfo, ok := GetAuthInfo(r.Context())
			if !ok {
				writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "no autenticado"})
				return
			}

			userRole := domain.Role(authInfo.Role)
			for _, role := range roles {
				if userRole == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			writeJSON(w, http.StatusForbidden, map[string]string{"error": "acceso no autorizado para este perfil"})
		})
	}
}
