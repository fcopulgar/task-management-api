# API backend

## Estado

`PENDIENTE DE IMPLEMENTACION`

Los endpoints estan planificados; no existen endpoints reales todavia.

## Auth

```text
POST /auth/login
POST /auth/logout
PUT  /auth/password
```

Reglas:

- Login entrega token con `session_id`.
- Logout revoca la sesion actual.
- Cambio de contrasena aplica a cualquier perfil autenticado.

## Usuarios — ADMIN

```text
POST   /users
GET    /users
GET    /users/{id}
PUT    /users/{id}
DELETE /users/{id}
```

Reglas:

- Solo `ADMIN`.
- `POST /users` solo permite crear `EXECUTOR` o `AUDITOR`.
- `DELETE /users/{id}` debe interpretarse preferentemente como desactivacion logica.

## Tareas — ADMIN y AUDITOR

```text
POST   /tasks
GET    /tasks
GET    /tasks/{id}
PUT    /tasks/{id}
DELETE /tasks/{id}
```

Reglas:

- `POST /tasks`: solo `ADMIN`.
- `PUT /tasks/{id}`: solo `ADMIN` y solo si estado `ASSIGNED`.
- `DELETE /tasks/{id}`: solo `ADMIN` y solo si estado `ASSIGNED`.
- `GET /tasks`: `ADMIN` y `AUDITOR` pueden listar todas.
- `GET /tasks/{id}`: `ADMIN` y `AUDITOR` pueden ver detalle.

## Tareas — EXECUTOR

```text
GET   /me/tasks
GET   /me/tasks/{id}
PATCH /me/tasks/{id}/status
POST  /me/tasks/{id}/comments
```

Reglas:

- Solo `EXECUTOR`.
- Solo sobre tareas propias.
- No puede cambiar estado si la tarea esta vencida.
- Solo puede comentar tarea vencida propia.
