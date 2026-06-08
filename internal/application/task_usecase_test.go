package application

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/domain"
)

type mockTaskRepo struct {
	createFn         func(ctx context.Context, task *domain.Task) error
	findByIDFn       func(ctx context.Context, id string) (*domain.Task, error)
	listFn           func(ctx context.Context) ([]domain.Task, error)
	listByAssigneeFn func(ctx context.Context, assigneeID string) ([]domain.Task, error)
	updateFn         func(ctx context.Context, task *domain.Task) error
	deleteFn         func(ctx context.Context, id string) error
}

func (m *mockTaskRepo) Create(ctx context.Context, t *domain.Task) error {
	if m.createFn != nil {
		return m.createFn(ctx, t)
	}
	return nil
}
func (m *mockTaskRepo) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}
	return nil, nil
}
func (m *mockTaskRepo) List(ctx context.Context) ([]domain.Task, error) {
	if m.listFn != nil {
		return m.listFn(ctx)
	}
	return nil, nil
}
func (m *mockTaskRepo) ListByAssignee(ctx context.Context, id string) ([]domain.Task, error) {
	if m.listByAssigneeFn != nil {
		return m.listByAssigneeFn(ctx, id)
	}
	return nil, nil
}
func (m *mockTaskRepo) Update(ctx context.Context, t *domain.Task) error {
	if m.updateFn != nil {
		return m.updateFn(ctx, t)
	}
	return nil
}
func (m *mockTaskRepo) Delete(ctx context.Context, id string) error {
	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}
	return nil
}

func TestCreateTaskSuccess(t *testing.T) {
	exec, _ := domain.NewUser("Exec", "exec@test.com", "hash", domain.RoleExecutor)
	exec.ID = "exec-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return exec, nil
		},
	}
	taskRepo := &mockTaskRepo{}

	deps := Dependencies{UserRepo: userRepo, TaskRepo: taskRepo}
	uc := NewTaskUseCase(deps)

	dueAt := time.Now().Add(24 * time.Hour)
	output, err := uc.CreateTask(context.Background(), dto.CreateTaskInput{
		Title:       "Tarea 1",
		Description: "Desc",
		DueAt:       dueAt,
		AssigneeID:  "exec-1",
	}, "admin-1")

	require.NoError(t, err)
	assert.Equal(t, "Tarea 1", output.Title)
	assert.Equal(t, domain.StatusAssigned, output.Status)
	assert.Equal(t, "exec-1", output.AssigneeID)
	assert.Equal(t, "admin-1", output.CreatedBy)
}

func TestCreateTaskAssigneeNotExecutor(t *testing.T) {
	aud, _ := domain.NewUser("Aud", "aud@test.com", "hash", domain.RoleAuditor)
	aud.ID = "aud-1"

	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return aud, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewTaskUseCase(deps)

	dueAt := time.Now().Add(24 * time.Hour)
	_, err := uc.CreateTask(context.Background(), dto.CreateTaskInput{
		Title:      "Tarea",
		DueAt:      dueAt,
		AssigneeID: "aud-1",
	}, "admin-1")

	assert.Error(t, err)
}

func TestCreateTaskAssigneeNotFound(t *testing.T) {
	userRepo := &mockUserRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.User, error) {
			return nil, nil
		},
	}

	deps := Dependencies{UserRepo: userRepo}
	uc := NewTaskUseCase(deps)

	dueAt := time.Now().Add(24 * time.Hour)
	_, err := uc.CreateTask(context.Background(), dto.CreateTaskInput{
		Title:      "Tarea",
		DueAt:      dueAt,
		AssigneeID: "nonexistent",
	}, "admin-1")

	assert.ErrorIs(t, err, ErrUserNotFound)
}

func TestListTasks(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	t1, _ := domain.NewTask("T1", "D1", dueAt, "exec-1", "admin-1")
	t1.ID = "t1"
	t2, _ := domain.NewTask("T2", "D2", dueAt, "exec-2", "admin-1")
	t2.ID = "t2"

	taskRepo := &mockTaskRepo{
		listFn: func(ctx context.Context) ([]domain.Task, error) {
			return []domain.Task{*t1, *t2}, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewTaskUseCase(deps)

	tasks, err := uc.ListTasks(context.Background())
	require.NoError(t, err)
	assert.Len(t, tasks, 2)
}

func TestGetTaskNotFound(t *testing.T) {
	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return nil, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewTaskUseCase(deps)

	_, err := uc.GetTask(context.Background(), "nonexistent")
	assert.ErrorIs(t, err, ErrTaskNotFound)
}

func TestUpdateTaskSuccess(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("Original", "Desc", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewTaskUseCase(deps)

	output, err := uc.UpdateTask(context.Background(), dto.UpdateTaskInput{
		ID:    "task-1",
		Title: "Actualizado",
	})

	require.NoError(t, err)
	assert.Equal(t, "Actualizado", output.Title)
}

func TestUpdateTaskNotAssigned(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"
	task.Status = domain.StatusStarted

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewTaskUseCase(deps)

	_, err := uc.UpdateTask(context.Background(), dto.UpdateTaskInput{
		ID:    "task-1",
		Title: "Actualizado",
	})

	assert.ErrorIs(t, err, ErrTaskNotModifiable)
}

func TestDeleteTaskSuccess(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	deleted := false
	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
		deleteFn: func(ctx context.Context, id string) error {
			deleted = true
			return nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewTaskUseCase(deps)

	err := uc.DeleteTask(context.Background(), "task-1")
	require.NoError(t, err)
	assert.True(t, deleted)
}

func TestDeleteTaskNotAssigned(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"
	task.Status = domain.StatusStarted

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewTaskUseCase(deps)

	err := uc.DeleteTask(context.Background(), "task-1")
	assert.ErrorIs(t, err, ErrTaskNotModifiable)
}
