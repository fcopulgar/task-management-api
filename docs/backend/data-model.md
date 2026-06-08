# Modelo de datos backend

## Estado

`PLANIFICADO`

El modelo de datos esta definido a nivel documental. Los modelos GORM estan `PENDIENTE DE IMPLEMENTACION`.

## Tablas esperadas

### users

- `id`
- `name`
- `email`
- `password_hash`
- `role`
- `must_change_password`
- `active`
- `created_at`
- `updated_at`

Restricciones esperadas:

- `email` unico.
- `role` limitado a `ADMIN`, `EXECUTOR`, `AUDITOR`.

### sessions

- `id`
- `user_id`
- `revoked_at`
- `expires_at`
- `created_at`

Restricciones esperadas:

- `user_id` referencia a usuario.
- `revoked_at` nulo representa sesion no revocada.

### tasks

- `id`
- `title`
- `description`
- `due_at`
- `status`
- `assignee_id`
- `created_by`
- `created_at`
- `updated_at`

Restricciones esperadas:

- `assignee_id` referencia a usuario `EXECUTOR`.
- `created_by` referencia a usuario creador.
- `status` limitado a estados permitidos.

### task_comments

- `id`
- `task_id`
- `user_id`
- `comment`
- `created_at`

Restricciones esperadas:

- `task_id` referencia a tarea.
- `user_id` referencia a usuario comentarista.

## Inicializacion

- Usar GORM `AutoMigrate`.
- No usar migraciones SQL versionadas en esta etapa.
