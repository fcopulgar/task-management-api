package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/anomalyco/task-management-api/internal/application/ports"
	"github.com/anomalyco/task-management-api/internal/domain"
)

func TestAuthenticateNoToken(t *testing.T) {
	mw := NewAuthMiddleware(nil, nil, nil)
	handler := mw.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthenticateInvalidToken(t *testing.T) {
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return nil, assert.AnError
		},
	}
	mw := NewAuthMiddleware(nil, nil, tokenSvc)
	handler := mw.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthenticateSessionRevoked(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	user.ID = "user-1"
	revokedAt := time.Now()
	session := &domain.Session{
		ID:        "session-1",
		UserID:    user.ID,
		RevokedAt: &revokedAt,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		CreatedAt: time.Now(),
	}

	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: user.ID, Role: user.Role, SessionID: "session-1"}, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			return session, nil
		},
	}

	mw := NewAuthMiddleware(nil, sessionRepo, tokenSvc)
	handler := mw.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer token")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthenticateSuccess(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	user.ID = "user-1"
	session, _ := domain.NewSession(user.ID, 24*time.Hour)
	session.ID = "session-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return user, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Session, error) {
			return session, nil
		},
	}
	tokenSvc := &mockTokenSvc{
		validateFn: func(ctx context.Context, tok string) (*ports.TokenClaims, error) {
			return &ports.TokenClaims{UserID: user.ID, Role: user.Role, SessionID: session.ID}, nil
		},
	}

	mw := NewAuthMiddleware(userRepo, sessionRepo, tokenSvc)

	var capturedAuth *AuthInfo
	handler := mw.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		info, ok := GetAuthInfo(r.Context())
		assert.True(t, ok)
		capturedAuth = info
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, capturedAuth)
	assert.Equal(t, user.ID, capturedAuth.UserID)
	assert.Equal(t, session.ID, capturedAuth.SessionID)
}

func TestRequirePasswordNotTemporaryBlocks(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return user, nil
		},
	}

	mw := NewAuthMiddleware(userRepo, nil, nil)

	handler := mw.RequirePasswordNotTemporary(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	ctx := SetAuthInfo(req.Context(), &ports.TokenClaims{UserID: user.ID, Role: user.Role, SessionID: "s1"})
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestRequirePasswordNotTemporaryAllowsPasswordRoute(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return user, nil
		},
	}

	mw := NewAuthMiddleware(userRepo, nil, nil)

	handler := mw.RequirePasswordNotTemporary(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodPost, "/auth/password", nil)
	ctx := SetAuthInfo(req.Context(), &ports.TokenClaims{UserID: user.ID, Role: user.Role, SessionID: "s1"})
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequirePasswordNotTemporaryAllowsLogoutRoute(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return user, nil
		},
	}

	mw := NewAuthMiddleware(userRepo, nil, nil)

	handler := mw.RequirePasswordNotTemporary(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodPost, "/auth/logout", nil)
	ctx := SetAuthInfo(req.Context(), &ports.TokenClaims{UserID: user.ID, Role: user.Role, SessionID: "s1"})
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequirePasswordNotTemporaryAllowsWhenPasswordChanged(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	user.MarkPasswordChanged()
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return user, nil
		},
	}

	mw := NewAuthMiddleware(userRepo, nil, nil)

	handler := mw.RequirePasswordNotTemporary(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	ctx := SetAuthInfo(req.Context(), &ports.TokenClaims{UserID: user.ID, Role: user.Role, SessionID: "s1"})
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAuthInfoNotSet(t *testing.T) {
	info, ok := GetAuthInfo(context.Background())
	assert.False(t, ok)
	assert.Nil(t, info)
}
