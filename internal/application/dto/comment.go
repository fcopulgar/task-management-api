package dto

import (
	"time"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type CreateCommentInput struct {
	TaskID  string
	UserID  string
	Comment string
}

type CommentOutput struct {
	ID        string
	TaskID    string
	UserID    string
	Comment   string
	CreatedAt time.Time
}

func CommentToOutput(c *domain.Comment) CommentOutput {
	return CommentOutput{
		ID:        c.ID,
		TaskID:    c.TaskID,
		UserID:    c.UserID,
		Comment:   c.Comment,
		CreatedAt: c.CreatedAt,
	}
}
