package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	t.Run("crear sesion valida", func(t *testing.T) {
		s, err := NewSession("user-1", 24*time.Hour)

		assert.NoError(t, err)
		assert.NotNil(t, s)
		assert.Equal(t, "user-1", s.UserID)
		assert.Nil(t, s.RevokedAt)
		assert.False(t, s.CreatedAt.IsZero())
		assert.True(t, s.ExpiresAt.After(s.CreatedAt))
		assert.False(t, s.IsRevoked())
	})

	t.Run("crear sesion sin user_id devuelve error", func(t *testing.T) {
		s, err := NewSession("", 1*time.Hour)
		assert.ErrorIs(t, err, ErrSessionUserIDRequired)
		assert.Nil(t, s)
	})
}

func TestSessionIsValid(t *testing.T) {
	s, _ := NewSession("user-1", 1*time.Hour)

	t.Run("sesion recien creada es valida", func(t *testing.T) {
		assert.True(t, s.IsValid(time.Now()))
	})

	t.Run("sesion revocada no es valida", func(t *testing.T) {
		now := time.Now()
		s.Revoke(now)
		assert.False(t, s.IsValid(now.Add(time.Second)))
		assert.True(t, s.IsRevoked())
	})

	t.Run("sesion expirada no es valida", func(t *testing.T) {
		s2, _ := NewSession("user-1", 1*time.Hour)
		futureTime := s2.CreatedAt.Add(2 * time.Hour)
		assert.True(t, s2.IsExpired(futureTime))
		assert.False(t, s2.IsValid(futureTime))
	})
}

func TestSessionRevoke(t *testing.T) {
	s, _ := NewSession("user-1", 1*time.Hour)
	assert.Nil(t, s.RevokedAt)
	assert.False(t, s.IsRevoked())

	now := time.Now()
	s.Revoke(now)
	assert.NotNil(t, s.RevokedAt)
	assert.True(t, s.IsRevoked())
}

func TestSessionIsExpired(t *testing.T) {
	s, _ := NewSession("user-1", 1*time.Hour)

	t.Run("no expirada inmediatamente", func(t *testing.T) {
		assert.False(t, s.IsExpired(time.Now()))
	})

	t.Run("expirada despues del tiempo", func(t *testing.T) {
		futureTime := s.CreatedAt.Add(2 * time.Hour)
		assert.True(t, s.IsExpired(futureTime))
	})
}
