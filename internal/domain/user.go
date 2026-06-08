package domain

import (
	"errors"
	"time"
)

type User struct {
	ID                string
	Name              string
	Email             string
	PasswordHash      string
	Role              Role
	MustChangePassword bool
	Active            bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

var (
	ErrUserNameRequired = errors.New("nombre de usuario requerido")
	ErrEmailRequired    = errors.New("email requerido")
)

func NewUser(name, email, passwordHash string, role Role) (*User, error) {
	if name == "" {
		return nil, ErrUserNameRequired
	}
	if email == "" {
		return nil, ErrEmailRequired
	}
	if !role.IsValid() {
		return nil, ErrInvalidRole
	}

	now := time.Now()
	return &User{
		Name:              name,
		Email:             email,
		PasswordHash:      passwordHash,
		Role:              role,
		MustChangePassword: true,
		Active:            true,
		CreatedAt:         now,
		UpdatedAt:         now,
	}, nil
}

func (u *User) IsActive() bool {
	return u.Active
}

func (u *User) CanLogin() error {
	if !u.Active {
		return ErrUserInactive
	}
	return nil
}

func (u *User) HasRole(r Role) bool {
	return u.Role == r
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) IsExecutor() bool {
	return u.Role == RoleExecutor
}

func (u *User) IsAuditor() bool {
	return u.Role == RoleAuditor
}

func (u *User) MarkPasswordChanged() {
	u.MustChangePassword = false
	u.UpdatedAt = time.Now()
}

func (u *User) Deactivate() {
	u.Active = false
	u.UpdatedAt = time.Now()
}

func (u *User) Activate() {
	u.Active = true
	u.UpdatedAt = time.Now()
}
