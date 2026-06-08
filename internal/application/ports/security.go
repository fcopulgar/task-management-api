package ports

import (
	"context"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type PasswordHasher interface {
	Hash(ctx context.Context, password string) (string, error)
	Compare(ctx context.Context, hash, password string) error
}

type TokenClaims struct {
	UserID    string
	Role      domain.Role
	SessionID string
}

type TokenService interface {
	Generate(ctx context.Context, userID string, role domain.Role, sessionID string, tokenDurationHours int) (string, error)
	Validate(ctx context.Context, tokenString string) (*TokenClaims, error)
}
