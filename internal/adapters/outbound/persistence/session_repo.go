package persistence

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type GormSessionRepository struct {
	db *gorm.DB
}

func NewGormSessionRepository(db *gorm.DB) *GormSessionRepository {
	return &GormSessionRepository{db: db}
}

func (r *GormSessionRepository) Create(ctx context.Context, session *domain.Session) error {
	model := SessionToModel(session)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}
	session.ID = model.ID
	session.CreatedAt = model.CreatedAt
	return nil
}

func (r *GormSessionRepository) FindByID(ctx context.Context, id string) (*domain.Session, error) {
	var model SessionModel
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return SessionFromModel(&model), nil
}

func (r *GormSessionRepository) Revoke(ctx context.Context, sessionID string) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&SessionModel{}).
		Where("id = ? AND revoked_at IS NULL", sessionID).
		Update("revoked_at", now).Error
}
