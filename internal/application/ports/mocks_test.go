package ports

import (
	"context"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type MockUserRepository struct {
	CreateFn     func(ctx context.Context, user *domain.User) error
	FindByIDFn   func(ctx context.Context, id string) (*domain.User, error)
	FindByEmailFn func(ctx context.Context, email string) (*domain.User, error)
	ListFn       func(ctx context.Context) ([]domain.User, error)
	UpdateFn     func(ctx context.Context, user *domain.User) error
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, user)
	}
	return nil
}

func (m *MockUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	if m.FindByIDFn != nil {
		return m.FindByIDFn(ctx, id)
	}
	return nil, nil
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if m.FindByEmailFn != nil {
		return m.FindByEmailFn(ctx, email)
	}
	return nil, nil
}

func (m *MockUserRepository) List(ctx context.Context) ([]domain.User, error) {
	if m.ListFn != nil {
		return m.ListFn(ctx)
	}
	return nil, nil
}

func (m *MockUserRepository) Update(ctx context.Context, user *domain.User) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, user)
	}
	return nil
}

type MockSessionRepository struct {
	CreateFn   func(ctx context.Context, session *domain.Session) error
	FindByIDFn func(ctx context.Context, id string) (*domain.Session, error)
	RevokeFn   func(ctx context.Context, sessionID string) error
}

func (m *MockSessionRepository) Create(ctx context.Context, session *domain.Session) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, session)
	}
	return nil
}

func (m *MockSessionRepository) FindByID(ctx context.Context, id string) (*domain.Session, error) {
	if m.FindByIDFn != nil {
		return m.FindByIDFn(ctx, id)
	}
	return nil, nil
}

func (m *MockSessionRepository) Revoke(ctx context.Context, sessionID string) error {
	if m.RevokeFn != nil {
		return m.RevokeFn(ctx, sessionID)
	}
	return nil
}

type MockTaskRepository struct {
	CreateFn         func(ctx context.Context, task *domain.Task) error
	FindByIDFn       func(ctx context.Context, id string) (*domain.Task, error)
	ListFn           func(ctx context.Context) ([]domain.Task, error)
	ListByAssigneeFn func(ctx context.Context, assigneeID string) ([]domain.Task, error)
	UpdateFn         func(ctx context.Context, task *domain.Task) error
	DeleteFn         func(ctx context.Context, id string) error
}

func (m *MockTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, task)
	}
	return nil
}

func (m *MockTaskRepository) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	if m.FindByIDFn != nil {
		return m.FindByIDFn(ctx, id)
	}
	return nil, nil
}

func (m *MockTaskRepository) List(ctx context.Context) ([]domain.Task, error) {
	if m.ListFn != nil {
		return m.ListFn(ctx)
	}
	return nil, nil
}

func (m *MockTaskRepository) ListByAssignee(ctx context.Context, assigneeID string) ([]domain.Task, error) {
	if m.ListByAssigneeFn != nil {
		return m.ListByAssigneeFn(ctx, assigneeID)
	}
	return nil, nil
}

func (m *MockTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, task)
	}
	return nil
}

func (m *MockTaskRepository) Delete(ctx context.Context, id string) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, id)
	}
	return nil
}

type MockCommentRepository struct {
	CreateFn       func(ctx context.Context, comment *domain.Comment) error
	FindByTaskIDFn func(ctx context.Context, taskID string) ([]domain.Comment, error)
}

func (m *MockCommentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, comment)
	}
	return nil
}

func (m *MockCommentRepository) FindByTaskID(ctx context.Context, taskID string) ([]domain.Comment, error) {
	if m.FindByTaskIDFn != nil {
		return m.FindByTaskIDFn(ctx, taskID)
	}
	return nil, nil
}

type MockPasswordHasher struct {
	HashFn    func(ctx context.Context, password string) (string, error)
	CompareFn func(ctx context.Context, hash, password string) error
}

func (m *MockPasswordHasher) Hash(ctx context.Context, password string) (string, error) {
	if m.HashFn != nil {
		return m.HashFn(ctx, password)
	}
	return "", nil
}

func (m *MockPasswordHasher) Compare(ctx context.Context, hash, password string) error {
	if m.CompareFn != nil {
		return m.CompareFn(ctx, hash, password)
	}
	return nil
}

type MockTokenService struct {
	GenerateFn func(ctx context.Context, userID string, role domain.Role, sessionID string, tokenDurationHours int) (string, error)
	ValidateFn func(ctx context.Context, tokenString string) (*TokenClaims, error)
}

func (m *MockTokenService) Generate(ctx context.Context, userID string, role domain.Role, sessionID string, tokenDurationHours int) (string, error) {
	if m.GenerateFn != nil {
		return m.GenerateFn(ctx, userID, role, sessionID, tokenDurationHours)
	}
	return "", nil
}

func (m *MockTokenService) Validate(ctx context.Context, tokenString string) (*TokenClaims, error) {
	if m.ValidateFn != nil {
		return m.ValidateFn(ctx, tokenString)
	}
	return nil, nil
}
