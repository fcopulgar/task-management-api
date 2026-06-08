package persistence

import (
	"context"

	"gorm.io/gorm"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type GormCommentRepository struct {
	db *gorm.DB
}

func NewGormCommentRepository(db *gorm.DB) *GormCommentRepository {
	return &GormCommentRepository{db: db}
}

func (r *GormCommentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	model := CommentToModel(comment)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}
	comment.ID = model.ID
	comment.CreatedAt = model.CreatedAt
	return nil
}

func (r *GormCommentRepository) FindByTaskID(ctx context.Context, taskID string) ([]domain.Comment, error) {
	var models []CommentModel
	if err := r.db.WithContext(ctx).Where("task_id = ?", taskID).Find(&models).Error; err != nil {
		return nil, err
	}
	comments := make([]domain.Comment, len(models))
	for i, m := range models {
		comments[i] = *CommentFromModel(&m)
	}
	return comments, nil
}
