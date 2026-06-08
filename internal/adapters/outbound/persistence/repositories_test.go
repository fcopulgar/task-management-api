//go:build integration

package persistence

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anomalyco/task-management-api/internal/domain"
)

func setupTestDB(t *testing.T) Config {
	t.Helper()

	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbName := "taskmanagement"
	sslMode := "disable"

	if h := os.Getenv("DB_HOST"); h != "" {
		host = h
	}
	if p := os.Getenv("DB_PORT"); p != "" {
		port = p
	}
	if u := os.Getenv("DB_USER"); u != "" {
		user = u
	}
	if pw := os.Getenv("DB_PASSWORD"); pw != "" {
		password = pw
	}
	if d := os.Getenv("DB_NAME"); d != "" {
		dbName = d
	}

	return Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
		SSLMode:  sslMode,
	}
}

func TestGormUserRepositoryCRUD(t *testing.T) {
	cfg := setupTestDB(t)
	db, err := NewConnection(cfg)
	require.NoError(t, err)

	err = RunMigrations(db)
	require.NoError(t, err)

	repo := NewGormUserRepository(db)
	ctx := context.Background()

	u, err := domain.NewUser("Test User", "test@example.com", "hashed", domain.RoleExecutor)
	require.NoError(t, err)

	err = repo.Create(ctx, u)
	require.NoError(t, err)
	assert.NotEmpty(t, u.ID)

	found, err := repo.FindByID(ctx, u.ID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "Test User", found.Name)
	assert.Equal(t, "test@example.com", found.Email)

	foundByEmail, err := repo.FindByEmail(ctx, "test@example.com")
	require.NoError(t, err)
	require.NotNil(t, foundByEmail)
	assert.Equal(t, u.ID, foundByEmail.ID)

	users, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, users, 1)

	found.Name = "Updated User"
	err = repo.Update(ctx, found)
	require.NoError(t, err)

	updated, err := repo.FindByID(ctx, u.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated User", updated.Name)

	notFound, err := repo.FindByID(ctx, "00000000-0000-0000-0000-000000000000")
	require.NoError(t, err)
	assert.Nil(t, notFound)

	db.Exec("DELETE FROM task_comments")
	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM sessions")
	db.Exec("DELETE FROM users")
}

func TestGormSessionRepository(t *testing.T) {
	cfg := setupTestDB(t)
	db, err := NewConnection(cfg)
	require.NoError(t, err)

	err = RunMigrations(db)
	require.NoError(t, err)

	repo := NewGormSessionRepository(db)
	ctx := context.Background()

	s, err := domain.NewSession("user-1", 24*time.Hour)
	require.NoError(t, err)

	err = repo.Create(ctx, s)
	require.NoError(t, err)
	assert.NotEmpty(t, s.ID)

	found, err := repo.FindByID(ctx, s.ID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "user-1", found.UserID)
	assert.False(t, found.IsRevoked())

	err = repo.Revoke(ctx, s.ID)
	require.NoError(t, err)

	revoked, err := repo.FindByID(ctx, s.ID)
	require.NoError(t, err)
	assert.True(t, revoked.IsRevoked())

	db.Exec("DELETE FROM task_comments")
	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM sessions")
	db.Exec("DELETE FROM users")
}

func TestGormTaskRepository(t *testing.T) {
	cfg := setupTestDB(t)
	db, err := NewConnection(cfg)
	require.NoError(t, err)

	err = RunMigrations(db)
	require.NoError(t, err)

	repo := NewGormTaskRepository(db)
	ctx := context.Background()

	dueAt := time.Now().Add(24 * time.Hour)
	task, err := domain.NewTask("Tarea Test", "Descripcion", dueAt, "exec-1", "admin-1")
	require.NoError(t, err)

	err = repo.Create(ctx, task)
	require.NoError(t, err)
	assert.NotEmpty(t, task.ID)

	found, err := repo.FindByID(ctx, task.ID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, "Tarea Test", found.Title)

	tasks, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, tasks, 1)

	byAssignee, err := repo.ListByAssignee(ctx, "exec-1")
	require.NoError(t, err)
	assert.Len(t, byAssignee, 1)

	empty, err := repo.ListByAssignee(ctx, "exec-2")
	require.NoError(t, err)
	assert.Len(t, empty, 0)

	found.Title = "Updated Task"
	err = repo.Update(ctx, found)
	require.NoError(t, err)

	updated, err := repo.FindByID(ctx, task.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated Task", updated.Title)

	err = repo.Delete(ctx, task.ID)
	require.NoError(t, err)

	deleted, err := repo.FindByID(ctx, task.ID)
	require.NoError(t, err)
	assert.Nil(t, deleted)

	db.Exec("DELETE FROM task_comments")
	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM sessions")
	db.Exec("DELETE FROM users")
}

func TestGormCommentRepository(t *testing.T) {
	cfg := setupTestDB(t)
	db, err := NewConnection(cfg)
	require.NoError(t, err)

	err = RunMigrations(db)
	require.NoError(t, err)

	repo := NewGormCommentRepository(db)
	ctx := context.Background()

	c, err := domain.NewComment("task-1", "exec-1", "Comentario de prueba")
	require.NoError(t, err)

	err = repo.Create(ctx, c)
	require.NoError(t, err)
	assert.NotEmpty(t, c.ID)

	comments, err := repo.FindByTaskID(ctx, "task-1")
	require.NoError(t, err)
	assert.Len(t, comments, 1)
	assert.Equal(t, "Comentario de prueba", comments[0].Comment)

	empty, err := repo.FindByTaskID(ctx, "task-999")
	require.NoError(t, err)
	assert.Len(t, empty, 0)

	db.Exec("DELETE FROM task_comments")
	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM sessions")
	db.Exec("DELETE FROM users")
}
