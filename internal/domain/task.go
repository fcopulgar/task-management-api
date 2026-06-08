package domain

import (
	"errors"
	"time"
)

type Task struct {
	ID          string
	Title       string
	Description string
	DueAt       time.Time
	Status      TaskStatus
	AssigneeID  string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var (
	ErrTaskTitleRequired    = errors.New("titulo de tarea requerido")
	ErrTaskDueAtRequired    = errors.New("fecha de vencimiento requerida")
	ErrTaskAssigneeRequired = errors.New("asignatario requerido")
)

func NewTask(title, description string, dueAt time.Time, assigneeID, createdBy string) (*Task, error) {
	if title == "" {
		return nil, ErrTaskTitleRequired
	}
	if dueAt.IsZero() {
		return nil, ErrTaskDueAtRequired
	}
	if assigneeID == "" {
		return nil, ErrTaskAssigneeRequired
	}

	now := time.Now()
	return &Task{
		Title:       title,
		Description: description,
		DueAt:       dueAt,
		Status:      StatusAssigned,
		AssigneeID:  assigneeID,
		CreatedBy:   createdBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (t *Task) IsOverdue(now time.Time) bool {
	return now.After(t.DueAt)
}

func (t *Task) CanBeModified() bool {
	return t.Status == StatusAssigned
}

func (t *Task) IsAssignedTo(userID string) bool {
	return t.AssigneeID == userID
}

func (t *Task) IsAssigned() bool {
	return t.Status == StatusAssigned
}

func (t *Task) TransitionTo(target TaskStatus) error {
	if t.Status.IsTerminal() {
		return ErrTaskTerminal
	}
	if !t.Status.CanTransitionTo(target) {
		return ErrInvalidTransition
	}

	t.Status = target
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) AssignTo(assigneeID string, assigneeRole Role) error {
	if assigneeRole != RoleExecutor {
		return ErrInvalidAssignee
	}
	if assigneeID == "" {
		return ErrTaskAssigneeRequired
	}
	t.AssigneeID = assigneeID
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) CanBeCommentedBy(userID string, now time.Time) error {
	if !t.IsAssignedTo(userID) {
		return ErrNotTaskOwner
	}
	if !t.IsOverdue(now) {
		return ErrCommentOnNonOverdue
	}
	return nil
}
