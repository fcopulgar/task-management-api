package domain

import (
	"errors"
	"time"
)

type Comment struct {
	ID        string
	TaskID    string
	UserID    string
	Comment   string
	CreatedAt time.Time
}

var (
	ErrCommentTaskIDRequired = errors.New("task_id requerido para el comentario")
	ErrCommentUserIDRequired = errors.New("user_id requerido para el comentario")
	ErrCommentTextRequired   = errors.New("texto del comentario requerido")
)

func NewComment(taskID, userID, text string) (*Comment, error) {
	if taskID == "" {
		return nil, ErrCommentTaskIDRequired
	}
	if userID == "" {
		return nil, ErrCommentUserIDRequired
	}
	if text == "" {
		return nil, ErrCommentTextRequired
	}

	return &Comment{
		TaskID:    taskID,
		UserID:    userID,
		Comment:   text,
		CreatedAt: time.Now(),
	}, nil
}
