package application

import "github.com/anomalyco/task-management-api/internal/application/ports"

type Dependencies struct {
	UserRepo    ports.UserRepository
	SessionRepo ports.SessionRepository
	Hasher      ports.PasswordHasher
	TokenSvc    ports.TokenService
}
