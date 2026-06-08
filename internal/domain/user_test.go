package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name         string
		userName     string
		email        string
		passwordHash string
		role         Role
		wantErr      error
	}{
		{"crear usuario valido executor", "Juan", "juan@test.com", "hash123", RoleExecutor, nil},
		{"crear usuario valido auditor", "Ana", "ana@test.com", "hash123", RoleAuditor, nil},
		{"crear usuario admin", "Admin", "admin@test.com", "hash123", RoleAdmin, nil},
		{"nombre vacio", "", "test@test.com", "hash", RoleExecutor, ErrUserNameRequired},
		{"email vacio", "Juan", "", "hash", RoleExecutor, ErrEmailRequired},
		{"rol invalido", "Juan", "juan@test.com", "hash", Role("INVALID"), ErrInvalidRole},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUser(tt.userName, tt.email, tt.passwordHash, tt.role)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, u)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, u)
			assert.Equal(t, tt.userName, u.Name)
			assert.Equal(t, tt.email, u.Email)
			assert.Equal(t, tt.passwordHash, u.PasswordHash)
			assert.Equal(t, tt.role, u.Role)
			assert.True(t, u.Active)
			assert.True(t, u.MustChangePassword)
			assert.False(t, u.CreatedAt.IsZero())
			assert.False(t, u.UpdatedAt.IsZero())
		})
	}
}

func TestUserCanLogin(t *testing.T) {
	u, _ := NewUser("Juan", "juan@test.com", "hash", RoleExecutor)

	t.Run("usuario activo puede iniciar sesion", func(t *testing.T) {
		assert.NoError(t, u.CanLogin())
	})

	t.Run("usuario inactivo no puede iniciar sesion", func(t *testing.T) {
		u.Deactivate()
		assert.ErrorIs(t, u.CanLogin(), ErrUserInactive)
	})

	t.Run("usuario reactivado puede iniciar sesion", func(t *testing.T) {
		u.Activate()
		assert.NoError(t, u.CanLogin())
	})
}

func TestUserRoles(t *testing.T) {
	admin, _ := NewUser("Admin", "admin@test.com", "hash", RoleAdmin)
	executor, _ := NewUser("Exec", "exec@test.com", "hash", RoleExecutor)
	auditor, _ := NewUser("Aud", "aud@test.com", "hash", RoleAuditor)

	assert.True(t, admin.IsAdmin())
	assert.False(t, admin.IsExecutor())
	assert.False(t, admin.IsAuditor())

	assert.False(t, executor.IsAdmin())
	assert.True(t, executor.IsExecutor())
	assert.False(t, executor.IsAuditor())

	assert.False(t, auditor.IsAdmin())
	assert.False(t, auditor.IsExecutor())
	assert.True(t, auditor.IsAuditor())

	assert.True(t, admin.HasRole(RoleAdmin))
	assert.False(t, admin.HasRole(RoleExecutor))
}

func TestUserMarkPasswordChanged(t *testing.T) {
	u, _ := NewUser("Juan", "juan@test.com", "hash", RoleExecutor)
	assert.True(t, u.MustChangePassword)

	u.MarkPasswordChanged()
	assert.False(t, u.MustChangePassword)
}

func TestUserDeactivateActivate(t *testing.T) {
	u, _ := NewUser("Juan", "juan@test.com", "hash", RoleExecutor)
	assert.True(t, u.IsActive())

	u.Deactivate()
	assert.False(t, u.IsActive())

	u.Activate()
	assert.True(t, u.IsActive())
}

func TestRoleIsValid(t *testing.T) {
	assert.True(t, RoleAdmin.IsValid())
	assert.True(t, RoleExecutor.IsValid())
	assert.True(t, RoleAuditor.IsValid())
	assert.False(t, Role("SUPER_ADMIN").IsValid())
	assert.False(t, Role("").IsValid())
}

func TestRoleString(t *testing.T) {
	assert.Equal(t, "ADMIN", RoleAdmin.String())
	assert.Equal(t, "EXECUTOR", RoleExecutor.String())
	assert.Equal(t, "AUDITOR", RoleAuditor.String())
}

func TestNewUser_DefaultTimestamps(t *testing.T) {
	before := time.Now()
	u, err := NewUser("Test", "test@test.com", "hash", RoleExecutor)
	after := time.Now()

	assert.NoError(t, err)
	assert.True(t, u.CreatedAt.After(before) || u.CreatedAt.Equal(before))
	assert.True(t, u.CreatedAt.Before(after) || u.CreatedAt.Equal(after))
}
