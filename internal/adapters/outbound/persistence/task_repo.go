package persistence

import (
	"context"

	"gorm.io/gorm"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository {
	return &GormTaskRepository{db: db}
}

func (r *GormTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	model := TaskToModel(task)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}
	task.ID = model.ID
	task.CreatedAt = model.CreatedAt
	task.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *GormTaskRepository) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	var model TaskModel
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return TaskFromModel(&model), nil
}

func (r *GormTaskRepository) List(ctx context.Context) ([]domain.Task, error) {
	var models []TaskModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	tasks := make([]domain.Task, len(models))
	for i, m := range models {
		tasks[i] = *TaskFromModel(&m)
	}
	return tasks, nil
}

func (r *GormTaskRepository) ListByAssignee(ctx context.Context, assigneeID string) ([]domain.Task, error) {
	var models []TaskModel
	if err := r.db.WithContext(ctx).Where("assignee_id = ?", assigneeID).Find(&models).Error; err != nil {
		return nil, err
	}
	tasks := make([]domain.Task, len(models))
	for i, m := range models {
		tasks[i] = *TaskFromModel(&m)
	}
	return tasks, nil
}

func (r *GormTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	model := TaskToModel(task)
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *GormTaskRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&TaskModel{}, "id = ?", id).Error
}
