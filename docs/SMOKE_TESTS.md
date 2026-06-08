# Smoke tests

## Estado

`IMPLEMENTADO Y EJECUTADO`

## Prerrequisitos

- `docker compose up -d` con PostgreSQL y la aplicacion corriendo.
- Un usuario `ADMIN` insertado manualmente (via `psql`) con `must_change_password = false`.
- `JWT_SECRET` configurado.

### Seed manual de ADMIN

```sql
INSERT INTO users (id, name, email, password_hash, role, must_change_password, active, created_at, updated_at)
VALUES (gen_random_uuid(), 'Admin', 'admin@test.com', '<bcrypt-hash>', 'ADMIN', false, true, NOW(), NOW());
```

## Smoke test 01 — Health check

```bash
curl -s http://localhost:8080/health
# {"status":"ok"}
```

## Smoke test 02 — Login exitoso

```bash
curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"admin123"}'
# {"token":"eyJ..."}
```

## Smoke test 03 — Login credenciales invalidas

```bash
curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"wrong"}'
# {"error":"credenciales invalidas"}  (401)
```

## Smoke test 04 — Usuario inactivo rechazado

1. Desactivar un usuario via `PUT /users/{id}` con `{"active":false}`.
2. Intentar login con ese usuario.
3. Debe retornar `401` "credenciales invalidas".

## Smoke test 05 — Logout revoca sesion

```bash
TOKEN=$(login)
curl -s -X POST http://localhost:8080/auth/logout \
  -H "Authorization: Bearer $TOKEN"
# {"message":"sesion cerrada"}  (200)

curl -s http://localhost:8080/users \
  -H "Authorization: Bearer $TOKEN"
# {"error":"sesion revocada"}  (401)
```

## Smoke test 06 — Token con sesion revocada es rechazado

Ver smoke test 05. La segunda llamada con el mismo token despues del logout debe fallar.

## Smoke test 07 — Contrasena temporal bloquea acceso

1. Crear un usuario con `POST /users` (nace con `must_change_password = true`).
2. Login con ese usuario.
3. Intentar acceder a `GET /tasks` o `GET /users`.
4. Debe retornar `403` "debe cambiar su contrasena antes de continuar".
5. `POST /auth/password` debe funcionar (cambio de contrasena).
6. `POST /auth/logout` debe funcionar (cierre de sesion).
7. Despues del cambio, las rutas protegidas deben ser accesibles.

## Smoke test 08 — ADMIN no puede crear otro ADMIN

```bash
TOKEN=$(admin_login)
curl -s -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"Admin2","email":"a2@test.com","password":"pwd","role":"ADMIN"}'
# {"error":"no se puede crear usuarios ADMIN"}  (403)
```

## Smoke test 09 — ADMIN crea EXECUTOR

```bash
TOKEN=$(admin_login)
curl -s -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"Exec","email":"exec@test.com","password":"pwd","role":"EXECUTOR"}'
# 201, MustChangePassword: true, Active: true
```

## Smoke test 10 — Desactivacion logica

```bash
TOKEN=$(admin_login)
USER_ID=$(get_user_id "exec@test.com")
curl -s -X DELETE http://localhost:8080/users/$USER_ID \
  -H "Authorization: Bearer $TOKEN"
# 204 No Content

curl -s http://localhost:8080/users/$USER_ID \
  -H "Authorization: Bearer $TOKEN"
# Active: false
```

## Smoke test 11 — Creacion de tarea solo a EXECUTOR

```bash
TOKEN=$(admin_login)
EXEC_ID=$(get_user_id "exec@test.com")
curl -s -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{\"title\":\"T1\",\"due_at\":\"2026-12-31T00:00:00Z\",\"assignee_id\":\"$EXEC_ID\"}"
# 201, Status: ASSIGNED
```

## Smoke test 12 — Tarea solo actualizable en ASSIGNED

1. Crear tarea (ASSIGNED).
2. `PUT /tasks/{id}` con `{"title":"Updated"}` -> 200 OK.
3. Cambiar estado a STARTED (via executor `PATCH /me/tasks/{id}/status`).
4. `PUT /tasks/{id}` con `{"title":"Updated2"}` -> 409 "la tarea no puede ser modificada en su estado actual".

## Smoke test 13 — AUDITOR solo lectura

```bash
# Crear auditor via ADMIN, cambiar contrasena
AUD_TOKEN=$(auditor_login)
curl -s http://localhost:8080/tasks \
  -H "Authorization: Bearer $AUD_TOKEN"
# 200, lista de tareas

curl -s -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $AUD_TOKEN" \
  -d '{"title":"T","due_at":"2026-12-31T00:00:00Z","assignee_id":"uuid"}'
# 403
```

## Smoke test 14 — EXECUTOR solo ve tareas propias

```bash
EXEC_TOKEN=$(executor_login)
curl -s http://localhost:8087/me/tasks \
  -H "Authorization: Bearer $EXEC_TOKEN"
# Solo tareas con assignee_id del executor autenticado
```

## Smoke test 15 — EXECUTOR cambia estado (flujo valido)

```bash
EXEC_TOKEN=$(executor_login)
TASK_ID=$(get_my_task_id)
curl -s -X PATCH http://localhost:8080/me/tasks/$TASK_ID/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $EXEC_TOKEN" \
  -d '{"new_status":"STARTED"}'
# 200, Status: STARTED
```

## Smoke test 16 — EXECUTOR bloqueado en tarea vencida

1. ADMIN crea tarea con `due_at` en el pasado asignada al executor.
2. Executor intenta `PATCH /me/tasks/{id}/status`.
3. Debe retornar `409` "la tarea esta vencida".

## Smoke test 17 — EXECUTOR comenta tarea vencida

1. Tarea vencida asignada al executor.
2. `POST /me/tasks/{id}/comments` con `{"comment":"..."}`.
3. Debe retornar `201` con el comentario creado.

## Smoke test 18 — EXECUTOR no comenta tarea no vencida

1. Tarea no vencida asignada al executor.
2. `POST /me/tasks/{id}/comments`.
3. Debe retornar `409` "solo se puede comentar tareas vencidas".

## Smoke test 19 — Transiciones invalidas rechazadas

1. Tarea en ASSIGNED.
2. `PATCH /me/tasks/{id}/status` con `{"new_status":"FINISHED_SUCCESS"}`.
3. Debe retornar `409` "transicion de estado no permitida".

## Resultado de smoke tests

Todos los 19 smoke tests fueron ejecutados exitosamente durante las validaciones E2E de las fases 5-8.
