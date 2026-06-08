package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/anomalyco/task-management-api/internal/application"
	"github.com/anomalyco/task-management-api/internal/application/ports"
	"github.com/anomalyco/task-management-api/internal/domain"
)

func setupExecutorTestRouter(deps application.Dependencies) chi.Router {
	r := chi.NewRouter()
	SetupExecutorRoutes(r, deps)
	return r
}

func newExecutorUser() *domain.User {
	u, _ := domain.NewUser("Exec", "exec@test.com", "hash", domain.RoleExecutor)
	u.ID = "exec-1"
	u.MarkPasswordChanged()
	return u
}

type mockCommentRepo struct {
	createFn       func(ctx context.Context, comment *domain.Comment) error
	findByTaskIDFn func(ctx context.Context, taskID string) ([]domain.Comment, error)
}

func (m *mockCommentRepo) Create(ctx context.Context, c *domain.Comment) error {
	if m.createFn != nil {
		return m.createFn(ctx, c)
	}
	return nil
}
func (m *mockCommentRepo) FindByTaskID(ctx context.Context, taskID string) ([]domain.Comment, error) {
	if m.findByTaskIDFn != nil {
		return m.findByTaskIDFn(ctx, taskID)
	}
	return nil, nil
}

func TestListMyTasksHandlerSuccess(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	t1, _ := domain.NewTask("T1", "D1", dueAt, "exec-1", "admin-1")
	t1.ID = "t1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return newExecutorUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("exec-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	taskRepo := &mockTaskRepo{
		listByAssigneeFn: func(ctx context.Context, assigneeID string) ([]domain.Task, error) {
			return []domain.Task{*t1}, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "exec-1", Role: domain.RoleExecutor, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupExecutorTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/me/tasks", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetMyTaskHandlerSuccess(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return newExecutorUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("exec-1", 24*time.Hour)
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
			return &ports.TokenClaims{UserID: "exec-1", Role: domain.RoleExecutor, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupExecutorTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/me/tasks/task-1", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetMyTaskHandlerNotOwner(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-2", "admin-1")
	task.ID = "task-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return newExecutorUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("exec-1", 24*time.Hour)
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
			return &ports.TokenClaims{UserID: "exec-1", Role: domain.RoleExecutor, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupExecutorTestRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/me/tasks/task-1", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestTransitionStatusHandlerSuccess(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return newExecutorUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("exec-1", 24*time.Hour)
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
			return &ports.TokenClaims{UserID: "exec-1", Role: domain.RoleExecutor, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupExecutorTestRouter(deps)

	body := `{"new_status":"STARTED"}`
	req := httptest.NewRequest(http.MethodPatch, "/me/tasks/task-1/status", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTransitionStatusHandlerOverdue(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	task, _ := domain.NewTask("T", "D", past, "exec-1", "admin-1")
	task.ID = "task-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return newExecutorUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("exec-1", 24*time.Hour)
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
			return &ports.TokenClaims{UserID: "exec-1", Role: domain.RoleExecutor, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, TokenSvc: tokenSvc,
	}
	router := setupExecutorTestRouter(deps)

	body := `{"new_status":"STARTED"}`
	req := httptest.NewRequest(http.MethodPatch, "/me/tasks/task-1/status", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestCommentOnTaskHandlerSuccess(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	task, _ := domain.NewTask("T", "D", past, "exec-1", "admin-1")
	task.ID = "task-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return newExecutorUser(), nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("exec-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}
	commentRepo := &mockCommentRepo{}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "exec-1", Role: domain.RoleExecutor, SessionID: "s1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TaskRepo: taskRepo, CommentRepo: commentRepo,
		TokenSvc: tokenSvc,
	}
	router := setupExecutorTestRouter(deps)

	body := `{"comment":"Comentario de prueba"}`
	req := httptest.NewRequest(http.MethodPost, "/me/tasks/task-1/comments", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
