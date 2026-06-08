package application

import "errors"

var (
	ErrUserNotFound      = errors.New("usuario no encontrado")
	ErrTaskNotFound       = errors.New("tarea no encontrada")
	ErrSessionNotFound    = errors.New("sesion no encontrada")
	ErrInvalidCredentials = errors.New("credenciales invalidas")
	ErrUnauthorized        = errors.New("no autorizado")
	ErrEmailAlreadyExists = errors.New("el email ya existe")
	ErrCannotCreateAdmin  = errors.New("no se puede crear usuarios ADMIN")
	ErrTaskNotModifiable  = errors.New("la tarea no puede ser modificada en su estado actual")
	ErrTaskOverdue        = errors.New("la tarea esta vencida")
)
