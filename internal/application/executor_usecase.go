package application

import (
	"context"
	"fmt"
	"time"

	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/application/ports"
	"github.com/anomalyco/task-management-api/internal/domain"
)

type ExecutorUseCase struct {
	taskRepo    ports.TaskRepository
	commentRepo ports.CommentRepository
}

func NewExecutorUseCase(deps Dependencies) *ExecutorUseCase {
	return &ExecutorUseCase{
		taskRepo:    deps.TaskRepo,
		commentRepo: deps.CommentRepo,
	}
}

func (uc *ExecutorUseCase) ListMyTasks(ctx context.Context, executorID string) ([]dto.TaskOutput, error) {
	tasks, err := uc.taskRepo.ListByAssignee(ctx, executorID)
	if err != nil {
		return nil, fmt.Errorf("error al listar tareas: %w", err)
	}

	outputs := make([]dto.TaskOutput, len(tasks))
	for i, t := range tasks {
		outputs[i] = dto.TaskToOutput(&t)
	}
	return outputs, nil
}

func (uc *ExecutorUseCase) GetMyTask(ctx context.Context, taskID, executorID string) (*dto.TaskOutput, error) {
	task, err := uc.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("error al buscar tarea: %w", err)
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}

	if !task.IsAssignedTo(executorID) {
		return nil, fmt.Errorf("%w", domain.ErrNotTaskOwner)
	}

	output := dto.TaskToOutput(task)
	return &output, nil
}

func (uc *ExecutorUseCase) TransitionTask(ctx context.Context, taskID, executorID string, newStatus domain.TaskStatus) (*dto.TaskOutput, error) {
	task, err := uc.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("error al buscar tarea: %w", err)
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}

	if !task.IsAssignedTo(executorID) {
		return nil, fmt.Errorf("%w", domain.ErrNotTaskOwner)
	}

	now := time.Now()
	if task.IsOverdue(now) {
		return nil, ErrTaskOverdue
	}

	if err := task.TransitionTo(newStatus); err != nil {
		return nil, err
	}

	if err := uc.taskRepo.Update(ctx, task); err != nil {
		return nil, fmt.Errorf("error al actualizar tarea: %w", err)
	}

	output := dto.TaskToOutput(task)
	return &output, nil
}

func (uc *ExecutorUseCase) CommentOnTask(ctx context.Context, taskID, executorID, commentText string) (*dto.CommentOutput, error) {
	task, err := uc.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("error al buscar tarea: %w", err)
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}

	if err := task.CanBeCommentedBy(executorID, time.Now()); err != nil {
		return nil, err
	}

	comment, err := domain.NewComment(taskID, executorID, commentText)
	if err != nil {
		return nil, err
	}

	if err := uc.commentRepo.Create(ctx, comment); err != nil {
		return nil, fmt.Errorf("error al crear comentario: %w", err)
	}

	output := dto.CommentToOutput(comment)
	return &output, nil
}
