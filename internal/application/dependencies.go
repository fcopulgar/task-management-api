package application

import "github.com/anomalyco/task-management-api/internal/application/ports"

type Dependencies struct {
	UserRepo    ports.UserRepository
	SessionRepo ports.SessionRepository
	TaskRepo    ports.TaskRepository
	CommentRepo ports.CommentRepository
	Hasher      ports.PasswordHasher
	TokenSvc    ports.TokenService
}
