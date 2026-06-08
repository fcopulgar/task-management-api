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

type mockTaskRepo struct {
	createFn         func(ctx context.Context, task *domain.Task) error
	findByIDFn       func(ctx context.Context, id string) (*domain.Task, error)
	listFn           func(ctx context.Context) ([]domain.Task, error)
	listByAssigneeFn func(ctx context.Context, assigneeID string) ([]domain.Task, error)
	updateFn         func(ctx context.Context, task *domain.Task) error
	deleteFn         func(ctx context.Context, id string) error
}

func (m *mockTaskRepo) Create(ctx context.Context, t *domain.Task) error {
	if m.createFn != nil {
		return m.createFn(ctx, t)
	}
	return nil
}
func (m *mockTaskRepo) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}
	return nil, nil
}
func (m *mockTaskRepo) List(ctx context.Context) ([]domain.Task, error) {
	if m.listFn != nil {
		return m.listFn(ctx)
	}
	return nil, nil
}
func (m *mockTaskRepo) ListByAssignee(ctx context.Context, id string) ([]domain.Task, error) {
	if m.listByAssigneeFn != nil {
		return m.listByAssigneeFn(ctx, id)
	}
	return nil, nil
}
func (m *mockTaskRepo) Update(ctx context.Context, t *domain.Task) error {
	if m.updateFn != nil {
		return m.updateFn(ctx, t)
	}
	return nil
}
func (m *mockTaskRepo) Delete(ctx context.Context, id string) error {
	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}
	return nil
}

func setupTaskTestRouter(deps application.Dependencies) chi.Router {
	r := chi.NewRouter()
	SetupTaskRoutes(r, deps)
	return r
}

func makeAdminAuthMiddleware() *AuthMiddleware {
	admin := newAdminUser()
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
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

	return NewAuthMiddleware(userRepo, sessionRepo, nil)
}

func makeAuditorAuthMiddleware() *AuthMiddleware {
	auditor, _ := domain.NewUser("Auditor", "aud@test.com", "hash", domain.RoleAuditor)
	auditor.ID = "aud-1"
	auditor.MarkPasswordChanged()
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return auditor, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("aud-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	return NewAuthMiddleware(userRepo, sessionRepo, nil)
}

func TestCreateTaskHandlerSuccess(t *testing.T) {
	exec, _ := domain.NewUser("Exec", "exec@test.com", "hash", domain.RoleExecutor)
	exec.ID = "exec-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			if id == "exec-1" {
				return exec, nil
			}
			return newAdminUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	taskRepo := &mockTaskRepo{}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupTaskTestRouter(deps)

	body := `{"title":"Tarea Test","description":"Desc","due_at":"2026-12-31T00:00:00Z","assignee_id":"exec-1"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestListTasksHandlerAdmin(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	t1, _ := domain.NewTask("T1", "D1", dueAt, "exec-1", "admin-1")
	t1.ID = "t1"
	t2, _ := domain.NewTask("T2", "D2", dueAt, "exec-2", "admin-1")
	t2.ID = "t2"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return newAdminUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	taskRepo := &mockTaskRepo{
		listFn: func(ctx context.Context) ([]domain.Task, error) {
			return []domain.Task{*t1, *t2}, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupTaskTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var tasks []dto.TaskOutput
	json.NewDecoder(w.Body).Decode(&tasks)
	assert.Len(t, tasks, 2)
}

func TestListTasksHandlerAuditor(t *testing.T) {
	auditor, _ := domain.NewUser("Auditor", "aud@test.com", "hash", domain.RoleAuditor)
	auditor.ID = "aud-1"
	auditor.MarkPasswordChanged()

	dueAt := time.Now().Add(24 * time.Hour)
	t1, _ := domain.NewTask("T1", "D1", dueAt, "exec-1", "admin-1")
	t1.ID = "t1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			if id == "aud-1" {
				return auditor, nil
			}
			return newAdminUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("aud-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	taskRepo := &mockTaskRepo{
		listFn: func(ctx context.Context) ([]domain.Task, error) {
			return []domain.Task{*t1}, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "aud-1", Role: domain.RoleAuditor, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupTaskTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTaskAuditorForbidden(t *testing.T) {
	auditor, _ := domain.NewUser("Auditor", "aud@test.com", "hash", domain.RoleAuditor)
	auditor.ID = "aud-1"
	auditor.MarkPasswordChanged()

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return auditor, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("aud-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "aud-1", Role: domain.RoleAuditor, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc,
	}
	router := setupTaskTestRouter(deps)

	body := `{"title":"T","due_at":"2026-12-31T00:00:00Z","assignee_id":"exec-1"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestUpdateTaskNotAssignedRejected(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"
	task.Status = domain.StatusStarted

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return newAdminUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupTaskTestRouter(deps)

	body := `{"title":"Updated"}`
	req := httptest.NewRequest(http.MethodPut, "/tasks/task-1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestDeleteTaskHandlerSuccess(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return newAdminUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("admin-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "admin-1", Role: domain.RoleAdmin, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupTaskTestRouter(deps)

	req := httptest.NewRequest(http.MethodDelete, "/tasks/task-1", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
