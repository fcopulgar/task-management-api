package domain

import "errors"

var (
	ErrInvalidRole       = errors.New("rol invalido")
	ErrUserInactive      = errors.New("usuario inactivo")
	ErrInvalidTransition = errors.New("transicion de estado no permitida")
	ErrTaskTerminal      = errors.New("la tarea ya esta en estado terminal")
	ErrInvalidAssignee   = errors.New("solo usuarios EXECUTOR pueden ser asignados a tareas")
	ErrNotTaskOwner      = errors.New("la tarea no pertenece al usuario")
	ErrCommentOnNonOverdue = errors.New("solo se puede comentar tareas vencidas")
	ErrSessionRevoked    = errors.New("sesion revocada")
	ErrSessionExpired    = errors.New("sesion expirada")
)
