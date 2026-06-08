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

## Planificado

- Backend en Go con router `chi`.
- Persistencia en PostgreSQL con GORM y `AutoMigrate`.
- Autenticacion JWT con `session_id`.
- Sesiones revocables persistidas en base de datos.
- Hash de contrasenas con bcrypt.
- Tests con `testing`, `testify` y `httptest` cuando corresponda.
- Desarrollo local futuro con Docker Compose y Makefile.

## No implementado todavia

- Codigo de aplicacion.
- Endpoints reales.
- Modelos de dominio o GORM.
- Casos de uso, handlers, repositorios y servicios.
- Docker Compose, Makefile y tests ejecutables.

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
- Verificacion de ausencia de codigo de aplicacion.

## Que no se pudo validar

- Build.
- Tests.
- Ejecucion local.
- Integraciones.
- Endpoints.

Estos puntos no se pudieron validar porque la implementacion esta `PENDIENTE DE IMPLEMENTACION`.
