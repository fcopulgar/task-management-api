package application

import (
	"context"
	"fmt"
	"time"

	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/application/ports"
	"github.com/anomalyco/task-management-api/internal/domain"
)

const DefaultTokenDurationHours = 24
const DefaultSessionDuration = 24 * time.Hour

type AuthUseCase struct {
	userRepo    ports.UserRepository
	sessionRepo ports.SessionRepository
	hasher      ports.PasswordHasher
	tokenSvc    ports.TokenService
}

func NewAuthUseCase(deps Dependencies) *AuthUseCase {
	return &AuthUseCase{
		userRepo:    deps.UserRepo,
		sessionRepo: deps.SessionRepo,
		hasher:      deps.Hasher,
		tokenSvc:    deps.TokenSvc,
	}
}

func (uc *AuthUseCase) Login(ctx context.Context, input dto.LoginInput) (*dto.LoginOutput, error) {
	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, fmt.Errorf("error al buscar usuario: %w", err)
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	if err := user.CanLogin(); err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := uc.hasher.Compare(ctx, user.PasswordHash, input.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	session, err := domain.NewSession(user.ID, DefaultSessionDuration)
	if err != nil {
		return nil, fmt.Errorf("error al crear sesion: %w", err)
	}

	if err := uc.sessionRepo.Create(ctx, session); err != nil {
		return nil, fmt.Errorf("error al persistir sesion: %w", err)
	}

	token, err := uc.tokenSvc.Generate(ctx, user.ID, user.Role, session.ID, DefaultTokenDurationHours)
	if err != nil {
		return nil, fmt.Errorf("error al generar token: %w", err)
	}

	return &dto.LoginOutput{Token: token}, nil
}

func (uc *AuthUseCase) Logout(ctx context.Context, sessionID string) error {
	return uc.sessionRepo.Revoke(ctx, sessionID)
}

func (uc *AuthUseCase) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("error al buscar usuario: %w", err)
	}
	if user == nil {
		return ErrUserNotFound
	}

	if err := uc.hasher.Compare(ctx, user.PasswordHash, oldPassword); err != nil {
		return fmt.Errorf("contrasena actual incorrecta: %w", ErrInvalidCredentials)
	}

	hashed, err := uc.hasher.Hash(ctx, newPassword)
	if err != nil {
		return fmt.Errorf("error al hashear contrasena: %w", err)
	}

	user.PasswordHash = hashed
	user.MarkPasswordChanged()

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("error al actualizar usuario: %w", err)
	}

	return nil
}
