package persistence

import (
	"context"

	"gorm.io/gorm"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(ctx context.Context, user *domain.User) error {
	model := UserToModel(user)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}
	user.ID = model.ID
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *GormUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var model UserModel
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return UserFromModel(&model), nil
}

func (r *GormUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var model UserModel
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return UserFromModel(&model), nil
}

func (r *GormUserRepository) List(ctx context.Context) ([]domain.User, error) {
	var models []UserModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	users := make([]domain.User, len(models))
	for i, m := range models {
		users[i] = *UserFromModel(&m)
	}
	return users, nil
}

func (r *GormUserRepository) Update(ctx context.Context, user *domain.User) error {
	model := UserToModel(user)
	return r.db.WithContext(ctx).Save(model).Error
}
