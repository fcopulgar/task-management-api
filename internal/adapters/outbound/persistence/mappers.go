package persistence

import "github.com/anomalyco/task-management-api/internal/domain"

func UserToModel(u *domain.User) *UserModel {
	return &UserModel{
		ID:                 u.ID,
		Name:               u.Name,
		Email:              u.Email,
		PasswordHash:       u.PasswordHash,
		Role:               string(u.Role),
		MustChangePassword: u.MustChangePassword,
		Active:             u.Active,
		CreatedAt:          u.CreatedAt,
		UpdatedAt:          u.UpdatedAt,
	}
}

func UserFromModel(m *UserModel) *domain.User {
	return &domain.User{
		ID:                 m.ID,
		Name:               m.Name,
		Email:              m.Email,
		PasswordHash:       m.PasswordHash,
		Role:               domain.Role(m.Role),
		MustChangePassword: m.MustChangePassword,
		Active:             m.Active,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
	}
}

func SessionToModel(s *domain.Session) *SessionModel {
	return &SessionModel{
		ID:        s.ID,
		UserID:    s.UserID,
		RevokedAt: s.RevokedAt,
		ExpiresAt: s.ExpiresAt,
		CreatedAt: s.CreatedAt,
	}
}

func SessionFromModel(m *SessionModel) *domain.Session {
	return &domain.Session{
		ID:        m.ID,
		UserID:    m.UserID,
		RevokedAt: m.RevokedAt,
		ExpiresAt: m.ExpiresAt,
		CreatedAt: m.CreatedAt,
	}
}

func TaskToModel(t *domain.Task) *TaskModel {
	return &TaskModel{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		DueAt:       t.DueAt,
		Status:      string(t.Status),
		AssigneeID:  t.AssigneeID,
		CreatedBy:   t.CreatedBy,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func TaskFromModel(m *TaskModel) *domain.Task {
	return &domain.Task{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		DueAt:       m.DueAt,
		Status:      domain.TaskStatus(m.Status),
		AssigneeID:  m.AssigneeID,
		CreatedBy:   m.CreatedBy,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func CommentToModel(c *domain.Comment) *CommentModel {
	return &CommentModel{
		ID:        c.ID,
		TaskID:    c.TaskID,
		UserID:    c.UserID,
		Comment:   c.Comment,
		CreatedAt: c.CreatedAt,
	}
}

func CommentFromModel(m *CommentModel) *domain.Comment {
	return &domain.Comment{
		ID:        m.ID,
		TaskID:    m.TaskID,
		UserID:    m.UserID,
		Comment:   m.Comment,
		CreatedAt: m.CreatedAt,
	}
}
