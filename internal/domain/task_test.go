package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)

	tests := []struct {
		name        string
		title       string
		description string
		dueAt       time.Time
		assigneeID  string
		createdBy   string
		wantErr     error
	}{
		{"tarea valida", "Tarea 1", "Descripcion", dueAt, "exec-1", "admin-1", nil},
		{"sin titulo", "", "Desc", dueAt, "exec-1", "admin-1", ErrTaskTitleRequired},
		{"sin fecha vencimiento", "Tarea 1", "Desc", time.Time{}, "exec-1", "admin-1", ErrTaskDueAtRequired},
		{"sin asignatario", "Tarea 1", "Desc", dueAt, "", "admin-1", ErrTaskAssigneeRequired},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := NewTask(tt.title, tt.description, tt.dueAt, tt.assigneeID, tt.createdBy)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, task)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, task)
			assert.Equal(t, tt.title, task.Title)
			assert.Equal(t, tt.description, task.Description)
			assert.Equal(t, tt.dueAt, task.DueAt)
			assert.Equal(t, tt.assigneeID, task.AssigneeID)
			assert.Equal(t, tt.createdBy, task.CreatedBy)
			assert.Equal(t, StatusAssigned, task.Status)
			assert.False(t, task.CreatedAt.IsZero())
		})
	}
}

func TestTaskTransitionTo(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")

	t.Run("ASSIGNED -> STARTED valido", func(t *testing.T) {
		err := task.TransitionTo(StatusStarted)
		assert.NoError(t, err)
		assert.Equal(t, StatusStarted, task.Status)
	})

	t.Run("STARTED -> WAITING valido", func(t *testing.T) {
		err := task.TransitionTo(StatusWaiting)
		assert.NoError(t, err)
		assert.Equal(t, StatusWaiting, task.Status)
	})

	t.Run("WAITING -> FINISHED_SUCCESS valido", func(t *testing.T) {
		err := task.TransitionTo(StatusFinishedSuccess)
		assert.NoError(t, err)
		assert.Equal(t, StatusFinishedSuccess, task.Status)
	})
}

func TestTaskInvalidTransitions(t *testing.T) {
	t.Run("ASSIGNED -> WAITING invalido", func(t *testing.T) {
		dueAt := time.Now().Add(24 * time.Hour)
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		err := task.TransitionTo(StatusWaiting)
		assert.ErrorIs(t, err, ErrInvalidTransition)
		assert.Equal(t, StatusAssigned, task.Status)
	})

	t.Run("ASSIGNED -> FINISHED_SUCCESS invalido", func(t *testing.T) {
		dueAt := time.Now().Add(24 * time.Hour)
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		err := task.TransitionTo(StatusFinishedSuccess)
		assert.ErrorIs(t, err, ErrInvalidTransition)
	})

	t.Run("STARTED -> ASSIGNED invalido", func(t *testing.T) {
		dueAt := time.Now().Add(24 * time.Hour)
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		task.Status = StatusStarted
		err := task.TransitionTo(StatusAssigned)
		assert.ErrorIs(t, err, ErrInvalidTransition)
	})

	t.Run("estado terminal no permite transiciones", func(t *testing.T) {
		dueAt := time.Now().Add(24 * time.Hour)
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		task.Status = StatusFinishedSuccess
		err := task.TransitionTo(StatusWaiting)
		assert.ErrorIs(t, err, ErrTaskTerminal)
	})

	t.Run("FINISHED_ERROR no permite transiciones", func(t *testing.T) {
		dueAt := time.Now().Add(24 * time.Hour)
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		task.Status = StatusFinishedError
		err := task.TransitionTo(StatusWaiting)
		assert.ErrorIs(t, err, ErrTaskTerminal)
	})
}

func TestTaskIsOverdue(t *testing.T) {
	future := time.Now().Add(24 * time.Hour)
	past := time.Now().Add(-1 * time.Hour)

	t.Run("tarea no vencida", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", future, "exec-1", "admin-1")
		assert.False(t, task.IsOverdue(time.Now()))
	})

	t.Run("tarea vencida", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", past, "exec-1", "admin-1")
		assert.True(t, task.IsOverdue(time.Now()))
	})
}

func TestTaskCanBeModified(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)

	t.Run("ASSIGNED permite modificacion", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		assert.True(t, task.CanBeModified())
	})

	t.Run("STARTED no permite modificacion", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		task.Status = StatusStarted
		assert.False(t, task.CanBeModified())
	})

	t.Run("FINISHED_SUCCESS no permite modificacion", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		task.Status = StatusFinishedSuccess
		assert.False(t, task.CanBeModified())
	})
}

func TestTaskIsAssignedTo(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")

	assert.True(t, task.IsAssignedTo("exec-1"))
	assert.False(t, task.IsAssignedTo("exec-2"))
}

func TestTaskAssignTo(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)

	t.Run("asignar a executor valido", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		err := task.AssignTo("exec-2", RoleExecutor)
		assert.NoError(t, err)
		assert.Equal(t, "exec-2", task.AssigneeID)
	})

	t.Run("asignar a admin rechazado", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		err := task.AssignTo("admin-1", RoleAdmin)
		assert.ErrorIs(t, err, ErrInvalidAssignee)
		assert.Equal(t, "exec-1", task.AssigneeID)
	})

	t.Run("asignar a auditor rechazado", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		err := task.AssignTo("aud-1", RoleAuditor)
		assert.ErrorIs(t, err, ErrInvalidAssignee)
		assert.Equal(t, "exec-1", task.AssigneeID)
	})

	t.Run("asignar id vacio rechazado", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")
		err := task.AssignTo("", RoleExecutor)
		assert.ErrorIs(t, err, ErrTaskAssigneeRequired)
	})
}

func TestTaskCanBeCommentedBy(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	future := time.Now().Add(24 * time.Hour)
	now := time.Now()

	t.Run("ejecutor puede comentar tarea propia vencida", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", past, "exec-1", "admin-1")
		err := task.CanBeCommentedBy("exec-1", now)
		assert.NoError(t, err)
	})

	t.Run("no se puede comentar tarea ajena", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", past, "exec-1", "admin-1")
		err := task.CanBeCommentedBy("exec-2", now)
		assert.ErrorIs(t, err, ErrNotTaskOwner)
	})

	t.Run("no se puede comentar tarea no vencida", func(t *testing.T) {
		task, _ := NewTask("Tarea", "Desc", future, "exec-1", "admin-1")
		err := task.CanBeCommentedBy("exec-1", now)
		assert.ErrorIs(t, err, ErrCommentOnNonOverdue)
	})
}

func TestTaskIsAssigned(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")

	assert.True(t, task.IsAssigned())

	task.Status = StatusStarted
	assert.False(t, task.IsAssigned())
}

func TestTaskStatusIsValid(t *testing.T) {
	assert.True(t, StatusAssigned.IsValid())
	assert.True(t, StatusStarted.IsValid())
	assert.True(t, StatusWaiting.IsValid())
	assert.True(t, StatusFinishedSuccess.IsValid())
	assert.True(t, StatusFinishedError.IsValid())
	assert.False(t, TaskStatus("INVALID").IsValid())
}

func TestTaskStatusCanTransitionTo(t *testing.T) {
	assert.True(t, StatusAssigned.CanTransitionTo(StatusStarted))
	assert.False(t, StatusAssigned.CanTransitionTo(StatusWaiting))
	assert.False(t, StatusAssigned.CanTransitionTo(StatusFinishedSuccess))

	assert.True(t, StatusStarted.CanTransitionTo(StatusWaiting))
	assert.True(t, StatusStarted.CanTransitionTo(StatusFinishedSuccess))
	assert.True(t, StatusStarted.CanTransitionTo(StatusFinishedError))
	assert.False(t, StatusStarted.CanTransitionTo(StatusAssigned))

	assert.True(t, StatusWaiting.CanTransitionTo(StatusWaiting))
	assert.True(t, StatusWaiting.CanTransitionTo(StatusFinishedSuccess))
	assert.True(t, StatusWaiting.CanTransitionTo(StatusFinishedError))
	assert.False(t, StatusWaiting.CanTransitionTo(StatusStarted))

	assert.False(t, StatusFinishedSuccess.CanTransitionTo(StatusWaiting))
	assert.False(t, StatusFinishedError.CanTransitionTo(StatusAssigned))
}

func TestTaskStatusIsTerminal(t *testing.T) {
	assert.False(t, StatusAssigned.IsTerminal())
	assert.False(t, StatusStarted.IsTerminal())
	assert.False(t, StatusWaiting.IsTerminal())
	assert.True(t, StatusFinishedSuccess.IsTerminal())
	assert.True(t, StatusFinishedError.IsTerminal())
}

func TestTaskStatusString(t *testing.T) {
	assert.Equal(t, "ASSIGNED", StatusAssigned.String())
	assert.Equal(t, "STARTED", StatusStarted.String())
	assert.Equal(t, "WAITING", StatusWaiting.String())
	assert.Equal(t, "FINISHED_SUCCESS", StatusFinishedSuccess.String())
	assert.Equal(t, "FINISHED_ERROR", StatusFinishedError.String())
}

func TestTaskTransitionUpdatesTimestamp(t *testing.T) {
	dueAt := time.Now().Add(24 * time.Hour)
	task, _ := NewTask("Tarea", "Desc", dueAt, "exec-1", "admin-1")

	oldUpdatedAt := task.UpdatedAt
	time.Sleep(time.Millisecond)

	task.TransitionTo(StatusStarted)
	assert.True(t, task.UpdatedAt.After(oldUpdatedAt))
}
