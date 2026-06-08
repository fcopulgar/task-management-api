package ports

import (
	"context"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user *domain.User) error
}

type SessionRepository interface {
	Create(ctx context.Context, session *domain.Session) error
	FindByID(ctx context.Context, id string) (*domain.Session, error)
	Revoke(ctx context.Context, sessionID string) error
}

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	FindByID(ctx context.Context, id string) (*domain.Task, error)
	List(ctx context.Context) ([]domain.Task, error)
	ListByAssignee(ctx context.Context, assigneeID string) ([]domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id string) error
}

type CommentRepository interface {
	Create(ctx context.Context, comment *domain.Comment) error
	FindByTaskID(ctx context.Context, taskID string) ([]domain.Comment, error)
}
