package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/application/ports"
	"github.com/anomalyco/task-management-api/internal/domain"
)

type mockUserRepo struct {
	findByEmailFn func(ctx context.Context, email string) (*domain.User, error)
	findByIDFn    func(ctx context.Context, id string) (*domain.User, error)
	updateFn      func(ctx context.Context, user *domain.User) error
}

func (m *mockUserRepo) Create(ctx context.Context, user *domain.User) error { return nil }
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
func (m *mockUserRepo) List(ctx context.Context) ([]domain.User, error)   { return nil, nil }
func (m *mockUserRepo) Update(ctx context.Context, user *domain.User) error {
	if m.updateFn != nil {
		return m.updateFn(ctx, user)
	}
	return nil
}

type mockSessionRepo struct {
	createFn func(ctx context.Context, session *domain.Session) error
	revokeFn func(ctx context.Context, sessionID string) error
}

func (m *mockSessionRepo) Create(ctx context.Context, session *domain.Session) error {
	if m.createFn != nil {
		return m.createFn(ctx, session)
	}
	return nil
}
func (m *mockSessionRepo) FindByID(ctx context.Context, id string) (*domain.Session, error) {
	return nil, nil
}
func (m *mockSessionRepo) Revoke(ctx context.Context, sessionID string) error {
	if m.revokeFn != nil {
		return m.revokeFn(ctx, sessionID)
	}
	return nil
}

type mockHasher struct {
	compareFn func(ctx context.Context, hash, password string) error
	hashFn    func(ctx context.Context, password string) (string, error)
}

func (m *mockHasher) Hash(ctx context.Context, password string) (string, error) {
	if m.hashFn != nil {
		return m.hashFn(ctx, password)
	}
	return "hashed-" + password, nil
}
func (m *mockHasher) Compare(ctx context.Context, hash, password string) error {
	if m.compareFn != nil {
		return m.compareFn(ctx, hash, password)
	}
	if hash == "hashed-"+password {
		return nil
	}
	return assert.AnError
}

type mockTokenSvc struct {
	generateFn func(ctx context.Context, userID string, role domain.Role, sessionID string, tokenDurationHours int) (string, error)
}

func (m *mockTokenSvc) Generate(ctx context.Context, userID string, role domain.Role, sessionID string, tokenDurationHours int) (string, error) {
	if m.generateFn != nil {
		return m.generateFn(ctx, userID, role, sessionID, tokenDurationHours)
	}
	return "mock-token-" + userID, nil
}
func (m *mockTokenSvc) Validate(ctx context.Context, tokenString string) (*ports.TokenClaims, error) {
	return nil, nil
}

func TestLoginSuccess(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hashed-secret", domain.RoleExecutor)
	user.ID = "user-1"

	userRepo := &mockUserRepo{
		findByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return user, nil
		},
	}
	sessionRepo := &mockSessionRepo{
		createFn: func(ctx context.Context, session *domain.Session) error {
			session.ID = "session-1"
			return nil
		},
	}
	hasher := &mockHasher{}
	tokenSvc := &mockTokenSvc{
		generateFn: func(ctx context.Context, userID string, role domain.Role, sessionID string, tokenDurationHours int) (string, error) {
			return "jwt-token-" + userID, nil
		},
	}

	deps := Dependencies{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
		Hasher:      hasher,
		TokenSvc:    tokenSvc,
	}

	uc := NewAuthUseCase(deps)
	output, err := uc.Login(context.Background(), dto.LoginInput{
		Email:    "test@test.com",
		Password: "secret",
	})

	require.NoError(t, err)
	assert.Equal(t, "jwt-token-"+user.ID, output.Token)
}

func TestLoginUserNotFound(t *testing.T) {
	userRepo := &mockUserRepo{
		findByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return nil, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewAuthUseCase(deps)
	_, err := uc.Login(context.Background(), dto.LoginInput{
		Email:    "no@test.com",
		Password: "secret",
	})

	assert.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestLoginInactiveUser(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hashed-secret", domain.RoleExecutor)
	user.ID = "user-1"
	user.Deactivate()

	userRepo := &mockUserRepo{
		findByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return user, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewAuthUseCase(deps)
	_, err := uc.Login(context.Background(), dto.LoginInput{
		Email:    "test@test.com",
		Password: "secret",
	})

	assert.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestLoginWrongPassword(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hashed-secret", domain.RoleExecutor)
	user.ID = "user-1"

	userRepo := &mockUserRepo{
		findByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return user, nil
		},
	}
	hasher := &mockHasher{
		compareFn: func(ctx context.Context, hash, password string) error {
			return assert.AnError
		},
	}

	deps := Dependencies{UserRepo: userRepo, Hasher: hasher}
	uc := NewAuthUseCase(deps)
	_, err := uc.Login(context.Background(), dto.LoginInput{
		Email:    "test@test.com",
		Password: "wrong",
	})

	assert.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestLogoutSuccess(t *testing.T) {
	sessionRepo := &mockSessionRepo{
		revokeFn: func(ctx context.Context, sessionID string) error {
			assert.Equal(t, "session-1", sessionID)
			return nil
		},
	}

	deps := Dependencies{SessionRepo: sessionRepo}
	uc := NewAuthUseCase(deps)
	err := uc.Logout(context.Background(), "session-1")

	assert.NoError(t, err)
}

func TestChangePasswordSuccess(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hashed-old", domain.RoleExecutor)
	user.ID = "user-1"
	assert.True(t, user.MustChangePassword)

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return user, nil
		},
		updateFn: func(ctx context.Context, u *domain.User) error {
			return nil
		},
	}
	hasher := &mockHasher{}

	deps := Dependencies{UserRepo: userRepo, Hasher: hasher}
	uc := NewAuthUseCase(deps)
	err := uc.ChangePassword(context.Background(), user.ID, "old", "new")

	require.NoError(t, err)
	assert.False(t, user.MustChangePassword)
	assert.Equal(t, "hashed-new", user.PasswordHash)
}

func TestChangePasswordWrongOldPassword(t *testing.T) {
	user, _ := domain.NewUser("Test", "test@test.com", "hashed-old", domain.RoleExecutor)
	user.ID = "user-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return user, nil
		},
	}
	hasher := &mockHasher{
		compareFn: func(ctx context.Context, hash, password string) error {
			return assert.AnError
		},
	}

	deps := Dependencies{UserRepo: userRepo, Hasher: hasher}
	uc := NewAuthUseCase(deps)
	err := uc.ChangePassword(context.Background(), user.ID, "wrong", "new")

	assert.Error(t, err)
}

func TestChangePasswordUserNotFound(t *testing.T) {
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return nil, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewAuthUseCase(deps)
	err := uc.ChangePassword(context.Background(), "nonexistent", "old", "new")

	assert.ErrorIs(t, err, ErrUserNotFound)
}
