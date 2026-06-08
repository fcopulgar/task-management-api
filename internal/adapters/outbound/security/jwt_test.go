package security

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anomalyco/task-management-api/internal/domain"
)

func TestJWTTokenServiceGenerateAndValidate(t *testing.T) {
	svc := NewJWTTokenService("test-secret-key")
	ctx := context.Background()

	token, err := svc.Generate(ctx, "user-1", domain.RoleExecutor, "session-1", 24)
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := svc.Validate(ctx, token)
	require.NoError(t, err)
	assert.Equal(t, "user-1", claims.UserID)
	assert.Equal(t, domain.RoleExecutor, claims.Role)
	assert.Equal(t, "session-1", claims.SessionID)
}

func TestJWTTokenServiceValidateInvalidToken(t *testing.T) {
	svc := NewJWTTokenService("test-secret-key")
	ctx := context.Background()

	_, err := svc.Validate(ctx, "invalid.token.string")
	assert.Error(t, err)
}

func TestJWTTokenServiceValidateWrongSecret(t *testing.T) {
	svc1 := NewJWTTokenService("secret-1")
	svc2 := NewJWTTokenService("secret-2")
	ctx := context.Background()

	token, err := svc1.Generate(ctx, "user-1", domain.RoleExecutor, "session-1", 24)
	require.NoError(t, err)

	_, err = svc2.Validate(ctx, token)
	assert.Error(t, err)
}

func TestJWTTokenServiceExpiredToken(t *testing.T) {
	svc := NewJWTTokenService("test-secret-key")
	ctx := context.Background()

	token, err := svc.Generate(ctx, "user-1", domain.RoleExecutor, "session-1", 0)
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	_, err = svc.Validate(ctx, token)
	assert.Error(t, err)
}

func TestJWTTokenServiceAllRoles(t *testing.T) {
	svc := NewJWTTokenService("test-secret-key")
	ctx := context.Background()

	roles := []domain.Role{domain.RoleAdmin, domain.RoleExecutor, domain.RoleAuditor}
	for _, role := range roles {
		token, err := svc.Generate(ctx, "user-1", role, "session-1", 24)
		require.NoError(t, err)

		claims, err := svc.Validate(ctx, token)
		require.NoError(t, err)
		assert.Equal(t, role, claims.Role)
	}
}
