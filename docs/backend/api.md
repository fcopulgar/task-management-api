# API backend

## Estado

`IMPLEMENTADO`

Todos los endpoints listados a continuacion estan implementados y funcionales.

## Auth

```
POST /auth/login
POST /auth/logout
POST /auth/password
```

### POST /auth/login

Autentica un usuario activo y retorna un JWT con `session_id`.

**Request:**
```json
{"email": "user@test.com", "password": "secret"}
```

**Response 200:**
```json
{"token": "eyJhbGciOiJIUzI1NiIs..."}
```

**Errores:**
- `401` — credenciales invalidas (usuario no encontrado, inactivo o contraseña incorrecta)

### POST /auth/logout

Cierra la sesión actual revocandola en base de datos.

**Requiere:** `Authorization: Bearer <token>`

**Response 200:**
```json
{"message": "sesion cerrada"}
```

### POST /auth/password

Cambia la contraseña del usuario autenticado.

**Requiere:** `Authorization: Bearer <token>`

**Request:**
```json
{"old_password": "actual", "new_password": "nueva"}
```

**Response:** `204 No Content`

**Errores:**
- `401` — contraseña actual incorrecta
- `404` — usuario no encontrado

## Usuarios — ADMIN

```
POST   /users
GET    /users
GET    /users/{id}
PUT    /users/{id}
DELETE /users/{id}
```

**Requiere:** `Authorization: Bearer <token>` + rol `ADMIN` + contraseña no temporal.

### POST /users

Crea un usuario `EXECUTOR` o `AUDITOR`. No permite crear `ADMIN`.

**Request:**
```json
{"name": "Nombre", "email": "user@test.com", "password": "temp123", "role": "EXECUTOR"}
```

**Response 201:**
```json
{
  "ID": "uuid",
  "Name": "Nombre",
  "Email": "user@test.com",
  "Role": "EXECUTOR",
  "MustChangePassword": true,
  "Active": true,
  "CreatedAt": "2026-01-01T00:00:00Z",
  "UpdatedAt": "2026-01-01T00:00:00Z"
}
```

**Errores:**
- `403` — no se puede crear usuarios ADMIN
- `409` — el email ya existe

### GET /users

Lista todos los usuarios.

**Response 200:** array de `UserOutput`

### GET /users/{id}

Obtiene detalle de un usuario.

**Response 200:** `UserOutput`
**Errores:** `404` — usuario no encontrado

### PUT /users/{id}

Actualiza datos de un usuario. Campos no enviados no se modifican.

**Request (parcial):**
```json
{"name": "Nuevo Nombre", "active": false}
```

**Response 200:** `UserOutput` actualizado
**Errores:** `404`, `409` — email duplicado

### DELETE /users/{id}

Desactiva logicamente un usuario (`active = false`).

**Response:** `204 No Content`
**Errores:** `404` — usuario no encontrado

## Tareas — ADMIN y AUDITOR

```
POST   /tasks          (ADMIN)
GET    /tasks          (ADMIN, AUDITOR)
GET    /tasks/{id}     (ADMIN, AUDITOR)
PUT    /tasks/{id}     (ADMIN, solo ASSIGNED)
DELETE /tasks/{id}     (ADMIN, solo ASSIGNED)
```

**Requiere:** `Authorization: Bearer <token>` + rol correspondiente + contraseña no temporal.

### POST /tasks

Crea una tarea asignada a un `EXECUTOR`. La tarea nace en `ASSIGNED`.

**Request:**
```json
{"title": "Tarea", "description": "Desc", "due_at": "2026-12-31T00:00:00Z", "assignee_id": "uuid-executor"}
```

**Response 201:** `TaskOutput` con `Status: "ASSIGNED"`

**Errores:**
- `403`/`500` — asignatario no es EXECUTOR o no existe

### GET /tasks

Lista todas las tareas.

**Response 200:** array de `TaskOutput`

### GET /tasks/{id}

Obtiene detalle de una tarea.

**Response 200:** `TaskOutput`
**Errores:** `404` — tarea no encontrada

### PUT /tasks/{id}

Actualiza una tarea. Solo permitido si la tarea esta en estado `ASSIGNED`.

**Request (parcial):**
```json
{"title": "Nuevo titulo", "assignee_id": "uuid-executor"}
```

**Response 200:** `TaskOutput` actualizado
**Errores:** `404`, `409` — la tarea no puede ser modificada en su estado actual

### DELETE /tasks/{id}

Elimina una tarea. Solo permitido si esta en estado `ASSIGNED`.

**Response:** `204 No Content`
**Errores:** `404`, `409` — no modificable

## Tareas — EXECUTOR

```
GET   /me/tasks
GET   /me/tasks/{id}
PATCH /me/tasks/{id}/status
POST  /me/tasks/{id}/comments
```

**Requiere:** `Authorization: Bearer <token>` + rol `EXECUTOR` + contraseña no temporal.

### GET /me/tasks

Lista las tareas propias del ejecutor autenticado.

**Response 200:** array de `TaskOutput` (solo tareas con `assignee_id` del usuario)

### GET /me/tasks/{id}

Obtiene detalle de una tarea propia.

**Response 200:** `TaskOutput`
**Errores:** `403` — la tarea no pertenece al usuario, `404` — no encontrada

### PATCH /me/tasks/{id}/status

Cambia el estado de una tarea propia segun el flujo permitido.

**Request:**
```json
{"new_status": "STARTED"}
```

**Response 200:** `TaskOutput` con nuevo estado

**Errores:**
- `403` — la tarea no pertenece al usuario
- `409` — tarea vencida, transición no permitida o estado terminal

**Transiciones permitidas:**
- `ASSIGNED -> STARTED`
- `STARTED -> WAITING | FINISHED_SUCCESS | FINISHED_ERROR`
- `WAITING -> WAITING | FINISHED_SUCCESS | FINISHED_ERROR`

### POST /me/tasks/{id}/comments

Comenta una tarea vencida propia. Solo permitido si la tarea esta vencida (`due_at < now`).

**Request:**
```json
{"comment": "Texto del comentario"}
```

**Response 201:** `CommentOutput`
**Errores:** `403` — no es propietario, `409` — tarea no vencida

## Formato de respuestas

### TaskOutput
```json
{
  "ID": "uuid",
  "Title": "...",
  "Description": "...",
  "DueAt": "2026-12-31T00:00:00Z",
  "Status": "ASSIGNED",
  "AssigneeID": "uuid",
  "CreatedBy": "uuid",
  "CreatedAt": "2026-01-01T00:00:00Z",
  "UpdatedAt": "2026-01-01T00:00:00Z"
}
```

### UserOutput
```json
{
  "ID": "uuid",
  "Name": "...",
  "Email": "...",
  "Role": "EXECUTOR",
  "MustChangePassword": true,
  "Active": true,
  "CreatedAt": "2026-01-01T00:00:00Z",
  "UpdatedAt": "2026-01-01T00:00:00Z"
}
```

### CommentOutput
```json
{
  "ID": "uuid",
  "TaskID": "uuid",
  "UserID": "uuid",
  "Comment": "...",
  "CreatedAt": "2026-01-01T00:00:00Z"
}
```

### Error
```json
{"error": "mensaje descriptivo"}
```

## Health

```
GET /health
```

**Response 200:**
```json
{"status":"ok"}
```
