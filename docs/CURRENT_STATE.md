# Estado actual

## Resumen

El proyecto `task-management-api` cuenta con harness documental, especificacion funcional inicial, stack tecnico definido, arquitectura objetivo, ADRs iniciales y plan de implementacion por fases. Todavia no existe implementacion funcional.

## Implementado

- Harness documental base.
- Reglas de trabajo para agentes.
- Estructura inicial de documentacion.
- Especificacion del producto.
- Requisitos funcionales y no funcionales iniciales.
- Arquitectura objetivo hexagonal minimalista.
- Documentacion backend planificada.
- ADRs iniciales `0001` a `0008`.
- `PLAN.md` con fases de implementacion.
- **Fase 1 — Bootstrap tecnico del proyecto**.
- Modulo Go inicializado (`go.mod`, `go.sum`).
- Servidor HTTP minimo con `chi` (solo endpoint `/health`).
- Configuracion de entorno (`internal/config/config.go`, `.env.example`).
- `Makefile`, `Dockerfile`, `docker-compose.yml`, `.gitignore`.
- **Fase 2 — Dominio y reglas centrales**.
- Entidades de dominio: `User`, `Session`, `Task`, `Comment`.
- Tipos de dominio: `Role` (`ADMIN`, `EXECUTOR`, `AUDITOR`), `TaskStatus` (`ASSIGNED`, `STARTED`, `WAITING`, `FINISHED_SUCCESS`, `FINISHED_ERROR`).
- Validacion de transiciones de estado.
- Reglas de vencimiento y propiedad de tareas.
- Errores de dominio.
- Tests unitarios de reglas criticas (36 tests, `testing` + `testify`).
- **Fase 3 — Puertos de aplicacion y contratos internos**.
- Interfaces de repositorios: `UserRepository`, `SessionRepository`, `TaskRepository`, `CommentRepository`.
- Interfaces de seguridad: `PasswordHasher`, `TokenService`.
- `TokenClaims` con `UserID`, `Role` y `SessionID`.
- DTOs de entrada/salida: `LoginInput/Output`, `ChangePasswordInput`, `CreateUserInput`, `UpdateUserInput`, `UserOutput`, `CreateTaskInput`, `UpdateTaskInput`, `TransitionTaskInput`, `TaskOutput`, `CreateCommentInput`, `CommentOutput`.
- Funciones de mapeo dominio->DTO: `UserToOutput`, `TaskToOutput`, `CommentToOutput`.
- Errores de aplicacion: `ErrUserNotFound`, `ErrTaskNotFound`, `ErrSessionNotFound`, `ErrInvalidCredentials`, `ErrUnauthorized`, `ErrEmailAlreadyExists`, `ErrCannotCreateAdmin`, `ErrTaskNotModifiable`, `ErrTaskOverdue`.
- Mocks de repositorios y servicios de seguridad para tests.
- Tests de contratos (interfaces + DTOs, 10 tests adicionales).
- **Fase 4 — Persistencia GORM y AutoMigrate**.
- Modelos GORM separados del dominio: `UserModel`, `SessionModel`, `TaskModel`, `CommentModel` (UUIDs via `gen_random_uuid()`).
- Mappers dominio↔modelos GORM: `UserToModel/FromModel`, `SessionToModel/FromModel`, `TaskToModel/FromModel`, `CommentToModel/FromModel`.
- Repositorios outbound: `GormUserRepository`, `GormSessionRepository`, `GormTaskRepository`, `GormCommentRepository`.
- Conexion PostgreSQL con `gorm.io/driver/postgres`.
- `AutoMigrate` ejecutado al iniciar el servidor.
- Integracion en `cmd/server/main.go`.
- Tests de persistencia (13 tests: 9 mappers + 4 repositorios con PostgreSQL real).
- `go.mod` actualizado a Go 1.25, `Dockerfile` a `golang:1.25-alpine`.
- **Fase 5 — Autenticacion, sesiones y cambio de contrasena**.
- Adaptadores de seguridad: `BcryptHasher` (`golang.org/x/crypto/bcrypt`), `JWTTokenService` (`github.com/golang-jwt/jwt/v5`).
- Caso de uso `AuthUseCase` con `Login`, `Logout`, `ChangePassword`.
- Endpoints HTTP: `POST /auth/login`, `POST /auth/logout`, `POST /auth/password`.
- Middleware `Authenticate`: extrae JWT, valida firma, verifica sesion no revocada, verifica usuario activo.
- Middleware `RequirePasswordNotTemporary`: bloquea acceso si `must_change_password = true` (excepto `/auth/password` y `/auth/logout`).
- Contexto de autenticacion via `SetAuthInfo`/`GetAuthInfo` con `UserID`, `Role`, `SessionID`.
- `Dependencies` container (`application.Dependencies`) para wire-up de puertos.
- Tests: 8 security + 9 use case + 6 handler + 10 middleware = 33 tests adicionales.
- **Fase 6 — Gestion de usuarios administrador**.
- Caso de uso `UserUseCase` con `CreateUser`, `ListUsers`, `GetUser`, `UpdateUser`, `DeactivateUser`.
- Handlers HTTP: `POST /users`, `GET /users`, `GET /users/{id}`, `PUT /users/{id}`, `DELETE /users/{id}`.
- Middleware `RequireRole`: autorizacion por perfil (solo `ADMIN` accede a `/users`).
- Restriccion: `ADMIN` no puede crear otros `ADMIN` (RF-006).
- Usuarios nuevos creados con `must_change_password = true`.
- Desactivacion logica (`active = false`) via `DELETE /users/{id}`.
- `password_hash` nunca expuesto en respuestas API.
- Validacion de email unico al crear/actualizar.
- Tests: 10 use case + 8 handler = 18 tests adicionales.
- **Fase 7 — Gestion de tareas administrador y auditor**.
- Caso de uso `TaskUseCase` con `CreateTask`, `ListTasks`, `GetTask`, `UpdateTask`, `DeleteTask`.
- Validacion de asignacion solo a `EXECUTOR` (busca usuario en repositorio, valida rol).
- Restriccion de actualizacion/eliminacion solo en estado `ASSIGNED` (`ErrTaskNotModifiable`).
- Middleware `RequireAnyRole(roles ...domain.Role)` para rutas compartidas.
- Rutas `/tasks` GET (ADMIN + AUDITOR), POST/PUT/DELETE (ADMIN solo).
- Handlers HTTP con `handleTaskError` para errores de tarea.
- `Dependencies` ampliado con `TaskRepo`.
- Tests: 8 use case + 6 handler = 14 tests adicionales.

## Planificado

- Flujos de ejecutor (Fase 8).
- Tests HTTP, hardening, documentacion final (Fase 9).

## No implementado todavia

- Casos de uso de autenticacion, usuarios y tareas.
- Handlers HTTP de negocio.
- Servicios de seguridad (bcrypt, JWT).
- Middleware de autenticacion y autorizacion.
- Endpoints de negocio.

## Endpoints planificados

Los endpoints estan documentados en `docs/backend/api.md` y se encuentran `PENDIENTE DE IMPLEMENTACION`.

## Decisiones confirmadas

- La documentacion se mantiene en espanol.
- Se usan ADRs versionados.
- El desarrollo se organiza mediante `PLAN.md`.
- El lenguaje principal sera Go.
- La arquitectura objetivo sera hexagonal minimalista.
- El router HTTP sera `chi`.
- La base de datos sera PostgreSQL.
- El ORM sera GORM.
- La inicializacion de esquema usara GORM `AutoMigrate`.
- No se usaran migraciones SQL versionadas en esta etapa.
- El dominio y la capa de aplicacion no dependeran de GORM.
- GORM vivira solo en adapters outbound de persistencia.
- La autenticacion usara JWT con `session_id`.
- Logout revocara la sesion en base de datos.
- Las contrasenas se hashearan con bcrypt.

## Pendiente de definicion

- Estrategia de despliegue.
- Politicas finales de secretos.
- Filtros avanzados para auditoria.
- Detalles listados en `docs/requirements/open-questions.md`.

## Riesgos tecnicos

- La estrategia de migracion futura debe revisarse si el esquema requiere evolucion controlada mas alla de `AutoMigrate`.
- La revocacion de sesiones requiere validacion persistente en endpoints protegidos.
- Las reglas de contrasena temporal deben aplicarse antes de permitir acceso a funcionalidades protegidas.

## Validacion disponible

- Validacion documental de estructura y archivos Markdown.
- Build: `go build ./cmd/server` (requiere Go 1.24+ o `docker compose build`).
- Ejecucion local: `docker compose up` o `go run ./cmd/server` + PostgreSQL.
- Health check: `GET /health` retorna `{"status":"ok"}`.

## Que no se pudo validar

- Tests de negocio.
- Integraciones de dominio y persistencia.
- Endpoints de negocio.

Estos puntos no se pudieron validar porque las fases correspondientes estan `PENDIENTE DE IMPLEMENTACION`.
