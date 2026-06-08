# task-management-api

## Descripcion

`task-management-api` es una API REST para gestion de usuarios y tareas con autenticacion, autorizacion por perfiles, cambio obligatorio de contrasena temporal, cierre de sesion y control de estados de tareas.

El sistema contempla tres perfiles: `ADMIN`, `EXECUTOR` y `AUDITOR`. El objetivo es mantener una base backend clara, segura, testeable y extensible, evitando sobredimensionar la solucion.

## Documentacion

- `PLAN.md`: plan de implementacion por fases.
- `AGENTS.md`: reglas obligatorias para agentes.
- `docs/CURRENT_STATE.md`: estado real y vigente del proyecto.
- `docs/backend/api.md`: documentacion completa de la API.
- `docs/requirements/`: alcance y requisitos funcionales/no funcionales.
- `docs/architecture/`: arquitectura objetivo.
- `docs/decisions/`: ADRs versionados (0001-0008).
- `docs/DOMAIN_GUARDRAILS.md`: reglas de negocio que no deben romperse.
- `docs/SMOKE_TESTS.md`: smoke tests documentados.
- `docs/PRODUCTION_CHECKLIST.md`: checklist para produccion.
- `docs/TECH_DEBT.md`: deuda tecnica registrada.
- `docs/standards/`: estandares de documentacion y validacion.

## Stack

- **Lenguaje:** Go 1.25+
- **Router HTTP:** `chi`
- **ORM:** GORM
- **Base de datos:** PostgreSQL 17
- **Inicializacion de esquema:** GORM `AutoMigrate`
- **Autenticacion:** JWT (HS256) con `session_id` en claims
- **Sesiones:** revocables, persistidas en base de datos
- **Hash de contrasenas:** bcrypt
- **Testing:** `testing`, `testify`, `httptest`

## Requisitos

- **Go 1.25+** (si se ejecuta sin Docker)
- **Docker + Docker Compose** (recomendado)
- **PostgreSQL 17** (si se ejecuta sin Docker)

## Instalacion y ejecucion

### Con Docker Compose (recomendado)

```bash
# Clonar el repositorio
git clone <repo-url> && cd task-management-api

# Crear archivo .env desde el ejemplo
cp .env.example .env

# Editar .env con valores reales (especialmente JWT_SECRET)
# vim .env

# Construir y levantar los servicios
docker compose up -d --build

# Verificar que esta corriendo
curl http://localhost:8080/health
# {"status":"ok"}
```

### Sin Docker (desarrollo local)

```bash
# Requisitos previos: Go 1.25+ y PostgreSQL corriendo en localhost:5432

cp .env.example .env
# Editar .env con JWT_SECRET y credenciales de PostgreSQL

# Instalar dependencias
go mod download

# Ejecutar migraciones y servidor
go run ./cmd/server

# O compilar y ejecutar
make build
./bin/server
```

## Tests

```bash
# Todos los tests (requiere PostgreSQL corriendo para tests de persistencia)
make test

# Solo tests unitarios (sin dependencia de base de datos)
go test ./internal/domain/... ./internal/application/... ./internal/adapters/inbound/http/... ./internal/adapters/outbound/security/...

# Con Docker (sin Go instalado)
docker run --rm -v "$(pwd):/app" -w /app -e GOTOOLCHAIN=auto golang:1.25-alpine go test ./internal/domain/... ./internal/application/... ./internal/adapters/inbound/http/... ./internal/adapters/outbound/security/...

# Tests de persistencia (requiere PostgreSQL)
# Con docker compose corriendo:
docker run --rm -v "$(pwd):/app" -w /app --network task-management-api_default -e GOTOOLCHAIN=auto -e DB_HOST=db -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=postgres -e DB_NAME=taskmanagement golang:1.25-alpine go test ./internal/adapters/outbound/persistence/...
```

## Endpoints

Documentacion completa en `docs/backend/api.md`.

| Metodo | Ruta | Perfil | Descripcion |
|--------|------|--------|-------------|
| `GET` | `/health` | Publico | Health check |
| `POST` | `/auth/login` | Publico | Iniciar sesion |
| `POST` | `/auth/logout` | Autenticado | Cerrar sesion |
| `POST` | `/auth/password` | Autenticado | Cambiar contrasena |
| `POST` | `/users` | ADMIN | Crear usuario |
| `GET` | `/users` | ADMIN | Listar usuarios |
| `GET` | `/users/{id}` | ADMIN | Ver usuario |
| `PUT` | `/users/{id}` | ADMIN | Actualizar usuario |
| `DELETE` | `/users/{id}` | ADMIN | Desactivar usuario |
| `POST` | `/tasks` | ADMIN | Crear tarea |
| `GET` | `/tasks` | ADMIN, AUDITOR | Listar tareas |
| `GET` | `/tasks/{id}` | ADMIN, AUDITOR | Ver tarea |
| `PUT` | `/tasks/{id}` | ADMIN | Actualizar tarea (solo ASSIGNED) |
| `DELETE` | `/tasks/{id}` | ADMIN | Eliminar tarea (solo ASSIGNED) |
| `GET` | `/me/tasks` | EXECUTOR | Listar mis tareas |
| `GET` | `/me/tasks/{id}` | EXECUTOR | Ver mi tarea |
| `PATCH` | `/me/tasks/{id}/status` | EXECUTOR | Cambiar estado |
| `POST` | `/me/tasks/{id}/comments` | EXECUTOR | Comentar tarea vencida |

## Ejemplos de uso

### Seed inicial: crear un ADMIN directamente en la base de datos

El sistema no incluye un endpoint publico para crear el primer `ADMIN`. Debe insertarse manualmente:

```bash
# Con docker compose corriendo, generar hash bcrypt:
docker compose exec app sh -c "echo 'package main
import (\"fmt\"; \"golang.org/x/crypto/bcrypt\")
func main() { h,_:=bcrypt.GenerateFromPassword([]byte(\"admin123\"),10); fmt.Print(string(h)) }' > /tmp/hash.go && cd /app && go run /tmp/hash.go"

# Insertar usuario ADMIN con el hash generado
docker compose exec db psql -U postgres -d taskmanagement -c "
INSERT INTO users (id, name, email, password_hash, role, must_change_password, active, created_at, updated_at)
VALUES (gen_random_uuid(), 'Admin', 'admin@test.com', '<HASH>', 'ADMIN', false, true, NOW(), NOW());
"
```

O alternativamente, usar una herramienta externa como `htpasswd -bnBC 10 "" admin123 | tr -d ':\n'` para generar el hash.

### Autenticacion

```bash
# Login como ADMIN
curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"admin123"}'
# {"token":"eyJhbGciOiJIUzI1NiIs..."}

# Guardar token para uso posterior
TOKEN="eyJhbGciOiJIUzI1NiIs..."
AUTH="Authorization: Bearer $TOKEN"
```

### Gestion de usuarios (ADMIN)

```bash
# Crear un EXECUTOR
curl -s -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -H "$AUTH" \
  -d '{"name":"Executor 1","email":"exec1@test.com","password":"temp123","role":"EXECUTOR"}'

# Crear un AUDITOR
curl -s -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -H "$AUTH" \
  -d '{"name":"Auditor 1","email":"aud1@test.com","password":"temp123","role":"AUDITOR"}'

# Listar todos los usuarios
curl -s http://localhost:8080/users -H "$AUTH"

# Desactivar un usuario (DELETE logico)
curl -s -X DELETE http://localhost:8080/users/{user-id} -H "$AUTH"

# Intentar crear otro ADMIN (rechazado)
curl -s -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -H "$AUTH" \
  -d '{"name":"Admin 2","email":"admin2@test.com","password":"pwd","role":"ADMIN"}'
# {"error":"no se puede crear usuarios ADMIN"}  (403)
```

### Gestion de tareas (ADMIN)

```bash
# Crear una tarea asignada a un EXECUTOR
curl -s -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "$AUTH" \
  -d '{"title":"Tarea 1","description":"Descripcion","due_at":"2026-12-31T00:00:00Z","assignee_id":"{executor-id}"}'

# Listar todas las tareas
curl -s http://localhost:8080/tasks -H "$AUTH"

# Actualizar una tarea (solo si esta en ASSIGNED)
curl -s -X PUT http://localhost:8080/tasks/{task-id} \
  -H "Content-Type: application/json" \
  -H "$AUTH" \
  -d '{"title":"Titulo actualizado"}'

# Eliminar una tarea
curl -s -X DELETE http://localhost:8080/tasks/{task-id} -H "$AUTH"
```

### Lectura de tareas (AUDITOR)

```bash
# Login como auditor (debe cambiar contrasena temporal primero)
AUD_TOKEN=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"aud1@test.com","password":"temp123"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['token'])")

# Cambiar contrasena obligatoria
curl -s -X POST http://localhost:8080/auth/password \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $AUD_TOKEN" \
  -d '{"old_password":"temp123","new_password":"auditor123"}'

# Re-login con nueva contrasena
AUD_TOKEN=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"aud1@test.com","password":"auditor123"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['token'])")

# Listar tareas (solo lectura)
curl -s http://localhost:8080/tasks -H "Authorization: Bearer $AUD_TOKEN"

# Intentar crear tarea (rechazado 403)
curl -s -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $AUD_TOKEN" \
  -d '{"title":"T","due_at":"2026-12-31T00:00:00Z","assignee_id":"id"}'
```

### Flujo de ejecutor (EXECUTOR)

```bash
# Login como executor
EXEC_TOKEN=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"exec1@test.com","password":"nueva"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['token'])")

# Listar mis tareas
curl -s http://localhost:8080/me/tasks -H "Authorization: Bearer $EXEC_TOKEN"

# Cambiar estado de una tarea (ASSIGNED -> STARTED)
curl -s -X PATCH http://localhost:8080/me/tasks/{task-id}/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $EXEC_TOKEN" \
  -d '{"new_status":"STARTED"}'

# Intentar cambiar estado de tarea vencida (rechazado 409)
curl -s -X PATCH http://localhost:8080/me/tasks/{overdue-task-id}/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $EXEC_TOKEN" \
  -d '{"new_status":"STARTED"}'
# {"error":"la tarea esta vencida"}

# Comentar tarea vencida propia
curl -s -X POST http://localhost:8080/me/tasks/{overdue-task-id}/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $EXEC_TOKEN" \
  -d '{"comment":"Trabajando en resolver el atraso"}'
```

### Logout y revocacion de sesion

```bash
# Cerrar sesion
curl -s -X POST http://localhost:8080/auth/logout -H "$AUTH"

# El token anterior ya no funciona
curl -s http://localhost:8080/users -H "$AUTH"
# {"error":"sesion revocada"}  (401)
```

## Estados de tarea y transiciones

```
ASSIGNED ──> STARTED ──> WAITING ──> FINISHED_SUCCESS
                         │
                         ├──> FINISHED_SUCCESS
                         ├──> FINISHED_ERROR
                         │
               STARTED ──┼──> FINISHED_SUCCESS
                         └──> FINISHED_ERROR

WAITING ──> WAITING (re-abrir)
```

- `ASSIGNED`: estado inicial al crear la tarea.
- `STARTED`: el ejecutor comienza a trabajar.
- `WAITING`: el ejecutor pone la tarea en espera.
- `FINISHED_SUCCESS` / `FINISHED_ERROR`: estados terminales.

## Makefile

```bash
make build        # Compilar binario
make run          # Ejecutar servidor
make test         # Ejecutar tests
make clean        # Limpiar binarios
make docker-build # Construir imagen Docker
make docker-up    # Levantar servicios con Docker Compose
make docker-down  # Detener servicios
make docker-logs  # Ver logs de los servicios
```

## Estructura del proyecto

```
.
├── cmd/server/           # Punto de entrada de la aplicacion
├── internal/
│   ├── config/           # Configuracion de entorno
│   ├── domain/           # Entidades y reglas de negocio
│   ├── application/      # Casos de uso, puertos y DTOs
│   │   ├── ports/        # Interfaces de repositorios y servicios
│   │   └── dto/          # Estructuras de entrada/salida
│   └── adapters/
│       ├── inbound/
│       │   └── http/     # Handlers HTTP y middleware
│       └── outbound/
│           ├── persistence/  # Modelos GORM y repositorios
│           └── security/     # Bcrypt y JWT
├── docs/                 # Documentacion completa
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── go.mod
└── .env.example
```
