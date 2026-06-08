# Estado actual

## Resumen

El proyecto `task-management-api` cuenta con implementacion funcional completa de las fases 1-8. API REST backend operativa con autenticacion JWT, sesiones revocables, gestion de usuarios por ADMIN, gestion de tareas por ADMIN y AUDITOR, y flujos de ejecutor con transiciones de estado y comentarios.

## Implementado

- Harness documental base.
- Reglas de trabajo para agentes.
- Estructura inicial de documentacion.
- Especificacion del producto.
- Requisitos funcionales y no funcionales iniciales.
- Arquitectura objetivo hexagonal minimalista.
- Documentacion backend actualizada.
- ADRs iniciales `0001` a `0008`.
- `PLAN.md` con fases de implementacion.
- **Fase 1 — Bootstrap tecnico del proyecto**.
- Modulo Go inicializado (`go.mod`, `go.sum`).
- Servidor HTTP minimo con `chi` (endpoint `/health`).
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
- DTOs de entrada/salida para todos los casos de uso.
- Funciones de mapeo dominio->DTO.
- Errores de aplicacion.
- Mocks de repositorios y servicios de seguridad para tests.
- Tests de contratos (interfaces + DTOs, 10 tests).
- **Fase 4 — Persistencia GORM y AutoMigrate**.
- Modelos GORM separados del dominio con UUIDs via `gen_random_uuid()`.
- Mappers dominio↔modelos GORM.
- Repositorios outbound: `GormUserRepository`, `GormSessionRepository`, `GormTaskRepository`, `GormCommentRepository`.
- Conexion PostgreSQL con `gorm.io/driver/postgres`.
- `AutoMigrate` ejecutado al iniciar el servidor.
- Tests de persistencia (13 tests: 9 mappers + 4 repositorios con PostgreSQL real).
- **Fase 5 — Autenticacion, sesiones y cambio de contrasena**.
- Adaptadores: `BcryptHasher`, `JWTTokenService` (HS256, `session_id` en claims).
- Caso de uso `AuthUseCase` con `Login`, `Logout`, `ChangePassword`.
- Endpoints: `POST /auth/login`, `POST /auth/logout`, `POST /auth/password`.
- Middleware `Authenticate`: valida JWT, verifica sesion no revocada, verifica usuario activo.
- Middleware `RequirePasswordNotTemporary`: bloquea acceso si `must_change_password = true`.
- Contexto de autenticacion via `SetAuthInfo`/`GetAuthInfo`.
- Tests: 33 (8 security + 9 use case + 6 handler + 10 middleware).
- **Fase 6 — Gestion de usuarios administrador**.
- Caso de uso `UserUseCase`: `CreateUser`, `ListUsers`, `GetUser`, `UpdateUser`, `DeactivateUser`.
- Endpoints: `POST/GET /users`, `GET/PUT/DELETE /users/{id}`.
- Middleware `RequireRole`: autorizacion por perfil (solo `ADMIN` accede a `/users`).
- Restriccion: `ADMIN` no puede crear otros `ADMIN`.
- `password_hash` nunca expuesto en respuestas API.
- Tests: 18 (10 use case + 8 handler).
- **Fase 7 — Gestion de tareas administrador y auditor**.
- Caso de uso `TaskUseCase`: `CreateTask`, `ListTasks`, `GetTask`, `UpdateTask`, `DeleteTask`.
- Validacion de asignacion solo a `EXECUTOR`.
- Restriccion de actualizacion/eliminacion solo en `ASSIGNED`.
- Middleware `RequireAnyRole(roles ...)` para rutas compartidas (ADMIN + AUDITOR).
- Endpoints `/tasks` GET (ADMIN+AUDITOR), POST/PUT/DELETE (ADMIN).
- Tests: 14 (8 use case + 6 handler).
- **Fase 8 — Flujos de ejecutor**.
- Caso de uso `ExecutorUseCase`: `ListMyTasks`, `GetMyTask`, `TransitionTask`, `CommentOnTask`.
- Validacion de propiedad de tarea.
- Bloqueo de cambio de estado en tareas vencidas.
- Comentarios solo en tareas vencidas propias.
- Endpoints `/me/tasks` GET (listar/detalle), PATCH (transicion), POST (comentarios).
- Tests: 16 (10 use case + 6 handler).
- **Fase 9 — Tests HTTP, hardening y documentacion final**.
- API documentada en `docs/backend/api.md` con todos los endpoints, request/response y errores.
- Smoke tests documentados en `docs/SMOKE_TESTS.md` (19 escenarios ejecutados).
- Checklist de produccion actualizada en `docs/PRODUCTION_CHECKLIST.md`.
- `docs/CURRENT_STATE.md` actualizado para reflejar estado final.

## Planificado

Sin fases planificadas pendientes. El desarrollo planificado en `PLAN.md` esta completo (Fases 0-9).

## No implementado todavia

Ningun item pendiente dentro del alcance definido.

## Endpoints implementados

Ver `docs/backend/api.md` para documentacion detallada de todos los endpoints.

- `GET /health`
- `POST /auth/login`
- `POST /auth/logout`
- `POST /auth/password`
- `POST /users`, `GET /users`, `GET /users/{id}`, `PUT /users/{id}`, `DELETE /users/{id}`
- `POST /tasks`, `GET /tasks`, `GET /tasks/{id}`, `PUT /tasks/{id}`, `DELETE /tasks/{id}`
- `GET /me/tasks`, `GET /me/tasks/{id}`, `PATCH /me/tasks/{id}/status`, `POST /me/tasks/{id}/comments`

## Decisiones confirmadas

- La documentacion se mantiene en espanol.
- Se usan ADRs versionados.
- El desarrollo se organiza mediante `PLAN.md`.
- El lenguaje principal es Go.
- La arquitectura objetivo es hexagonal minimalista.
- El router HTTP es `chi`.
- La base de datos es PostgreSQL.
- El ORM es GORM.
- La inicializacion de esquema usa GORM `AutoMigrate`.
- No se usan migraciones SQL versionadas en esta etapa.
- El dominio y la capa de aplicacion no dependen de GORM.
- GORM vive solo en adapters outbound de persistencia.
- La autenticacion usa JWT con `session_id`.
- Logout revoca la sesion en base de datos.
- Las contrasenas se hashean con bcrypt.

## Pendiente de definicion

- Estrategia de despliegue.
- Politicas finales de secretos.
- Filtros avanzados para auditoria.
- Detalles listados en `docs/requirements/open-questions.md`.

## Riesgos tecnicos

- La estrategia de migracion futura debe revisarse si el esquema requiere evolucion controlada mas alla de `AutoMigrate`.
- La revocacion de sesiones requiere validacion persistente en endpoints protegidos. (Implementado)
- Las reglas de contrasena temporal deben aplicarse antes de permitir acceso a funcionalidades protegidas. (Implementado)

## Validacion disponible

- Build: `go build ./cmd/server` (requiere Go 1.25+ o `docker compose build`).
- Ejecucion local: `docker compose up` o `go run ./cmd/server` + PostgreSQL.
- Health check: `GET /health` retorna `{"status":"ok"}`.
- Tests unitarios: `go test ./internal/domain/... ./internal/application/... ./internal/adapters/...` (~140 tests).
- Tests de integracion: repositorios GORM requieren PostgreSQL corriendo.
- Smoke tests: 19 escenarios documentados en `docs/SMOKE_TESTS.md`.
- Linter: `go vet ./...`.

## Que no se pudo validar

- Despliegue en produccion (no definido).
- Estrategia de respaldo de base de datos (no definida).
- Monitoreo operacional (no definido).
