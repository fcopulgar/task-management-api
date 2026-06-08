package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/domain"
)

func TestCreateUserSuccess(t *testing.T) {
	userRepo := &mockUserRepo{
		findByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return nil, nil
		},
	}
	hasher := &mockHasher{}

	deps := Dependencies{UserRepo: userRepo, Hasher: hasher}
	uc := NewUserUseCase(deps)

	output, err := uc.CreateUser(context.Background(), dto.CreateUserInput{
		Name:     "Nuevo Executor",
		Email:    "exec@test.com",
		Password: "temp123",
		Role:     domain.RoleExecutor,
	})

	require.NoError(t, err)
	assert.Equal(t, "Nuevo Executor", output.Name)
	assert.Equal(t, "exec@test.com", output.Email)
	assert.Equal(t, domain.RoleExecutor, output.Role)
	assert.True(t, output.MustChangePassword)
	assert.True(t, output.Active)
}

func TestCreateUserCannotCreateAdmin(t *testing.T) {
	deps := Dependencies{}
	uc := NewUserUseCase(deps)

	_, err := uc.CreateUser(context.Background(), dto.CreateUserInput{
		Name:     "Admin",
		Email:    "admin@test.com",
		Password: "secret",
		Role:     domain.RoleAdmin,
	})

	assert.ErrorIs(t, err, ErrCannotCreateAdmin)
}

func TestCreateUserEmailAlreadyExists(t *testing.T) {
	existing, _ := domain.NewUser("Existente", "dup@test.com", "hash", domain.RoleExecutor)
	existing.ID = "existing-id"

	userRepo := &mockUserRepo{
		findByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
			return existing, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewUserUseCase(deps)

	_, err := uc.CreateUser(context.Background(), dto.CreateUserInput{
		Name:     "Otro",
		Email:    "dup@test.com",
		Password: "secret",
		Role:     domain.RoleExecutor,
	})

	assert.ErrorIs(t, err, ErrEmailAlreadyExists)
}

func TestListUsers(t *testing.T) {
	u1, _ := domain.NewUser("User 1", "u1@test.com", "hash1", domain.RoleExecutor)
	u1.ID = "id-1"
	u2, _ := domain.NewUser("User 2", "u2@test.com", "hash2", domain.RoleAuditor)
	u2.ID = "id-2"

	userRepo := &mockUserRepo{
		listFn: func(ctx context.Context) ([]domain.User, error) {
			return []domain.User{*u1, *u2}, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewUserUseCase(deps)

	users, err := uc.ListUsers(context.Background())
	require.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "User 1", users[0].Name)
	assert.Equal(t, "User 2", users[1].Name)
}

func TestGetUserNotFound(t *testing.T) {
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return nil, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewUserUseCase(deps)

	_, err := uc.GetUser(context.Background(), "nonexistent")
	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestGetUserSuccess(t *testing.T) {
	u, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleAuditor)
	u.ID = "user-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return u, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewUserUseCase(deps)

	output, err := uc.GetUser(context.Background(), "user-1")
	require.NoError(t, err)
	assert.Equal(t, "Test", output.Name)
	assert.Equal(t, domain.RoleAuditor, output.Role)
}

func TestUpdateUserName(t *testing.T) {
	u, _ := domain.NewUser("Original", "orig@test.com", "hash", domain.RoleExecutor)
	u.ID = "user-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return u, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewUserUseCase(deps)

	output, err := uc.UpdateUser(context.Background(), dto.UpdateUserInput{
		ID:   "user-1",
		Name: "Actualizado",
	})
	require.NoError(t, err)
	assert.Equal(t, "Actualizado", output.Name)
}

func TestUpdateUserDeactivate(t *testing.T) {
	u, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	u.ID = "user-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return u, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewUserUseCase(deps)

	active := false
	output, err := uc.UpdateUser(context.Background(), dto.UpdateUserInput{
		ID:     "user-1",
		Active: &active,
	})
	require.NoError(t, err)
	assert.False(t, output.Active)
}

func TestDeactivateUser(t *testing.T) {
	u, _ := domain.NewUser("Test", "test@test.com", "hash", domain.RoleExecutor)
	u.ID = "user-1"

	updateCalled := false
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return u, nil
		},
		updateFn: func(ctx context.Context, user *domain.User) error {
			updateCalled = true
			return nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewUserUseCase(deps)

	err := uc.DeactivateUser(context.Background(), "user-1")
	require.NoError(t, err)
	assert.True(t, updateCalled)
	assert.False(t, u.IsActive())
}

func TestDeactivateUserNotFound(t *testing.T) {
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return nil, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewUserUseCase(deps)

	err := uc.DeactivateUser(context.Background(), "nonexistent")
	assert.ErrorIs(t, err, ErrUserNotFound)
}
