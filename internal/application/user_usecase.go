package application

import (
	"context"
	"fmt"

	"github.com/anomalyco/task-management-api/internal/application/dto"
	"github.com/anomalyco/task-management-api/internal/application/ports"
	"github.com/anomalyco/task-management-api/internal/domain"
)

type UserUseCase struct {
	userRepo ports.UserRepository
	hasher   ports.PasswordHasher
}

func NewUserUseCase(deps Dependencies) *UserUseCase {
	return &UserUseCase{
		userRepo: deps.UserRepo,
		hasher:   deps.Hasher,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, input dto.CreateUserInput) (*dto.UserOutput, error) {
	if input.Role == domain.RoleAdmin {
		return nil, ErrCannotCreateAdmin
	}

	existing, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, fmt.Errorf("error al verificar email: %w", err)
	}
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	hashedPassword, err := uc.hasher.Hash(ctx, input.Password)
	if err != nil {
		return nil, fmt.Errorf("error al hashear contrasena: %w", err)
	}

	user, err := domain.NewUser(input.Name, input.Email, hashedPassword, input.Role)
	if err != nil {
		return nil, err
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("error al crear usuario: %w", err)
	}

	output := dto.UserToOutput(user)
	return &output, nil
}

func (uc *UserUseCase) ListUsers(ctx context.Context) ([]dto.UserOutput, error) {
	users, err := uc.userRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error al listar usuarios: %w", err)
	}

	outputs := make([]dto.UserOutput, len(users))
	for i, u := range users {
		outputs[i] = dto.UserToOutput(&u)
	}
	return outputs, nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (*dto.UserOutput, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error al buscar usuario: %w", err)
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	output := dto.UserToOutput(user)
	return &output, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, input dto.UpdateUserInput) (*dto.UserOutput, error) {
	user, err := uc.userRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("error al buscar usuario: %w", err)
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" && input.Email != user.Email {
		existing, err := uc.userRepo.FindByEmail(ctx, input.Email)
		if err != nil {
			return nil, fmt.Errorf("error al verificar email: %w", err)
		}
		if existing != nil && existing.ID != user.ID {
			return nil, ErrEmailAlreadyExists
		}
		user.Email = input.Email
	}
	if input.Role.IsValid() {
		user.Role = input.Role
	}
	if input.Active != nil {
		if *input.Active {
			user.Activate()
		} else {
			user.Deactivate()
		}
	}

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("error al actualizar usuario: %w", err)
	}

	output := dto.UserToOutput(user)
	return &output, nil
}

func (uc *UserUseCase) DeactivateUser(ctx context.Context, id string) error {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error al buscar usuario: %w", err)
	}
	if user == nil {
		return ErrUserNotFound
	}

	user.Deactivate()

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("error al desactivar usuario: %w", err)
	}

	return nil
}
