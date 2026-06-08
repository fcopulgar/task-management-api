package http

import (
	"github.com/go-chi/chi/v5"

	"github.com/anomalyco/task-management-api/internal/application"
	"github.com/anomalyco/task-management-api/internal/domain"
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

func SetupUserRoutes(r chi.Router, deps application.Dependencies) {
	userUC := application.NewUserUseCase(deps)
	userHandler := NewUserHandler(userUC)

	authMW := NewAuthMiddleware(deps.UserRepo, deps.SessionRepo, deps.TokenSvc)

	r.Route("/users", func(r chi.Router) {
		r.Use(authMW.Authenticate)
		r.Use(RequireRole(domain.RoleAdmin))
		r.Use(authMW.RequirePasswordNotTemporary)

		r.Post("/", userHandler.Create)
		r.Get("/", userHandler.List)
		r.Get("/{id}", userHandler.Get)
		r.Put("/{id}", userHandler.Update)
		r.Delete("/{id}", userHandler.Delete)
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

func SetupTaskRoutes(r chi.Router, deps application.Dependencies) {
	taskUC := application.NewTaskUseCase(deps)
	taskHandler := NewTaskHandler(taskUC)

	authMW := NewAuthMiddleware(deps.UserRepo, deps.SessionRepo, deps.TokenSvc)

	r.Route("/tasks", func(r chi.Router) {
		r.Use(authMW.Authenticate)
		r.Use(authMW.RequirePasswordNotTemporary)

		r.With(RequireAnyRole(domain.RoleAdmin, domain.RoleAuditor)).Get("/", taskHandler.List)
		r.With(RequireAnyRole(domain.RoleAdmin, domain.RoleAuditor)).Get("/{id}", taskHandler.Get)

		r.With(RequireRole(domain.RoleAdmin)).Post("/", taskHandler.Create)
		r.With(RequireRole(domain.RoleAdmin)).Put("/{id}", taskHandler.Update)
		r.With(RequireRole(domain.RoleAdmin)).Delete("/{id}", taskHandler.Delete)
	})
}
