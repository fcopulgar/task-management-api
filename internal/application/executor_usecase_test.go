package application

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type mockCommentRepo struct {
	createFn       func(ctx context.Context, comment *domain.Comment) error
	findByTaskIDFn func(ctx context.Context, taskID string) ([]domain.Comment, error)
}

func (m *mockCommentRepo) Create(ctx context.Context, c *domain.Comment) error {
	if m.createFn != nil {
		return m.createFn(ctx, c)
	}
	return nil
}
func (m *mockCommentRepo) FindByTaskID(ctx context.Context, taskID string) ([]domain.Comment, error) {
	if m.findByTaskIDFn != nil {
		return m.findByTaskIDFn(ctx, taskID)
	}
	return nil, nil
}

func TestListMyTasks(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	t1, _ := domain.NewTask("T1", "D1", dueAt, "exec-1", "admin-1")
	t1.ID = "t1"
	t2, _ := domain.NewTask("T2", "D2", dueAt, "exec-1", "admin-1")
	t2.ID = "t2"

	taskRepo := &mockTaskRepo{
		listByAssigneeFn: func(ctx context.Context, assigneeID string) ([]domain.Task, error) {
			assert.Equal(t, "exec-1", assigneeID)
			return []domain.Task{*t1, *t2}, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewExecutorUseCase(deps)

	tasks, err := uc.ListMyTasks(context.Background(), "exec-1")
	require.NoError(t, err)
	assert.Len(t, tasks, 2)
}

func TestGetMyTaskSuccess(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewExecutorUseCase(deps)

	result, err := uc.GetMyTask(context.Background(), "task-1", "exec-1")
	require.NoError(t, err)
	assert.Equal(t, "T", result.Title)
}

func TestGetMyTaskNotOwner(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewExecutorUseCase(deps)

	_, err := uc.GetMyTask(context.Background(), "task-1", "exec-2")
	assert.ErrorIs(t, err, domain.ErrNotTaskOwner)
}

func TestTransitionTaskSuccess(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewExecutorUseCase(deps)

	result, err := uc.TransitionTask(context.Background(), "task-1", "exec-1", domain.StatusStarted)
	require.NoError(t, err)
	assert.Equal(t, domain.StatusStarted, result.Status)
}

func TestTransitionTaskNotOwner(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewExecutorUseCase(deps)

	_, err := uc.TransitionTask(context.Background(), "task-1", "exec-2", domain.StatusStarted)
	assert.ErrorIs(t, err, domain.ErrNotTaskOwner)
}

func TestTransitionTaskOverdue(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	task, _ := domain.NewTask("T", "D", past, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewExecutorUseCase(deps)

	_, err := uc.TransitionTask(context.Background(), "task-1", "exec-1", domain.StatusStarted)
	assert.ErrorIs(t, err, ErrTaskOverdue)
}

func TestTransitionTaskInvalid(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewExecutorUseCase(deps)

	_, err := uc.TransitionTask(context.Background(), "task-1", "exec-1", domain.StatusFinishedSuccess)
	assert.ErrorIs(t, err, domain.ErrInvalidTransition)
}

func TestCommentOnTaskSuccess(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	task, _ := domain.NewTask("T", "D", past, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}
	commentRepo := &mockCommentRepo{}

	deps := Dependencies{TaskRepo: taskRepo, CommentRepo: commentRepo}
	uc := NewExecutorUseCase(deps)

	result, err := uc.CommentOnTask(context.Background(), "task-1", "exec-1", "Comentario")
	require.NoError(t, err)
	assert.Equal(t, "Comentario", result.Comment)
	assert.Equal(t, "task-1", result.TaskID)
}

func TestCommentOnTaskNotOverdue(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := domain.NewTask("T", "D", dueAt, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewExecutorUseCase(deps)

	_, err := uc.CommentOnTask(context.Background(), "task-1", "exec-1", "Comentario")
	assert.ErrorIs(t, err, domain.ErrCommentOnNonOverdue)
}

func TestCommentOnTaskNotOwner(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	task, _ := domain.NewTask("T", "D", past, "exec-1", "admin-1")
	task.ID = "task-1"

	taskRepo := &mockTaskRepo{
		findByIDFn: func(ctx context.Context, id string) (*domain.Task, error) {
			return task, nil
		},
	}

	deps := Dependencies{TaskRepo: taskRepo}
	uc := NewExecutorUseCase(deps)

	_, err := uc.CommentOnTask(context.Background(), "task-1", "exec-2", "Comentario")
	assert.ErrorIs(t, err, domain.ErrNotTaskOwner)
}
