package http

import (
	"github.com/go-chi/chi/v5"

	"github.com/anomalyco/task-management-api/internal/application"
)

func SetupAuthRoutes(r chi.Router, deps application.Dependencies) {
	authUC := application.NewAuthUseCase(deps)
	authHandler := NewAuthHandler(authUC)

	authMW := NewAuthMiddleware(deps.UserRepo, deps.SessionRepo, deps.TokenSvc)

	r.Post("/auth/login", authHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(authMW.Authenticate)

		r.Post("/auth/logout", authHandler.Logout)
		r.Post("/auth/password", authHandler.ChangePassword)
	})
}

func SetupProtectedRoutes(r chi.Router, deps application.Dependencies) {
	authMW := NewAuthMiddleware(deps.UserRepo, deps.SessionRepo, deps.TokenSvc)

	r.Group(func(r chi.Router) {
		r.Use(authMW.Authenticate)
		r.Use(authMW.RequirePasswordNotTemporary)

		// endpoints protegidos se registraran en fases futuras
	})
}
