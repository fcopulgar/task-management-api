# task-management-api

## Descripcion

`task-management-api` es una API REST planificada para gestion de usuarios y tareas con autenticacion, autorizacion por perfiles, cambio obligatorio de contrasena temporal, cierre de sesion y control de estados de tareas.

El sistema contempla tres perfiles: `ADMIN`, `EXECUTOR` y `AUDITOR`. El objetivo es mantener una base backend clara, segura, testeable y extensible, evitando sobredimensionar la solucion.

## Estado actual

- `IMPLEMENTADO`: harness documental base, reglas para agentes, fuentes de verdad, estructura de documentacion, ADRs iniciales, plan de implementacion por fases, bootstrap tecnico (Go, `chi`, servidor minimo, Docker Compose, Makefile).
- `PLANIFICADO`: implementacion de dominio, aplicacion, persistencia GORM, autenticacion JWT con `session_id`, sesiones revocables, bcrypt, endpoints de negocio y tests.
- `PENDIENTE DE IMPLEMENTACION`: codigo de dominio, casos de uso, endpoints, modelos GORM, repositorios, servicios de autenticacion, integraciones y tests de negocio.
- `PENDIENTE DE DEFINICION`: detalles operativos de despliegue, politicas finales de secretos, filtros avanzados de auditoria y decisiones funcionales listadas en preguntas abiertas.

## Alcance funcional planificado

- Autenticacion con login, logout y cambio de contrasena.
- Restriccion de acceso para usuarios inactivos o con contrasena temporal pendiente.
- Gestion de usuarios por `ADMIN`.
- Gestion de tareas por `ADMIN` con restricciones por estado.
- Consulta de tareas por `AUDITOR`.
- Trabajo de tareas propias por `EXECUTOR`, incluyendo cambio de estado y comentarios sobre tareas vencidas propias.
- Control de transiciones entre `ASSIGNED`, `STARTED`, `WAITING`, `FINISHED_SUCCESS` y `FINISHED_ERROR`.

## Stack definido

- Lenguaje: Go.
- Router HTTP: `chi`.
- ORM: GORM.
- Base de datos: PostgreSQL.
- Inicializacion de esquema: GORM `AutoMigrate`.
- Autenticacion: JWT con `session_id` incluido en claims.
- Logout: revocacion de sesion persistida en base de datos.
- Hash de contrasenas: bcrypt.
- Testing: `testing`, `testify` y `httptest` cuando corresponda.
- Desarrollo local futuro: Docker Compose y Makefile.
- Documentacion: Markdown, Mermaid y ADRs versionados.

## Estructura documental

- `PLAN.md`: plan activo de implementacion por fases.
- `AGENTS.md`: reglas obligatorias para agentes.
- `docs/CURRENT_STATE.md`: estado real y vigente del proyecto.
- `docs/requirements/`: alcance y requisitos.
- `docs/architecture/`: arquitectura objetivo.
- `docs/backend/`: contratos backend planificados.
- `docs/decisions/`: ADRs versionados.
- `docs/standards/`: estandares de documentacion, validacion y cierre.
- `docs/TECH_DEBT.md`: deuda tecnica real registrada.

## Fuentes de verdad

Las fuentes de verdad se leen en el orden definido por `AGENTS.md`:

1. `README.md`
2. `PLAN.md`
3. `AGENTS.md`
4. `docs/CURRENT_STATE.md`
5. `docs/requirements/`
6. `docs/DOMAIN_GUARDRAILS.md`
7. `docs/architecture/`
8. `docs/decisions/`
9. `docs/TECH_DEBT.md`

## Desarrollo por fases

El desarrollo se organiza por fases en `PLAN.md`. Cada fase debe ejecutarse dentro de su alcance exacto, sin adelantar trabajo de fases futuras y sin modificar decisiones historicas sin crear un nuevo ADR.

## ADRs

Las decisiones arquitectonicas iniciales estan en `docs/decisions/`. Los ADRs son historial tecnico versionado: no se borran; si una decision cambia, se crea un nuevo ADR y el anterior se marca como `Reemplazado` o `Deprecado`.

## Fuera de alcance actual

- Frontend.
- Aplicacion mobile.
- Workers.
- CLI.
- Migraciones SQL versionadas.
- Carpeta `migrations/`.
- Herramientas como `golang-migrate` o `goose`.
- Endpoints, modelos, servicios, handlers, repositorios o casos de uso implementados en esta fase documental.

## Proximos pasos

Ejecutar `Fase 1 — Bootstrap tecnico del proyecto` desde `PLAN.md`, creando solo la base tecnica minima definida para esa fase.
