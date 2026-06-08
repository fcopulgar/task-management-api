package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainErrors(t *testing.T) {
	assert.EqualError(t, ErrInvalidRole, "rol invalido")
	assert.EqualError(t, ErrUserInactive, "usuario inactivo")
	assert.EqualError(t, ErrInvalidTransition, "transicion de estado no permitida")
	assert.EqualError(t, ErrTaskTerminal, "la tarea ya esta en estado terminal")
	assert.EqualError(t, ErrInvalidAssignee, "solo usuarios EXECUTOR pueden ser asignados a tareas")
	assert.EqualError(t, ErrNotTaskOwner, "la tarea no pertenece al usuario")
	assert.EqualError(t, ErrCommentOnNonOverdue, "solo se puede comentar tareas vencidas")
	assert.EqualError(t, ErrSessionRevoked, "sesion revocada")
	assert.EqualError(t, ErrSessionExpired, "sesion expirada")
}
