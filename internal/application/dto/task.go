package dto

import (
	"time"

	"github.com/anomalyco/task-management-api/internal/domain"
)

type CreateTaskInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueAt       time.Time `json:"due_at"`
	AssigneeID  string    `json:"assignee_id"`
}

type UpdateTaskInput struct {
	ID          string    `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueAt       time.Time `json:"due_at"`
	AssigneeID  string    `json:"assignee_id"`
}

type TransitionTaskInput struct {
	TaskID    string           `json:"-"`
	UserID    string           `json:"-"`
	NewStatus domain.TaskStatus `json:"new_status"`
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
