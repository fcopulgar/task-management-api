# Plan activo

## Epica activa

Especificación e implementación incremental de `task-management-api`.

## Contexto

El proyecto cuenta con harness documental, especificación funcional inicial, stack técnico definido, ADRs iniciales y arquitectura objetivo. Todavía no existe implementación funcional.

## Alcance

Desarrollar una API REST backend para gestión de usuarios y tareas con autenticación, autorización por perfiles, sesiones revocables, cambio obligatorio de contraseña temporal y control de estados de tareas.

## Fuera de alcance

- Frontend.
- Mobile.
- Workers.
- CLI.
- Migraciones SQL versionadas.
- Lógica de negocio fuera de las fases definidas.
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
- `COMPLETADO ✅`: fase cerrada con validación y documentación actualizada.
- `BLOQUEADO ⛔`: fase detenida por una dependencia real.

## Fase 0 — Definición del proyecto

**Estado:** `COMPLETADO ✅`

### Objetivo

Completar requisitos, tecnologías, decisiones tecnicas, arquitectura objetivo y plan inicial.

### Alcance exacto

- Documentar alcance funcional.
- Documentar alcance no funcional.
- Documentar tecnologías elegidas.
- Crear ADRs iniciales.
- Crear plan de implementación por fases.
- Actualizar estado actual del proyecto.

### Tareas

- [x] Completar requisitos funcionales.
- [x] Completar requisitos no funcionales.
- [x] Completar decisiones tecnicas iniciales.
- [x] Crear ADRs iniciales.
- [x] Crear plan de implementación real.
- [x] Actualizar documentación vigente.

### Criterios de aceptacion

- Los requisitos principales estan documentados.
- Las tecnologías principales estan documentadas.
- Las decisiones arquitectónicas iniciales tienen ADR.
- Existe un plan de implementación por fases.
- No se ha implementado codigo.

### Entregables

- `docs/requirements/`
- `docs/architecture/`
- `docs/backend/`
- `docs/decisions/`
- `PLAN.md` actualizado.

## Fase 1 — Bootstrap técnico del proyecto

**Estado:** `COMPLETADO ✅`

### Objetivo

Crear estructura base de codigo, configuración minima, `go.mod`, servidor mínimo, Makefile, Docker Compose y validación base.

### Alcance exacto

- Inicializar módulo Go.
- Crear estructura minima compatible con arquitectura hexagonal.
- Crear servidor HTTP mínimo sin endpoints de negocio.
- Agregar configuración local minima.
- Agregar Makefile y Docker Compose para desarrollo local.
- Preparar validación base de build/test.

### Tareas

- [x] Crear `go.mod` con dependencias estrictamente necesarias para bootstrap.
- [x] Crear punto de entrada del servidor.
- [x] Crear configuración minima de entorno.
- [x] Crear Makefile con comandos base.
- [x] Crear Docker Compose con PostgreSQL y aplicación si corresponde.
- [x] Agregar validación minima de arranque sin lógica de negocio.
- [x] Actualizar documentación de ejecución local.

### Criterios de aceptacion

- El proyecto compila.
- Existe servidor mínimo ejecutable.
- No existen endpoints de negocio implementados.
- Docker Compose y Makefile reflejan el stack definido.
- La documentación indica comandos reales de validación.

### Entregables

- Estructura base Go.
- `go.mod`.
- Makefile.
- Docker Compose.
- Documentación de ejecución local actualizada.

## Fase 2 — Dominio y reglas centrales

**Estado:** `COMPLETADO ✅`

### Objetivo

Implementar entidades de dominio, roles, estados, transiciones, errores y tests unitarios de reglas criticas.

### Alcance exacto

- Entidades de dominio sin GORM.
- Roles `ADMIN`, `EXECUTOR`, `AUDITOR`.
- Estados de tarea y transiciones permitidas.
- Errores de dominio.
- Reglas criticas de usuarios, sesiones y tareas.

### Tareas

- [x] Implementar entidades de dominio de usuario, sesión, tarea y comentario.
- [x] Implementar tipos para roles y estados.
- [x] Implementar validación de transiciones de estado.
- [x] Implementar reglas de vencimiento y propiedad de tareas.
- [x] Implementar errores de dominio.
- [x] Crear tests unitarios de reglas criticas.

### Criterios de aceptacion

- El dominio no importa GORM ni HTTP.
- Las transiciones invalidas son rechazadas por tests.
- Las reglas criticas tienen cobertura unitaria.
- No se implementa persistencia ni endpoints reales.

### Entregables

- Paquetes de dominio.
- Tests unitarios de dominio.
- Documentación actualizada si cambian reglas.

## Fase 3 — Puertos de aplicación y contratos internos

**Estado:** `COMPLETADO ✅`

### Objetivo

Definir interfaces de repositorios, servicios de seguridad, token service y estructuras de input/output de casos de uso.

### Alcance exacto

- Puertos inbound y outbound.
- Contratos de repositorios.
- Contratos para hashing, tokens y sesiones.
- DTOs internos de casos de uso.
- Errores de aplicación.

### Tareas

- [x] Definir interfaces de repositorios requeridas.
- [x] Definir servicios de contraseña y token.
- [x] Definir contratos de sesión revocable.
- [x] Definir inputs y outputs internos por caso de uso.
- [x] Crear tests de contratos cuando corresponda.

### Criterios de aceptacion

- Application no importa GORM.
- Los contratos permiten implementar autenticación, usuarios y tareas en fases futuras.
- No hay handlers HTTP ni persistencia concreta.

### Entregables

- Puertos de aplicación.
- Contratos internos.
- Tests de contratos aplicables.

## Fase 4 — Persistencia GORM y AutoMigrate

**Estado:** `COMPLETADO ✅`

### Objetivo

Crear modelos GORM, mappers, repositorios, conexión PostgreSQL y AutoMigrate.

### Alcance exacto

- Modelos GORM separados del dominio.
- Mappers dominio/persistencia.
- Repositorios outbound.
- Conexión PostgreSQL.
- AutoMigrate para tablas requeridas.

### Tareas

- [x] Crear modelos GORM de usuarios, sesiones, tareas y comentarios.
- [x] Crear mappers entre dominio y modelos GORM.
- [x] Implementar repositorios con GORM.
- [x] Configurar conexión PostgreSQL.
- [x] Ejecutar AutoMigrate al iniciar segun configuración.
- [x] Agregar tests de persistencia si la infraestructura local lo permite.

### Criterios de aceptacion

- GORM vive solo en adapters outbound.
- No existe carpeta `migrations/`.
- AutoMigrate crea o actualiza el esquema esperado.
- El dominio sigue desacoplado de infraestructura.

### Entregables

- Adapter de persistencia GORM.
- Modelos GORM.
- Mappers.
- Configuración de base de datos.

## Fase 5 — Autenticación, sesiones y cambio de contraseña

**Estado:** `COMPLETADO ✅`

### Objetivo

Implementar login, emisión de JWT, sesiones revocables, logout y cambio obligatorio de contraseña.

### Alcance exacto

- Login.
- Emisión de JWT con `session_id`.
- Validación de sesión activa.
- Logout por revocación persistida.
- Cambio de contraseña.
- Restricción por `must_change_password`.

### Tareas

- [x] Implementar caso de uso de login.
- [x] Implementar creación o asociación de sesión.
- [x] Implementar emisión y validación de JWT.
- [x] Implementar logout con revocación persistida.
- [x] Implementar cambio de contraseña con bcrypt.
- [x] Implementar middleware de autenticación y restricción por contraseña temporal.
- [x] Agregar tests unitarios y HTTP aplicables.

### Criterios de aceptacion

- El JWT incluye `session_id`.
- Una sesión revocada no permite consumir endpoints protegidos.
- Un usuario inactivo no puede iniciar sesión.
- Un usuario con contraseña temporal pendiente solo puede cambiar contraseña y cerrar sesión.

### Entregables

- Endpoints de autenticación.
- Casos de uso de autenticación.
- Middleware de seguridad.
- Tests de autenticación.

## Fase 6 — Gestión de usuarios administrador

**Estado:** `COMPLETADO ✅`

### Objetivo

Implementar CRUD de usuarios para `ADMIN`, respetando restricciones de rol, contraseña temporal y desactivación lógica.

### Alcance exacto

- Crear usuarios `EXECUTOR` y `AUDITOR`.
- Listar usuarios.
- Ver detalle.
- Actualizar usuarios.
- Desactivar usuarios.
- Evitar exposicion de `password_hash`.

### Tareas

- [x] Implementar casos de uso de administración de usuarios.
- [x] Implementar handlers HTTP de usuarios.
- [x] Aplicar autorización `ADMIN`.
- [x] Rechazar creación de usuarios `ADMIN`.
- [x] Crear usuarios con contraseña temporal o `must_change_password = true`.
- [x] Implementar desactivación lógica.
- [x] Agregar tests de permisos y respuestas.

### Criterios de aceptacion

- Solo `ADMIN` puede gestionar usuarios.
- `ADMIN` no puede crear otros `ADMIN`.
- No se retorna `password_hash` por API.
- La eliminación se interpreta como desactivación lógica.

### Entregables

- Endpoints de usuarios.
- Casos de uso de usuarios.
- Tests de gestión de usuarios.

## Fase 7 — Gestión de tareas administrador y auditor

**Estado:** `COMPLETADO ✅`

### Objetivo

Implementar creación, listado, detalle, actualización y eliminación de tareas segun permisos y estado.

### Alcance exacto

- CRUD de tareas para `ADMIN` segun reglas.
- Lectura de tareas para `AUDITOR`.
- Asignación solo a `EXECUTOR`.
- Restricción de actualización y eliminación por estado `ASSIGNED`.

### Tareas

- [x] Implementar casos de uso de tareas para administrador.
- [x] Implementar lectura de tareas para auditor.
- [x] Implementar handlers HTTP compartidos segun permisos.
- [x] Validar asignación a usuarios `EXECUTOR`.
- [x] Rechazar actualización o eliminación fuera de `ASSIGNED`.
- [x] Agregar tests de permisos, asignación y estados.

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

**Estado:** `COMPLETADO ✅`

### Objetivo

Implementar listado de tareas propias, detalle, cambio de estado y comentarios sobre tareas vencidas.

### Alcance exacto

- Listado de tareas propias.
- Detalle de tarea propia.
- Cambio de estado con flujo permitido.
- Bloqueo de cambio de estado en tareas vencidas.
- Comentarios en tareas vencidas propias.

### Tareas

- [x] Implementar casos de uso de tareas propias.
- [x] Implementar cambio de estado por ejecutor.
- [x] Implementar comentario de tarea vencida propia.
- [x] Implementar handlers `/me/tasks`.
- [x] Rechazar acceso a tareas de otros ejecutores.
- [x] Agregar tests de propiedad, vencimiento y transiciones.

### Criterios de aceptacion

- `EXECUTOR` solo ve tareas propias.
- `EXECUTOR` no modifica datos generales de tarea.
- `EXECUTOR` no cambia estado de tareas vencidas.
- `EXECUTOR` puede comentar tareas vencidas propias.

### Entregables

- Endpoints de ejecutor.
- Casos de uso de ejecutor.
- Tests de flujo de ejecutor.

## Fase 9 — Tests HTTP, hardening y documentación final

**Estado:** `COMPLETADO ✅`

### Objetivo

Completar tests de integración HTTP minimos, smoke tests, documentación de endpoints, validaciones finales y limpieza de deuda.

### Alcance exacto

- Tests HTTP representativos.
- Smoke tests reales.
- Revisión de seguridad basica.
- Documentación final de endpoints implementados.
- Actualización de deuda técnica real.

### Tareas

- [x] Completar tests HTTP de flujos criticos.
- [x] Definir y ejecutar smoke tests reales.
- [x] Revisar errores y respuestas de seguridad.
- [x] Actualizar documentación de API segun implementación final.
- [x] Actualizar checklist de producción.
- [x] Registrar deuda técnica real pendiente.

### Criterios de aceptacion

- Los flujos criticos tienen cobertura HTTP minima.
- Los smoke tests estan documentados y ejecutados.
- La documentación no promete funcionalidades inexistentes.
- La deuda técnica real queda registrada.

### Entregables

- Tests HTTP.
- Smoke tests documentados.
- Documentación final vigente.
- Checklist y deuda técnica actualizados.
