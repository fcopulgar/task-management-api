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

type mockHasher struct {
	compareFn func(ctx context.Context, hash, password string) error
	hashFn    func(ctx context.Context, password string) (string, error)
}

func (m *mockHasher) Hash(ctx context.Context, p string) (string, error) {
	if m.hashFn != nil {
		return m.hashFn(ctx, p)
	}
	return "hashed-" + p, nil
}
func (m *mockHasher) Compare(ctx context.Context, h, p string) error {
	if m.compareFn != nil {
		return m.compareFn(ctx, h, p)
	}
	return nil
}

type mockTokenSvc struct {
	generateFn func(ctx context.Context, userID string, role domain.Role, sessionID string, dur int) (string, error)
	validateFn func(ctx context.Context, token string) (*ports.TokenClaims, error)
}

func (m *mockTokenSvc) Generate(ctx context.Context, uid string, r domain.Role, sid string, dur int) (string, error) {
	if m.generateFn != nil {
		return m.generateFn(ctx, uid, r, sid, dur)
	}
	return "token", nil
}
func (m *mockTokenSvc) Validate(ctx context.Context, tok string) (*ports.TokenClaims, error) {
	if m.validateFn != nil {
		return m.validateFn(ctx, tok)
	}
	return nil, nil
}

type mockUserRepo struct {
	findByEmailFn func(ctx context.Context, email string) (*domain.User, error)
	findByIDFn    func(ctx context.Context, id string) (*domain.User, error)
	updateFn      func(ctx context.Context, user *domain.User) error
}

func (m *mockUserRepo) Create(ctx context.Context, u *domain.User) error    { return nil }
func (m *mockUserRepo) List(ctx context.Context) ([]domain.User, error)      { return nil, nil }
func (m *mockUserRepo) FindByID(ctx context.Context, id string) (*domain.User, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}
	return nil, nil
}
func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if m.findByEmailFn != nil {
		return m.findByEmailFn(ctx, email)
	}
	return nil, nil
}
func (m *mockUserRepo) Update(ctx context.Context, u *domain.User) error {
	if m.updateFn != nil {
		return m.updateFn(ctx, u)
	}
	return nil
}

type mockSessionRepo struct {
	createFn    func(ctx context.Context, s *domain.Session) error
	findByIDFn  func(ctx context.Context, id string) (*domain.Session, error)
	revokeFn    func(ctx context.Context, id string) error
}

func (m *mockSessionRepo) Create(ctx context.Context, s *domain.Session) error {
	if m.createFn != nil {
		return m.createFn(ctx, s)
	}
	return nil
}
func (m *mockSessionRepo) FindByID(ctx context.Context, id string) (*domain.Session, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}
	return nil, nil
}
func (m *mockSessionRepo) Revoke(ctx context.Context, id string) error {
	if m.revokeFn != nil {
		return m.revokeFn(ctx, id)
	}
	return nil
}

func setupTestRouter(deps application.Dependencies) chi.Router {
	r := chi.NewRouter()
	SetupAuthRoutes(r, deps)
	return r
}

func TestLoginHandlerSuccess(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hashed-secret", domain.RoleExecutor)
	user.ID = "user-1"
	userRepo := &mockUserRepo{
		findByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return user, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		createFn: func(ctx context.Context, s *domain.Session) error {
			s.ID = "session-1"
			return nil
		},
	}
	hasher := &mockHasher{}
	tokenSvc := &mockTokenSvc{
		generateFn: func(ctx context.Context, uid string, r domain.Role, sid string, dur int) (string, error) {
			return "jwt-token", nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		Hasher: hasher, TokenSvc: tokenSvc,
	}

	router := setupTestRouter(deps)

	body := `{"email":"test@test.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var output dto.LoginOutput
	json.NewDecoder(w.Body).Decode(&output)
	assert.Equal(t, "jwt-token", output.Token)
}

func TestLoginHandlerInvalidCredentials(t *testing.T) {
	userRepo := &mockUserRepo{
		findByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return nil, nil
		},
	}
	deps := application.Dependencies{UserRepo: userRepo}
	router := setupTestRouter(deps)

	body := `{"email":"no@test.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestLoginHandlerBadRequest(t *testing.T) {
	deps := application.Dependencies{}
	router := setupTestRouter(deps)

	req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader("invalid"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestChangePasswordHandlerSuccess(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hashed-old", domain.RoleExecutor)
	user.ID = "test-user"
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return user, nil
		},
		updateFn: func(ctx context.Context, u *domain.User) error { return nil },
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("test-user", 24*time.Hour)
			s.ID = id
			return s, nil
		},
	}
	hasher := &mockHasher{}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: user.ID, Role: user.Role, SessionID: "session-1"}, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		Hasher: hasher, TokenSvc: tokenSvc,
	}
	router := setupTestRouter(deps)

	body := `{"old_password":"old","new_password":"new"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/password", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.False(t, user.MustChangePassword)
}

func TestChangePasswordHandlerNoAuth(t *testing.T) {
	deps := application.Dependencies{}
	router := setupTestRouter(deps)

	body := `{"old_password":"old","new_password":"new"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/password", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestLogoutHandlerSuccess(t *testing.T) {
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			s, _ := domain.NewSession("user-1", 24*time.Hour)
			s.ID = id
			return s, nil
		},
		revokeFn: func(ctx context.Context, id string) error { return nil },
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: "user-1", Role: domain.RoleExecutor, SessionID: "session-1"}, nil
		},
	}
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			u, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
			u.ID = "user-1"
			return u, nil
		},
	}

	deps := application.Dependencies{
		UserRepo: userRepo, SessionRepo: sessionRepo,
		TokenSvc: tokenSvc,
	}
	router := setupTestRouter(deps)

	req := httptest.NewRequest(http.MethodPost, "/auth/logout", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
