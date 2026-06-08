package domain

import (
	"errors"
	"time"
)

type Session struct {
	ID        string
	UserID    string
	RevokedAt *time.Time
	ExpiresAt time.Time
	CreatedAt time.Time
}

var (
	ErrSessionUserIDRequired = errors.New("user_id requerido para la sesion")
)

func NewSession(userID string, duration time.Duration) (*Session, error) {
	if userID == "" {
		return nil, ErrSessionUserIDRequired
	}

	now := time.Now()
	return &Session{
		UserID:    userID,
		ExpiresAt: now.Add(duration),
		CreatedAt: now,
	}, nil
}

func (s *Session) IsRevoked() bool {
	return s.RevokedAt != nil
}

func (s *Session) IsExpired(now time.Time) bool {
	return now.After(s.ExpiresAt)
}

func (s *Session) IsValid(now time.Time) bool {
	return !s.IsRevoked() && !s.IsExpired(now)
}

func (s *Session) Revoke(now time.Time) {
	s.RevokedAt = &now
}
