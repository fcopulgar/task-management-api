package persistence

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/anomalyco/task-management-api/internal/domain"
)

func TestUserMapperRoundTrip(t *testing.T) {
	u := &domain.User{
		ID:                 "user-1",
		Name:               "Juan",
		Email:              "juan@test.com",
		PasswordHash:       "hash123",
		Role:               domain.RoleExecutor,
		MustChangePassword: true,
		Active:             true,
		CreatedAt:          time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:          time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC),
	}

	model := UserToModel(u)
	result := UserFromModel(model)

	assert.Equal(t, u.ID, result.ID)
	assert.Equal(t, u.Name, result.Name)
	assert.Equal(t, u.Email, result.Email)
	assert.Equal(t, u.PasswordHash, result.PasswordHash)
	assert.Equal(t, u.Role, result.Role)
	assert.Equal(t, u.MustChangePassword, result.MustChangePassword)
	assert.Equal(t, u.Active, result.Active)
	assert.True(t, u.CreatedAt.Equal(result.CreatedAt))
	assert.True(t, u.UpdatedAt.Equal(result.UpdatedAt))
}

func TestUserModelTableName(t *testing.T) {
	m := UserModel{}
	assert.Equal(t, "users", m.TableName())
}

func TestSessionMapperRoundTrip(t *testing.T) {
	revokedAt := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	s := &domain.Session{
		ID:        "session-1",
		UserID:    "user-1",
		RevokedAt: &revokedAt,
		ExpiresAt: time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	model := SessionToModel(s)
	result := SessionFromModel(model)

	assert.Equal(t, s.ID, result.ID)
	assert.Equal(t, s.UserID, result.UserID)
	assert.NotNil(t, result.RevokedAt)
	assert.True(t, s.ExpiresAt.Equal(result.ExpiresAt))
	assert.True(t, s.CreatedAt.Equal(result.CreatedAt))
}

func TestSessionMapperNoRevoke(t *testing.T) {
	s := &domain.Session{
		ID:        "session-1",
		UserID:    "user-1",
		RevokedAt: nil,
		ExpiresAt: time.Now(),
		CreatedAt: time.Now(),
	}

	model := SessionToModel(s)
	result := SessionFromModel(model)

	assert.Nil(t, result.RevokedAt)
}

func TestSessionModelTableName(t *testing.T) {
	m := SessionModel{}
	assert.Equal(t, "sessions", m.TableName())
}

func TestTaskMapperRoundTrip(t *testing.T) {
	dueAt := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)
	createdAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

	tsk := &domain.Task{
		ID:          "task-1",
		Title:       "Tarea 1",
		Description: "Descripcion",
		DueAt:       dueAt,
		Status:      domain.StatusStarted,
		AssigneeID:  "exec-1",
		CreatedBy:   "admin-1",
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	model := TaskToModel(tsk)
	result := TaskFromModel(model)

	assert.Equal(t, tsk.ID, result.ID)
	assert.Equal(t, tsk.Title, result.Title)
	assert.Equal(t, tsk.Description, result.Description)
	assert.True(t, tsk.DueAt.Equal(result.DueAt))
	assert.Equal(t, tsk.Status, result.Status)
	assert.Equal(t, tsk.AssigneeID, result.AssigneeID)
	assert.Equal(t, tsk.CreatedBy, result.CreatedBy)
	assert.True(t, tsk.CreatedAt.Equal(result.CreatedAt))
	assert.True(t, tsk.UpdatedAt.Equal(result.UpdatedAt))
}

func TestTaskModelTableName(t *testing.T) {
	m := TaskModel{}
	assert.Equal(t, "tasks", m.TableName())
}

func TestCommentMapperRoundTrip(t *testing.T) {
	c := &domain.Comment{
		ID:        "comment-1",
		TaskID:    "task-1",
		UserID:    "exec-1",
		Comment:   "Comentario",
		CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	model := CommentToModel(c)
	result := CommentFromModel(model)

	assert.Equal(t, c.ID, result.ID)
	assert.Equal(t, c.TaskID, result.TaskID)
	assert.Equal(t, c.UserID, result.UserID)
	assert.Equal(t, c.Comment, result.Comment)
	assert.True(t, c.CreatedAt.Equal(result.CreatedAt))
}

func TestCommentModelTableName(t *testing.T) {
	m := CommentModel{}
	assert.Equal(t, "task_comments", m.TableName())
}
