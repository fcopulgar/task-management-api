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

## Planificado

- Persistencia en PostgreSQL con GORM y `AutoMigrate`.
- Autenticacion JWT con `session_id`.
- Sesiones revocables persistidas en base de datos.
- Hash de contrasenas con bcrypt.
- Casos de uso, handlers, repositorios (Fases 3-9).

## No implementado todavia

- Modelos GORM.
- Puertos de aplicacion y contratos.
- Casos de uso, handlers, repositorios y servicios de negocio.
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
