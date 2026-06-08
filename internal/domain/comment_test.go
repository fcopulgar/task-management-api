package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewComment(t *testing.T) {
	tests := []struct {
		name    string
		taskID  string
		userID  string
		text    string
		wantErr error
	}{
		{"comentario valido", "task-1", "exec-1", "Comentario de prueba", nil},
		{"taskID vacio", "", "exec-1", "Comentario", ErrCommentTaskIDRequired},
		{"userID vacio", "task-1", "", "Comentario", ErrCommentUserIDRequired},
		{"texto vacio", "task-1", "exec-1", "", ErrCommentTextRequired},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewComment(tt.taskID, tt.userID, tt.text)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, c)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, c)
			assert.Equal(t, tt.taskID, c.TaskID)
			assert.Equal(t, tt.userID, c.UserID)
			assert.Equal(t, tt.text, c.Comment)
			assert.False(t, c.CreatedAt.IsZero())
		})
	}
}
