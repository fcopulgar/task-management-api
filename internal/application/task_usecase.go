package application

import (
	"context"
	"fmt"
	"time"

	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/application/ports"
	"github.com/anomalyco/task-management-api/internal/domain"
)

type TaskUseCase struct {
	taskRepo ports.TaskRepository
	userRepo ports.UserRepository
}

func NewTaskUseCase(deps Dependencies) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: deps.TaskRepo,
		userRepo: deps.UserRepo,
	}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, input dto.CreateTaskInput, createdBy string) (*dto.TaskOutput, error) {
	assignee, err := uc.userRepo.FindByID(ctx, input.AssigneeID)
	if err != nil {
		return nil, fmt.Errorf("error al buscar asignatario: %w", err)
	}
	if assignee == nil {
		return nil, ErrUserNotFound
	}

	if !assignee.IsExecutor() {
		return nil, fmt.Errorf("solo usuarios EXECUTOR pueden ser asignados: %w", domain.ErrInvalidAssignee)
	}

	task, err := domain.NewTask(input.Title, input.Description, input.DueAt, input.AssigneeID, createdBy)
	if err != nil {
		return nil, err
	}

	if err := uc.taskRepo.Create(ctx, task); err != nil {
		return nil, fmt.Errorf("error al crear tarea: %w", err)
	}

	output := dto.TaskToOutput(task)
	return &output, nil
}

func (uc *TaskUseCase) ListTasks(ctx context.Context) ([]dto.TaskOutput, error) {
	tasks, err := uc.taskRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error al listar tareas: %w", err)
	}

	outputs := make([]dto.TaskOutput, len(tasks))
	for i, t := range tasks {
		outputs[i] = dto.TaskToOutput(&t)
	}
	return outputs, nil
}

func (uc *TaskUseCase) GetTask(ctx context.Context, id string) (*dto.TaskOutput, error) {
	task, err := uc.taskRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error al buscar tarea: %w", err)
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}

	output := dto.TaskToOutput(task)
	return &output, nil
}

func (uc *TaskUseCase) UpdateTask(ctx context.Context, input dto.UpdateTaskInput) (*dto.TaskOutput, error) {
	task, err := uc.taskRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("error al buscar tarea: %w", err)
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}

	if !task.CanBeModified() {
		return nil, ErrTaskNotModifiable
	}

	if input.Title != "" {
		task.Title = input.Title
	}
	if input.Description != "" {
		task.Description = input.Description
	}
	if !input.DueAt.IsZero() {
		task.DueAt = input.DueAt
	}
	if input.AssigneeID != "" && input.AssigneeID != task.AssigneeID {
		assignee, err := uc.userRepo.FindByID(ctx, input.AssigneeID)
		if err != nil {
			return nil, fmt.Errorf("error al buscar asignatario: %w", err)
		}
		if assignee == nil || !assignee.IsExecutor() {
			return nil, fmt.Errorf("solo usuarios EXECUTOR pueden ser asignados: %w", domain.ErrInvalidAssignee)
		}
		task.AssigneeID = input.AssigneeID
	}

	task.UpdatedAt = time.Now()

	if err := uc.taskRepo.Update(ctx, task); err != nil {
		return nil, fmt.Errorf("error al actualizar tarea: %w", err)
	}

	output := dto.TaskToOutput(task)
	return &output, nil
}

func (uc *TaskUseCase) DeleteTask(ctx context.Context, id string) error {
	task, err := uc.taskRepo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error al buscar tarea: %w", err)
	}
	if task == nil {
		return ErrTaskNotFound
	}

	if !task.CanBeModified() {
		return ErrTaskNotModifiable
	}

	return uc.taskRepo.Delete(ctx, id)
}
