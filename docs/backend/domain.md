# Dominio backend

## Estado

`PLANIFICADO`

## Usuario

Campos esperados: `id`, `name`, `email`, `password_hash`, `role`, `must_change_password`, `active`, `created_at`, `updated_at`.

Reglas:

- `email` debe ser único.
- `role` debe ser `ADMIN`, `EXECUTOR` o `AUDITOR`.
- `password_hash` nunca debe exponerse por API.
- `must_change_password` indica si el usuario debe cambiar contraseña.
- `active = false` representa usuario desactivado.

## Sesión

Campos esperados: `id`, `user_id`, `revoked_at`, `expires_at`, `created_at`.

Reglas:

- El JWT debe incluir `session_id`.
- Una sesión revocada no debe permitir consumir endpoints protegidos.
- Logout debe marcar la sesión como revocada.

## Tarea

Campos esperados: `id`, `title`, `description`, `due_at`, `status`, `assignee_id`, `created_by`, `created_at`, `updated_at`.

Reglas:

- `assignee_id` debe corresponder a un usuario `EXECUTOR`.
- Una tarea nueva nace como `ASSIGNED`.
- No se puede actualizar ni eliminar si su estado no es `ASSIGNED`.

## Comentario de tarea

Campos esperados: `id`, `task_id`, `user_id`, `comment`, `created_at`.

Reglas:

- Solo ejecutores pueden comentar tareas propias vencidas.
- El comentario debe quedar asociado a la tarea y al usuario.
