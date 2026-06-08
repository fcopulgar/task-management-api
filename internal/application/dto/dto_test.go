package dto

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/anomalyco/task-management-api/internal/domain"
)

func TestUserToOutput(t *testing.T) {
	u := &domain.User{
		ID:                 "user-1",
		Name:               "Juan",
		Email:              "juan@test.com",
		PasswordHash:       "secret-hash",
		Role:               domain.RoleExecutor,
		MustChangePassword: true,
		Active:             true,
		CreatedAt:          time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:          time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	output := UserToOutput(u)

	assert.Equal(t, "user-1", output.ID)
	assert.Equal(t, "Juan", output.Name)
	assert.Equal(t, "juan@test.com", output.Email)
	assert.Equal(t, domain.RoleExecutor, output.Role)
	assert.True(t, output.MustChangePassword)
	assert.True(t, output.Active)
	assert.Equal(t, time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), output.CreatedAt)
	assert.Equal(t, time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), output.UpdatedAt)
}

func TestTaskToOutput(t *testing.T) {
	dueAt := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)
	createdAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	tsk := &domain.Task{
		ID:          "task-1",
		Title:       "Tarea 1",
		Description: "Descripcion",
		DueAt:       dueAt,
		Status:      domain.StatusAssigned,
		AssigneeID:  "exec-1",
		CreatedBy:   "admin-1",
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	}

	output := TaskToOutput(tsk)

	assert.Equal(t, "task-1", output.ID)
	assert.Equal(t, "Tarea 1", output.Title)
	assert.Equal(t, "Descripcion", output.Description)
	assert.Equal(t, dueAt, output.DueAt)
	assert.Equal(t, domain.StatusAssigned, output.Status)
	assert.Equal(t, "exec-1", output.AssigneeID)
	assert.Equal(t, "admin-1", output.CreatedBy)
}

func TestCommentToOutput(t *testing.T) {
	createdAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	c := &domain.Comment{
		ID:        "comment-1",
		TaskID:    "task-1",
		UserID:    "exec-1",
		Comment:   "Comentario de prueba",
		CreatedAt: createdAt,
	}

	output := CommentToOutput(c)

	assert.Equal(t, "comment-1", output.ID)
	assert.Equal(t, "task-1", output.TaskID)
	assert.Equal(t, "exec-1", output.UserID)
	assert.Equal(t, "Comentario de prueba", output.Comment)
	assert.Equal(t, createdAt, output.CreatedAt)
}

func TestLoginInput(t *testing.T) {
	input := LoginInput{
		Email:    "juan@test.com",
		Password: "secret",
	}

	assert.Equal(t, "juan@test.com", input.Email)
	assert.Equal(t, "secret", input.Password)
}

func TestCreateUserInput(t *testing.T) {
	input := CreateUserInput{
		Name:     "Nuevo Executor",
		Email:    "exec@test.com",
		Password: "temp123",
		Role:     domain.RoleExecutor,
	}

	assert.Equal(t, "Nuevo Executor", input.Name)
	assert.Equal(t, domain.RoleExecutor, input.Role)
}

func TestUpdateUserInput(t *testing.T) {
	active := false
	input := UpdateUserInput{
		ID:     "user-1",
		Name:   "Actualizado",
		Email:  "update@test.com",
		Role:   domain.RoleAuditor,
		Active: &active,
	}

	assert.Equal(t, "user-1", input.ID)
	assert.False(t, *input.Active)
}

func TestCreateTaskInput(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	input := CreateTaskInput{
		Title:       "Nueva Tarea",
		Description: "Desc",
		DueAt:       dueAt,
		AssigneeID:  "exec-1",
	}

	assert.Equal(t, "Nueva Tarea", input.Title)
	assert.Equal(t, "exec-1", input.AssigneeID)
}

func TestTransitionTaskInput(t *testing.T) {
	input := TransitionTaskInput{
		TaskID:    "task-1",
		UserID:    "exec-1",
		NewStatus: domain.StatusStarted,
	}

	assert.Equal(t, domain.StatusStarted, input.NewStatus)
}
