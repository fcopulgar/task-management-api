package dto

import (
	"time"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type CreateTaskInput struct {
	Title       string
	Description string
	DueAt       time.Time
	AssigneeID  string
}

type UpdateTaskInput struct {
	ID          string
	Title       string
	Description string
	DueAt       time.Time
	AssigneeID  string
}

type TransitionTaskInput struct {
	TaskID    string
	UserID    string
	NewStatus domain.TaskStatus
}

type TaskOutput struct {
	ID          string
	Title       string
	Description string
	DueAt       time.Time
	Status      domain.TaskStatus
	AssigneeID  string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func TaskToOutput(t *domain.Task) TaskOutput {
	return TaskOutput{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		DueAt:       t.DueAt,
		Status:      t.Status,
		AssigneeID:  t.AssigneeID,
		CreatedBy:   t.CreatedBy,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
