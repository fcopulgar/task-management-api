package security

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/anomalyco/task-management-api/internal/application/ports"
	"github.com/anomalyco/task-management-api/internal/domain"
)

type JWTTokenService struct {
	secret []byte
}

func NewJWTTokenService(secret string) *JWTTokenService {
	return &JWTTokenService{secret: []byte(secret)}
}

type customClaims struct {
	jwt.RegisteredClaims
	Role      string `json:"role"`
	SessionID string `json:"session_id"`
}

func (s *JWTTokenService) Generate(_ context.Context, userID string, role domain.Role, sessionID string, tokenDurationHours int) (string, error) {
	now := time.Now()
	claims := customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(tokenDurationHours) * time.Hour)),
		},
		Role:      string(role),
		SessionID: sessionID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *JWTTokenService) Validate(_ context.Context, tokenString string) (*ports.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metodo de firma inesperado: %v", t.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token invalido")
	}

	return &ports.TokenClaims{
		UserID:    claims.Subject,
		Role:      domain.Role(claims.Role),
		SessionID: claims.SessionID,
	}, nil
}
