package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationErrors(t *testing.T) {
	assert.EqualError(t, ErrUserNotFound, "usuario no encontrado")
	assert.EqualError(t, ErrTaskNotFound, "tarea no encontrada")
	assert.EqualError(t, ErrSessionNotFound, "sesion no encontrada")
	assert.EqualError(t, ErrInvalidCredentials, "credenciales invalidas")
	assert.EqualError(t, ErrUnauthorized, "no autorizado")
	assert.EqualError(t, ErrEmailAlreadyExists, "el email ya existe")
	assert.EqualError(t, ErrCannotCreateAdmin, "no se puede crear usuarios ADMIN")
	assert.EqualError(t, ErrTaskNotModifiable, "la tarea no puede ser modificada en su estado actual")
	assert.EqualError(t, ErrTaskOverdue, "la tarea esta vencida")
}
