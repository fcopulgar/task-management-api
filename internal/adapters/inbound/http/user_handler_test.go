package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/anomalyco/task-management-api/internal/application"
	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/application/ports"
	"github.com/anomalyco/task-management-api/internal/domain"
)

func setupUserTestRouter(deps application.Dependencies) chi.Router {
	r := chi.NewRouter()
	SetupUserRoutes(r, deps)
	return r
}

func newAdminUser() *domain.User {
	u, _ := domain.NewUser("Admin", "admin@test.com", "hash", domain.RoleAdmin)
	u.ID = "admin-1"
	u.MarkPasswordChanged()
	return u
}

func newAdminUserWithID(id string) *domain.User {
	u, _ := domain.NewUser("Admin", "admin@test.com", "hash", domain.RoleAdmin)
	u.ID = id
	u.MarkPasswordChanged()
	return u
}

func TestCreateUserHandlerSuccess(t *testing.T) {
	userRepo := &mockUserRepo{
		findByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return nil, nil
		},
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			u := newAdminUser()
			
			return u, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc, Hasher: &mockHasher{},
	}
	router := setupUserTestRouter(deps)

	body := `{"name":"Executor","email":"exec@test.com","password":"pwd","role":"EXECUTOR"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateUserHandlerCannotCreateAdmin(t *testing.T) {
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			u := newAdminUser()
			
			return u, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc, Hasher: &mockHasher{},
	}
	router := setupUserTestRouter(deps)

	body := `{"name":"Admin2","email":"admin2@test.com","password":"pwd","role":"ADMIN"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestListUsersHandlerSuccess(t *testing.T) {
	u1, _ := domain.NewUser("User 1", "u1@test.com", "hash", domain.RoleExecutor)
	u1.ID = "u1"
	u2, _ := domain.NewUser("User 2", "u2@test.com", "hash", domain.RoleAuditor)
	u2.ID = "u2"

	userRepo := &mockUserRepo{
		listFn: func(ctx context.Context) ([]domain.User, error) {
			return []domain.User{*u1, *u2}, nil
		},
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			u := newAdminUser()
			
			return u, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc,
	}
	router := setupUserTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []dto.UserOutput
	json.NewDecoder(w.Body).Decode(&users)
	assert.Len(t, users, 2)
	assert.Equal(t, "User 1", users[0].Name)
}

func TestGetUserHandlerNotFound(t *testing.T) {
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			if id == "admin-1" {
				u := newAdminUser()
				
				return u, nil
			}
			return nil, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc,
	}
	router := setupUserTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/users/nonexistent", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetUserHandlerSuccess(t *testing.T) {
	u, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	u.ID = "user-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			if id == "user-1" {
				return u, nil
			}
			admin := newAdminUserWithID(id)
			
			return admin, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc,
	}
	router := setupUserTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/users/user-1", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var output dto.UserOutput
	json.NewDecoder(w.Body).Decode(&output)
	assert.Equal(t, "test@test.com", output.Email)
}

func TestUpdateUserHandlerSuccess(t *testing.T) {
	u, _ := domain.NewUser("Original", "orig@test.com", "hash", domain.RoleExecutor)
	u.ID = "user-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			if id == "user-1" {
				return u, nil
			}
			admin := newAdminUserWithID(id)
			
			return admin, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc,
	}
	router := setupUserTestRouter(deps)

	body := `{"name":"Actualizado"}`
	req := httptest.NewRequest(http.MethodPut, "/users/user-1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var output dto.UserOutput
	json.NewDecoder(w.Body).Decode(&output)
	assert.Equal(t, "Actualizado", output.Name)
}

func TestDeleteUserHandlerSuccess(t *testing.T) {
	u, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	u.ID = "user-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			if id == "user-1" {
				return u, nil
			}
			admin := newAdminUserWithID(id)
			
			return admin, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc,
	}
	router := setupUserTestRouter(deps)

	req := httptest.NewRequest(http.MethodDelete, "/users/user-1", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.False(t, u.IsActive())
}

func TestUserRoutesRequireAdminRole(t *testing.T) {
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("exec-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "exec-1", Role: domain.RoleExecutor, SessionID: "s1"}, nil
		},
	}
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			u, _ := domain.NewUser("Exec", "exec@test.com", "hash", domain.RoleExecutor)
			u.ID = "exec-1"
			return u, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc,
	}
	router := setupUserTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestUserRoutesNoAuth(t *testing.T) {
	deps := application.Dependencies{}
	router := setupUserTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
