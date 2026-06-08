package dto

import (
	"time"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type CreateUserInput struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Role     domain.Role `json:"role"`
}

type UpdateUserInput struct {
	ID     string      `json:"-"`
	Name   string      `json:"name"`
	Email  string      `json:"email"`
	Role   domain.Role `json:"role"`
	Active *bool       `json:"active"`
}

type UserOutput struct {
	ID                 string
	Name               string
	Email              string
	Role               domain.Role
	MustChangePassword bool
	Active             bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func UserToOutput(u *domain.User) UserOutput {
	return UserOutput{
		ID:                 u.ID,
		Name:               u.Name,
		Email:              u.Email,
		Role:               u.Role,
		MustChangePassword: u.MustChangePassword,
		Active:             u.Active,
		CreatedAt:          u.CreatedAt,
		UpdatedAt:          u.UpdatedAt,
	}
}
