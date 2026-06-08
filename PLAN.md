# Plan activo

## Epica activa

Especificacion e implementacion incremental de `task-management-api`.

## Contexto

El proyecto cuenta con harness documental, especificacion funcional inicial, stack tecnico definido, ADRs iniciales y arquitectura objetivo. Todavia no existe implementacion funcional.

## Alcance

Desarrollar una API REST backend para gestion de usuarios y tareas con autenticacion, autorizacion por perfiles, sesiones revocables, cambio obligatorio de contrasena temporal y control de estados de tareas.

## Fuera de alcance

- Frontend.
- Mobile.
- Workers.
- CLI.
- Migraciones SQL versionadas.
- Logica de negocio fuera de las fases definidas.
- Cambios de arquitectura sin ADR nuevo.

## Fuentes de verdad

- `README.md`
- `PLAN.md`
- `AGENTS.md`
- `docs/CURRENT_STATE.md`
- `docs/requirements/`
- `docs/DOMAIN_GUARDRAILS.md`
- `docs/architecture/`
- `docs/decisions/`
- `docs/TECH_DEBT.md`

## Convenciones de estado

- `PENDIENTE ⬜`: fase no iniciada.
- `EN PROGRESO 🟨`: fase iniciada y no cerrada.
- `COMPLETADO ✅`: fase cerrada con validacion y documentacion actualizada.
- `BLOQUEADO ⛔`: fase detenida por una dependencia real.

## Fase 0 — Definicion del proyecto

**Estado:** `COMPLETADO ✅`

### Objetivo

Completar requisitos, tecnologias, decisiones tecnicas, arquitectura objetivo y plan inicial.

### Alcance exacto

- Documentar alcance funcional.
- Documentar alcance no funcional.
- Documentar tecnologias elegidas.
- Crear ADRs iniciales.
- Crear plan de implementacion por fases.
- Actualizar estado actual del proyecto.

### Tareas

- [x] Completar requisitos funcionales.
- [x] Completar requisitos no funcionales.
- [x] Completar decisiones tecnicas iniciales.
- [x] Crear ADRs iniciales.
- [x] Crear plan de implementacion real.
- [x] Actualizar documentacion vigente.

### Criterios de aceptacion

- Los requisitos principales estan documentados.
- Las tecnologias principales estan documentadas.
- Las decisiones arquitectonicas iniciales tienen ADR.
- Existe un plan de implementacion por fases.
- No se ha implementado codigo.

### Entregables

- `docs/requirements/`
- `docs/architecture/`
- `docs/backend/`
- `docs/decisions/`
- `PLAN.md` actualizado.

## Fase 1 — Bootstrap tecnico del proyecto

**Estado:** `COMPLETADO ✅`

### Objetivo

Crear estructura base de codigo, configuracion minima, `go.mod`, servidor minimo, Makefile, Docker Compose y validacion base.

### Alcance exacto

- Inicializar modulo Go.
- Crear estructura minima compatible con arquitectura hexagonal.
- Crear servidor HTTP minimo sin endpoints de negocio.
- Agregar configuracion local minima.
- Agregar Makefile y Docker Compose para desarrollo local.
- Preparar validacion base de build/test.

### Tareas

- [x] Crear `go.mod` con dependencias estrictamente necesarias para bootstrap.
- [x] Crear punto de entrada del servidor.
- [x] Crear configuracion minima de entorno.
- [x] Crear Makefile con comandos base.
- [x] Crear Docker Compose con PostgreSQL y aplicacion si corresponde.
- [x] Agregar validacion minima de arranque sin logica de negocio.
- [x] Actualizar documentacion de ejecucion local.

### Criterios de aceptacion

- El proyecto compila.
- Existe servidor minimo ejecutable.
- No existen endpoints de negocio implementados.
- Docker Compose y Makefile reflejan el stack definido.
- La documentacion indica comandos reales de validacion.

### Entregables

- Estructura base Go.
- `go.mod`.
- Makefile.
- Docker Compose.
- Documentacion de ejecucion local actualizada.

## Fase 2 — Dominio y reglas centrales

**Estado:** `PENDIENTE ⬜`

### Objetivo

Implementar entidades de dominio, roles, estados, transiciones, errores y tests unitarios de reglas criticas.

### Alcance exacto

- Entidades de dominio sin GORM.
- Roles `ADMIN`, `EXECUTOR`, `AUDITOR`.
- Estados de tarea y transiciones permitidas.
- Errores de dominio.
- Reglas criticas de usuarios, sesiones y tareas.

### Tareas

- [ ] Implementar entidades de dominio de usuario, sesion, tarea y comentario.
- [ ] Implementar tipos para roles y estados.
- [ ] Implementar validacion de transiciones de estado.
- [ ] Implementar reglas de vencimiento y propiedad de tareas.
- [ ] Implementar errores de dominio.
- [ ] Crear tests unitarios de reglas criticas.

### Criterios de aceptacion

- El dominio no importa GORM ni HTTP.
- Las transiciones invalidas son rechazadas por tests.
- Las reglas criticas tienen cobertura unitaria.
- No se implementa persistencia ni endpoints reales.

### Entregables

- Paquetes de dominio.
- Tests unitarios de dominio.
- Documentacion actualizada si cambian reglas.

## Fase 3 — Puertos de aplicacion y contratos internos

**Estado:** `PENDIENTE ⬜`

### Objetivo

Definir interfaces de repositorios, servicios de seguridad, token service y estructuras de input/output de casos de uso.

### Alcance exacto

- Puertos inbound y outbound.
- Contratos de repositorios.
- Contratos para hashing, tokens y sesiones.
- DTOs internos de casos de uso.
- Errores de aplicacion.

### Tareas

- [ ] Definir interfaces de repositorios requeridas.
- [ ] Definir servicios de contrasena y token.
- [ ] Definir contratos de sesion revocable.
- [ ] Definir inputs y outputs internos por caso de uso.
- [ ] Crear tests de contratos cuando corresponda.

### Criterios de aceptacion

- Application no importa GORM.
- Los contratos permiten implementar autenticacion, usuarios y tareas en fases futuras.
- No hay handlers HTTP ni persistencia concreta.

### Entregables

- Puertos de aplicacion.
- Contratos internos.
- Tests de contratos aplicables.

## Fase 4 — Persistencia GORM y AutoMigrate

**Estado:** `PENDIENTE ⬜`

### Objetivo

Crear modelos GORM, mappers, repositorios, conexion PostgreSQL y AutoMigrate.

### Alcance exacto

- Modelos GORM separados del dominio.
- Mappers dominio/persistencia.
- Repositorios outbound.
- Conexion PostgreSQL.
- AutoMigrate para tablas requeridas.

### Tareas

- [ ] Crear modelos GORM de usuarios, sesiones, tareas y comentarios.
- [ ] Crear mappers entre dominio y modelos GORM.
- [ ] Implementar repositorios con GORM.
- [ ] Configurar conexion PostgreSQL.
- [ ] Ejecutar AutoMigrate al iniciar segun configuracion.
- [ ] Agregar tests de persistencia si la infraestructura local lo permite.

### Criterios de aceptacion

- GORM vive solo en adapters outbound.
- No existe carpeta `migrations/`.
- AutoMigrate crea o actualiza el esquema esperado.
- El dominio sigue desacoplado de infraestructura.

### Entregables

- Adapter de persistencia GORM.
- Modelos GORM.
- Mappers.
- Configuracion de base de datos.

## Fase 5 — Autenticacion, sesiones y cambio de contrasena

**Estado:** `PENDIENTE ⬜`

### Objetivo

Implementar login, emision de JWT, sesiones revocables, logout y cambio obligatorio de contrasena.

### Alcance exacto

- Login.
- Emision de JWT con `session_id`.
- Validacion de sesion activa.
- Logout por revocacion persistida.
- Cambio de contrasena.
- Restriccion por `must_change_password`.

### Tareas

- [ ] Implementar caso de uso de login.
- [ ] Implementar creacion o asociacion de sesion.
- [ ] Implementar emision y validacion de JWT.
- [ ] Implementar logout con revocacion persistida.
- [ ] Implementar cambio de contrasena con bcrypt.
- [ ] Implementar middleware de autenticacion y restriccion por contrasena temporal.
- [ ] Agregar tests unitarios y HTTP aplicables.

### Criterios de aceptacion

- El JWT incluye `session_id`.
- Una sesion revocada no permite consumir endpoints protegidos.
- Un usuario inactivo no puede iniciar sesion.
- Un usuario con contrasena temporal pendiente solo puede cambiar contrasena y cerrar sesion.

### Entregables

- Endpoints de autenticacion.
- Casos de uso de autenticacion.
- Middleware de seguridad.
- Tests de autenticacion.

## Fase 6 — Gestion de usuarios administrador

**Estado:** `PENDIENTE ⬜`

### Objetivo

Implementar CRUD de usuarios para `ADMIN`, respetando restricciones de rol, contrasena temporal y desactivacion logica.

### Alcance exacto

- Crear usuarios `EXECUTOR` y `AUDITOR`.
- Listar usuarios.
- Ver detalle.
- Actualizar usuarios.
- Desactivar usuarios.
- Evitar exposicion de `password_hash`.

### Tareas

- [ ] Implementar casos de uso de administracion de usuarios.
- [ ] Implementar handlers HTTP de usuarios.
- [ ] Aplicar autorizacion `ADMIN`.
- [ ] Rechazar creacion de usuarios `ADMIN`.
- [ ] Crear usuarios con contrasena temporal o `must_change_password = true`.
- [ ] Implementar desactivacion logica.
- [ ] Agregar tests de permisos y respuestas.

### Criterios de aceptacion

- Solo `ADMIN` puede gestionar usuarios.
- `ADMIN` no puede crear otros `ADMIN`.
- No se retorna `password_hash` por API.
- La eliminacion se interpreta como desactivacion logica.

### Entregables

- Endpoints de usuarios.
- Casos de uso de usuarios.
- Tests de gestion de usuarios.

## Fase 7 — Gestion de tareas administrador y auditor

**Estado:** `PENDIENTE ⬜`

### Objetivo

Implementar creacion, listado, detalle, actualizacion y eliminacion de tareas segun permisos y estado.

### Alcance exacto

- CRUD de tareas para `ADMIN` segun reglas.
- Lectura de tareas para `AUDITOR`.
- Asignacion solo a `EXECUTOR`.
- Restriccion de actualizacion y eliminacion por estado `ASSIGNED`.

### Tareas

- [ ] Implementar casos de uso de tareas para administrador.
- [ ] Implementar lectura de tareas para auditor.
- [ ] Implementar handlers HTTP compartidos segun permisos.
- [ ] Validar asignacion a usuarios `EXECUTOR`.
- [ ] Rechazar actualizacion o eliminacion fuera de `ASSIGNED`.
- [ ] Agregar tests de permisos, asignacion y estados.

### Criterios de aceptacion

- `ADMIN` puede crear tareas asignadas a ejecutores.
- `AUDITOR` solo puede leer tareas.
- No se asignan tareas a `ADMIN` ni `AUDITOR`.
- Las tareas no `ASSIGNED` no pueden ser actualizadas ni eliminadas por `ADMIN`.

### Entregables

- Endpoints de tareas para administrador y auditor.
- Casos de uso de tareas.
- Tests de permisos y reglas de estado.

## Fase 8 — Flujos de ejecutor

**Estado:** `PENDIENTE ⬜`

### Objetivo

Implementar listado de tareas propias, detalle, cambio de estado y comentarios sobre tareas vencidas.

### Alcance exacto

- Listado de tareas propias.
- Detalle de tarea propia.
- Cambio de estado con flujo permitido.
- Bloqueo de cambio de estado en tareas vencidas.
- Comentarios en tareas vencidas propias.

### Tareas

- [ ] Implementar casos de uso de tareas propias.
- [ ] Implementar cambio de estado por ejecutor.
- [ ] Implementar comentario de tarea vencida propia.
- [ ] Implementar handlers `/me/tasks`.
- [ ] Rechazar acceso a tareas de otros ejecutores.
- [ ] Agregar tests de propiedad, vencimiento y transiciones.

### Criterios de aceptacion

- `EXECUTOR` solo ve tareas propias.
- `EXECUTOR` no modifica datos generales de tarea.
- `EXECUTOR` no cambia estado de tareas vencidas.
- `EXECUTOR` puede comentar tareas vencidas propias.

### Entregables

- Endpoints de ejecutor.
- Casos de uso de ejecutor.
- Tests de flujo de ejecutor.

## Fase 9 — Tests HTTP, hardening y documentacion final

**Estado:** `PENDIENTE ⬜`

### Objetivo

Completar tests de integracion HTTP minimos, smoke tests, documentacion de endpoints, validaciones finales y limpieza de deuda.

### Alcance exacto

- Tests HTTP representativos.
- Smoke tests reales.
- Revision de seguridad basica.
- Documentacion final de endpoints implementados.
- Actualizacion de deuda tecnica real.

### Tareas

- [ ] Completar tests HTTP de flujos criticos.
- [ ] Definir y ejecutar smoke tests reales.
- [ ] Revisar errores y respuestas de seguridad.
- [ ] Actualizar documentacion de API segun implementacion final.
- [ ] Actualizar checklist de produccion.
- [ ] Registrar deuda tecnica real pendiente.

### Criterios de aceptacion

- Los flujos criticos tienen cobertura HTTP minima.
- Los smoke tests estan documentados y ejecutados.
- La documentacion no promete funcionalidades inexistentes.
- La deuda tecnica real queda registrada.

### Entregables

- Tests HTTP.
- Smoke tests documentados.
- Documentacion final vigente.
- Checklist y deuda tecnica actualizados.
