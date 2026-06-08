# Especificación del proyecto, decisiones técnicas y plan por fases

El repositorio ya tiene un harness documental base.

Ahora quiero completar la especificación del proyecto `task-management-api`, documentar las tecnologías elegidas, registrar decisiones arquitectónicas iniciales mediante ADRs versionados y crear un `PLAN.md` de implementación por fases.

Este prompt NO debe implementar código.

Este prompt NO debe crear lógica de negocio.

Este prompt NO debe crear endpoints reales.

Este prompt NO debe crear modelos, servicios, controladores, handlers ni repositorios.

Este prompt solo debe llenar y ajustar documentación para que el proyecto quede listo para desarrollarse posteriormente por fases.

Trata este proyecto como un backend profesional real, no como una maqueta, demo, challenge, ejercicio académico, evaluación, entrevista o prueba.

No escribas en ninguna documentación que esto corresponde a una prueba, challenge, entrevista, postulación, evaluación o ejercicio.

---

# 1. Información del proyecto

## Nombre del proyecto

```text
task-management-api
```

## Descripción del producto

`task-management-api` es una API REST para gestión de usuarios y tareas con autenticación, autorización por perfiles, cambio obligatorio de contraseña temporal, cierre de sesión y control de estados de tareas.

El sistema permite que administradores gestionen usuarios y tareas, que ejecutores trabajen sobre sus tareas asignadas, y que auditores consulten el estado de tareas de cualquier usuario.

El foco del proyecto es mantener una base backend clara, segura, testeable y extensible, evitando sobredimensionar la solución.

---

# 2. Tecnologías definidas

## Lenguaje

```text
Go
```

## Router HTTP

```text
chi
```

## ORM

```text
GORM
```

## Base de datos

```text
PostgreSQL
```

## Inicialización de esquema

```text
GORM AutoMigrate
```

No usar migraciones SQL versionadas en esta etapa.

No crear carpeta:

```text
migrations/
```

No agregar herramientas como:

```text
golang-migrate
goose
```

No agregar comandos:

```text
make migrate-up
make migrate-down
```

## Autenticación

```text
JWT con session_id incluido en claims
```

## Cierre de sesión

```text
Revocación de sesión persistida en base de datos
```

## Hash de contraseñas

```text
bcrypt
```

## Testing

```text
testing
testify
httptest cuando corresponda
```

## Desarrollo local futuro

```text
Docker Compose
Makefile
```

## Documentación

```text
Markdown
Mermaid
ADRs versionados
```

---

# 3. Decisiones técnicas iniciales

Documentar estas decisiones en la documentación correspondiente y en ADRs versionados:

* Usar Go como lenguaje principal.
* Usar arquitectura hexagonal minimalista.
* Usar `chi` como router HTTP.
* Usar PostgreSQL como base de datos.
* Usar GORM como ORM.
* Usar GORM `AutoMigrate` para inicialización de esquema.
* No usar migraciones SQL versionadas en esta etapa.
* Mantener dominio desacoplado de GORM.
* Mantener capa de aplicación desacoplada de GORM.
* Encapsular GORM solo en adapters outbound de persistencia.
* Usar JWT para autenticación.
* Incluir `session_id` dentro del token.
* Implementar logout mediante revocación de sesión persistida.
* Hashear contraseñas con bcrypt.
* Mantener handlers HTTP delgados.
* Concentrar reglas de negocio en dominio y casos de uso.
* Organizar el desarrollo por fases usando `PLAN.md`.
* Mantener toda la documentación en español.
* Usar ADRs versionados como historial técnico.
* No borrar ADRs si una decisión cambia; crear un nuevo ADR y marcar el anterior como reemplazado o deprecado.

---

# 4. Restricciones explícitas

* No implementar código en este prompt.
* No crear archivos de aplicación.
* No crear lógica de negocio.
* No crear endpoints reales.
* No crear modelos reales.
* No crear handlers.
* No crear services.
* No crear repositories.
* No crear use cases.
* No agregar dependencias.
* No ejecutar instalación de paquetes.
* No crear archivos de lock.
* No crear archivos generados por frameworks.
* No modificar estructura de código si existiera.
* No crear frontend.
* No crear mobile.
* No crear workers.
* No crear `AI_USAGE.md`.
* No crear migraciones SQL versionadas.
* No crear carpeta `migrations/`.
* No mencionar en documentación que esto es una prueba, challenge, postulación, entrevista, evaluación o ejercicio.
* No escribir documentación en inglés.

Si un requisito no está claro, documentarlo como:

```text
PENDIENTE DE DEFINICIÓN
```

Si algo todavía no se implementa, documentarlo como:

```text
PENDIENTE DE IMPLEMENTACIÓN
```

Si algo no se puede confirmar, documentarlo como:

```text
NO CONFIRMADO
```

---

# 5. Alcance funcional

El sistema debe contemplar tres perfiles:

```text
ADMIN
EXECUTOR
AUDITOR
```

Mapeo conceptual:

```text
ADMIN = Administrador
EXECUTOR = Ejecutor
AUDITOR = Auditor
```

---

## 5.1 Autenticación

El sistema debe permitir:

* login de usuario
* identificación del usuario autenticado
* identificación del perfil del usuario autenticado
* autorización según perfil
* actualización de contraseña para cualquier perfil
* cierre de sesión

Reglas:

* el login debe entregar un token
* el token debe permitir identificar usuario, perfil y sesión
* el token debe incluir `session_id`
* el cierre de sesión debe invalidar la sesión actual
* las contraseñas deben almacenarse hasheadas
* un usuario inactivo no puede iniciar sesión
* los usuarios creados con contraseña temporal deben cambiarla en el primer ingreso
* un usuario con contraseña temporal pendiente solo debe poder cambiar contraseña y cerrar sesión

---

## 5.2 Perfil ADMIN

Un usuario `ADMIN` puede administrar usuarios y tareas.

### Gestión de usuarios

Puede:

* crear usuarios
* listar usuarios
* ver detalle de usuario
* actualizar usuarios
* desactivar usuarios

Reglas:

* un `ADMIN` solo puede crear usuarios `EXECUTOR` o `AUDITOR`
* un `ADMIN` no puede crear otros usuarios `ADMIN`
* los usuarios creados por un `ADMIN` nacen con contraseña temporal
* los usuarios creados por un `ADMIN` deben cambiar la contraseña en su primer ingreso
* no se debe retornar nunca el hash de contraseña por API
* se debe preferir desactivación lógica antes que eliminación física

### Gestión de tareas

Puede:

* crear tareas
* listar tareas
* ver detalle de tarea
* actualizar tareas
* eliminar tareas

Reglas:

* una tarea debe tener al menos título, descripción y fecha de vencimiento
* una tarea debe ser asignada a un usuario `EXECUTOR`
* no se puede asignar una tarea a un usuario `ADMIN`
* no se puede asignar una tarea a un usuario `AUDITOR`
* una tarea nueva nace en estado `ASSIGNED`
* un `ADMIN` solo puede actualizar o eliminar tareas en estado `ASSIGNED`

---

## 5.3 Perfil EXECUTOR

Un usuario `EXECUTOR` puede trabajar sobre sus tareas asignadas.

Puede:

* listar sus tareas asignadas
* ver detalle de sus tareas asignadas
* actualizar el estado de sus tareas
* comentar tareas vencidas propias

Reglas:

* un `EXECUTOR` solo puede ver tareas propias
* un `EXECUTOR` no puede ver tareas de otros ejecutores
* un `EXECUTOR` no puede modificar datos generales de la tarea
* un `EXECUTOR` solo puede modificar el estado mediante el flujo permitido
* un `EXECUTOR` no puede cambiar el estado de una tarea vencida
* un `EXECUTOR` solo puede comentar tareas vencidas propias

---

## 5.4 Perfil AUDITOR

Un usuario `AUDITOR` puede consultar información de tareas.

Puede:

* listar tareas de cualquier usuario
* ver detalle de tareas
* ver estado de tareas

Reglas:

* un `AUDITOR` no puede crear usuarios
* un `AUDITOR` no puede modificar usuarios
* un `AUDITOR` no puede crear tareas
* un `AUDITOR` no puede modificar tareas
* un `AUDITOR` no puede eliminar tareas

---

# 6. Estados de tareas

Usar internamente estos estados:

```text
ASSIGNED
STARTED
WAITING
FINISHED_SUCCESS
FINISHED_ERROR
```

Mapeo conceptual:

```text
ASSIGNED = Asignado
STARTED = Iniciado
WAITING = En espera
FINISHED_SUCCESS = Finalizado con éxito
FINISHED_ERROR = Finalizado con error
```

Flujo permitido:

```text
ASSIGNED -> STARTED

STARTED -> WAITING
STARTED -> FINISHED_SUCCESS
STARTED -> FINISHED_ERROR

WAITING -> WAITING
WAITING -> FINISHED_SUCCESS
WAITING -> FINISHED_ERROR
```

Reglas:

* una tarea nueva siempre nace en `ASSIGNED`
* los estados terminales son `FINISHED_SUCCESS` y `FINISHED_ERROR`
* las transiciones inválidas deben ser rechazadas
* el ejecutor no puede cambiar el estado de una tarea vencida
* el administrador no puede actualizar ni eliminar tareas que no estén en `ASSIGNED`

---

# 7. Entidades principales esperadas

Documentar estas entidades en `docs/backend/domain.md` y `docs/backend/data-model.md`.

No implementar modelos todavía.

## Usuario

Campos esperados:

```text
id
name
email
password_hash
role
must_change_password
active
created_at
updated_at
```

Reglas:

* `email` debe ser único
* `role` debe ser `ADMIN`, `EXECUTOR` o `AUDITOR`
* `password_hash` nunca debe exponerse por API
* `must_change_password` indica si el usuario debe cambiar contraseña
* `active = false` representa usuario desactivado

## Sesión

Campos esperados:

```text
id
user_id
revoked_at
expires_at
created_at
```

Reglas:

* el JWT debe incluir `session_id`
* una sesión revocada no debe permitir consumir endpoints protegidos
* logout debe marcar la sesión como revocada

## Tarea

Campos esperados:

```text
id
title
description
due_at
status
assignee_id
created_by
created_at
updated_at
```

Reglas:

* `assignee_id` debe corresponder a un usuario `EXECUTOR`
* una tarea nueva nace como `ASSIGNED`
* no se puede actualizar ni eliminar si su estado no es `ASSIGNED`

## Comentario de tarea

Campos esperados:

```text
id
task_id
user_id
comment
created_at
```

Reglas:

* solo ejecutores pueden comentar tareas propias vencidas
* el comentario debe quedar asociado a la tarea y al usuario

---

# 8. Endpoints planificados

Documentar estos endpoints en `docs/backend/api.md`.

No implementarlos todavía.

## Auth

```text
POST /auth/login
POST /auth/logout
PUT  /auth/password
```

## Usuarios — ADMIN

```text
POST   /users
GET    /users
GET    /users/{id}
PUT    /users/{id}
DELETE /users/{id}
```

Reglas:

* solo `ADMIN`
* `POST /users` solo permite crear `EXECUTOR` o `AUDITOR`
* `DELETE /users/{id}` debe interpretarse preferentemente como desactivación lógica

## Tareas — ADMIN y AUDITOR

```text
POST   /tasks
GET    /tasks
GET    /tasks/{id}
PUT    /tasks/{id}
DELETE /tasks/{id}
```

Reglas:

* `POST /tasks`: solo `ADMIN`
* `PUT /tasks/{id}`: solo `ADMIN` y solo si estado `ASSIGNED`
* `DELETE /tasks/{id}`: solo `ADMIN` y solo si estado `ASSIGNED`
* `GET /tasks`: `ADMIN` y `AUDITOR` pueden listar todas
* `GET /tasks/{id}`: `ADMIN` y `AUDITOR` pueden ver detalle

## Tareas — EXECUTOR

```text
GET   /me/tasks
GET   /me/tasks/{id}
PATCH /me/tasks/{id}/status
POST  /me/tasks/{id}/comments
```

Reglas:

* solo `EXECUTOR`
* solo sobre tareas propias
* no puede cambiar estado si la tarea está vencida
* solo puede comentar tarea vencida propia

---

# 9. Archivos que debes actualizar

Actualizar, como mínimo:

```text
README.md
AGENTS.md
PLAN.md

docs/CURRENT_STATE.md
docs/DOMAIN_GUARDRAILS.md
docs/AGENT_WORKFLOW.md
docs/TECH_DEBT.md
docs/PRODUCTION_CHECKLIST.md
docs/SMOKE_TESTS.md

docs/requirements/README.md
docs/requirements/product-scope.md
docs/requirements/functional-requirements.md
docs/requirements/non-functional-requirements.md
docs/requirements/open-questions.md

docs/architecture/README.md
docs/architecture/overview.md
docs/architecture/module-boundaries.md
docs/architecture/data-flow.md
docs/architecture/security.md

docs/decisions/README.md

docs/plans/README.md

docs/standards/documentation-style.md
docs/standards/definition-of-done.md
docs/standards/validation.md
```

Crear documentación adicional solo si está justificada por el stack y el alcance.

Crear si no existen:

```text
docs/backend/
docs/backend/README.md
docs/backend/domain.md
docs/backend/api.md
docs/backend/data-model.md
docs/backend/services.md
docs/backend/testing.md
```

No crear frontend, mobile, workers, CLI ni data docs porque no aplican al alcance actual.

---

# 10. README.md

Actualizar `README.md` para que describa el proyecto real.

Debe incluir:

* nombre del proyecto
* descripción profesional del producto
* alcance funcional resumido
* stack definido
* estado actual
* estructura documental
* fuentes de verdad
* cómo se desarrolla por fases
* dónde están los ADRs
* dónde está el plan activo
* qué está fuera de alcance
* próximos pasos

No incluir instrucciones de ejecución si todavía no hay implementación.

No prometer funcionalidades implementadas si solo están planificadas.

Distinguir claramente:

```text
IMPLEMENTADO
PLANIFICADO
PENDIENTE DE DEFINICIÓN
PENDIENTE DE IMPLEMENTACIÓN
```

---

# 11. Requisitos

Actualizar `docs/requirements/`.

## `product-scope.md`

Debe incluir:

* propósito del producto
* problema que resuelve
* usuarios/actores
* alcance inicial
* fuera de alcance
* supuestos
* pendientes

## `functional-requirements.md`

Convertir los requisitos entregados en una especificación clara.

Agrupar por dominio, actor, módulo o flujo.

Cada requisito debe tener:

```md
### RF-XXX — Nombre del requisito

**Estado:** PLANIFICADO

**Descripción:** ...

**Reglas:**

- ...

**Criterios de aceptación:**

- ...
```

Requisitos mínimos esperados:

* RF-001 — Login de usuario
* RF-002 — Logout de usuario
* RF-003 — Cambio de contraseña
* RF-004 — Restricción por contraseña temporal
* RF-005 — Gestión de usuarios por administrador
* RF-006 — Restricción de creación de administradores
* RF-007 — Gestión de tareas por administrador
* RF-008 — Asignación de tareas solo a ejecutores
* RF-009 — Restricción de actualización/eliminación por estado
* RF-010 — Listado de tareas propias para ejecutor
* RF-011 — Detalle de tarea propia para ejecutor
* RF-012 — Actualización de estado por ejecutor
* RF-013 — Bloqueo de actualización de tarea vencida
* RF-014 — Comentario sobre tarea vencida
* RF-015 — Visualización de tareas por auditor
* RF-016 — Control de transiciones de estado

## `non-functional-requirements.md`

Documentar requisitos como:

* seguridad
* mantenibilidad
* arquitectura
* testing
* documentación
* ejecución local
* trazabilidad de decisiones
* claridad de errores
* separación de responsabilidades

Cada requisito debe tener:

```md
### RNF-XXX — Nombre del requisito

**Estado:** PLANIFICADO

**Descripción:** ...

**Criterios de aceptación:**

- ...
```

Requisitos mínimos esperados:

* RNF-001 — Arquitectura hexagonal minimalista
* RNF-002 — Dominio desacoplado de infraestructura
* RNF-003 — Persistencia con PostgreSQL y GORM
* RNF-004 — Inicialización con AutoMigrate
* RNF-005 — Seguridad de contraseñas con bcrypt
* RNF-006 — Autenticación con JWT y sesiones revocables
* RNF-007 — Documentación en español
* RNF-008 — ADRs versionados
* RNF-009 — Tests de reglas críticas
* RNF-010 — Handlers delgados

## `open-questions.md`

Registrar dudas reales.

No inventar dudas innecesarias.

Preguntas sugeridas si aplican:

* ¿La eliminación de usuarios debe ser siempre lógica?
* ¿El administrador puede cambiar la contraseña de un usuario?
* ¿El auditor puede filtrar tareas por usuario, estado o fecha?
* ¿El comentario de tarea vencida puede agregarse más de una vez?
* ¿El estado `WAITING -> WAITING` debe registrarse como evento o simplemente permitirse como operación idempotente?
* ¿La contraseña temporal será generada automáticamente o recibida en el request de creación de usuario?

---

# 12. Arquitectura

Actualizar `docs/architecture/`.

Debe documentar la arquitectura objetivo sin implementarla.

Debe incluir:

* arquitectura hexagonal minimalista
* módulos previstos
* límites entre módulos
* dependencias permitidas
* dependencias prohibidas
* flujos principales esperados
* seguridad esperada
* persistencia esperada
* uso de GORM solo en adapter outbound
* uso de AutoMigrate
* uso de JWT y sesiones revocables

Si agregas diagramas Mermaid, etiquetarlos como:

```text
DISEÑO OBJETIVO
```

No afirmar que algo está implementado si no lo está.

Diagramas deseables:

* arquitectura hexagonal objetivo
* flujo de autenticación objetivo
* flujo de autorización por rol objetivo
* flujo de logout objetivo
* flujo de cambio obligatorio de contraseña objetivo
* flujo de tareas por administrador objetivo
* flujo de tareas por ejecutor objetivo
* flujo de auditor objetivo
* estados de tarea objetivo

---

# 13. Guardrails de dominio

Actualizar `docs/DOMAIN_GUARDRAILS.md`.

Debe contener reglas que futuros agentes no deben romper.

Incluir como mínimo:

## Usuarios

* Solo `ADMIN` puede crear usuarios.
* `ADMIN` solo puede crear usuarios `EXECUTOR` o `AUDITOR`.
* `ADMIN` no puede crear otros `ADMIN`.
* Todo usuario creado por `ADMIN` nace con contraseña temporal o con `must_change_password = true`.
* Un usuario inactivo no puede iniciar sesión.
* Nunca retornar `password_hash` por API.

## Autenticación

* Login válido debe crear o asociar una sesión.
* JWT debe incluir `session_id`.
* Logout debe revocar la sesión.
* Un token con sesión revocada no debe ser aceptado.
* Usuarios con `must_change_password = true` solo pueden cambiar contraseña y cerrar sesión.

## Tareas

* Una tarea nueva nace en `ASSIGNED`.
* Solo se puede asignar una tarea a un usuario `EXECUTOR`.
* `ADMIN` solo puede actualizar o eliminar tareas en `ASSIGNED`.
* `EXECUTOR` solo puede listar y ver tareas propias.
* `EXECUTOR` solo puede cambiar estado de tareas propias.
* `EXECUTOR` no puede cambiar estado de una tarea vencida.
* `EXECUTOR` solo puede comentar tareas vencidas propias.
* `AUDITOR` solo puede leer tareas.
* Las transiciones de estado deben respetar el flujo definido.

## Arquitectura

* Dominio sin dependencias externas.
* Dominio no importa GORM.
* Application no importa GORM.
* Handlers no contienen reglas de negocio.
* Repositories no contienen reglas de negocio.
* GORM vive solo en adapters outbound.
* DTOs HTTP no deben ser usados como entidades de dominio.
* Modelos de dominio y modelos GORM deben mantenerse separados.

## Persistencia

* Usar PostgreSQL.
* Usar GORM.
* Usar `AutoMigrate`.
* No usar migraciones SQL versionadas en esta etapa.
* No crear carpeta `migrations/`.

---

# 14. ADRs versionados

Crear ADRs iniciales en:

```text
docs/decisions/
```

Usar numeración secuencial:

```text
0001-*.md
0002-*.md
0003-*.md
```

Crear estos ADRs iniciales:

```text
0001-use-go.md
0002-use-hexagonal-architecture.md
0003-use-chi-router.md
0004-use-postgresql.md
0005-use-gorm.md
0006-use-gorm-automigrate.md
0007-use-jwt-with-revocable-sessions.md
0008-use-bcrypt-for-password-hashing.md
```

Cada ADR debe usar este formato:

```md
# NNNN — Título de la decisión

**Estado:** Aceptado  
**Fecha:** YYYY-MM-DD

## Contexto

...

## Decisión

...

## Consecuencias

...

## Alternativas consideradas

...

## Relación con otros ADRs

...
```

Reglas:

* los ADRs no se borran
* si cambia una decisión, se crea un nuevo ADR
* el ADR anterior se marca como `Reemplazado` o `Deprecado`
* no reescribir la historia técnica

---

# 15. PLAN.md

Actualizar `PLAN.md` para crear un plan real de implementación por fases.

El plan debe ser corto, accionable y ordenado.

No debe implementar nada.

Debe contener:

* nombre de la épica activa
* contexto
* alcance
* fuera de alcance
* fuentes de verdad
* convenciones de estado
* fases numeradas
* tareas detalladas
* criterios de aceptación por fase
* entregables por fase

Cada fase debe usar este formato:

```md
## Fase X — Nombre claro de la fase

**Estado:** `PENDIENTE ⬜`

### Objetivo

...

### Alcance exacto

...

### Tareas

- [ ] tarea concreta
- [ ] otra tarea concreta

### Criterios de aceptación

- criterio verificable
- criterio verificable

### Entregables

- entregable concreto
- entregable concreto
```

Crear estas fases:

## Fase 1 — Bootstrap técnico del proyecto

Objetivo:

Crear estructura base de código, configuración mínima, `go.mod`, servidor mínimo, Makefile, Docker Compose y validación base.

No implementar lógica de negocio completa.

## Fase 2 — Dominio y reglas centrales

Objetivo:

Implementar entidades de dominio, roles, estados, transiciones, errores y tests unitarios de reglas críticas.

## Fase 3 — Puertos de aplicación y contratos internos

Objetivo:

Definir interfaces de repositorios, servicios de seguridad, token service y estructuras de input/output de casos de uso.

## Fase 4 — Persistencia GORM y AutoMigrate

Objetivo:

Crear modelos GORM, mappers, repositorios, conexión PostgreSQL y AutoMigrate.

## Fase 5 — Autenticación, sesiones y cambio de contraseña

Objetivo:

Implementar login, emisión de JWT, sesiones revocables, logout y cambio obligatorio de contraseña.

## Fase 6 — Gestión de usuarios administrador

Objetivo:

Implementar CRUD de usuarios para `ADMIN`, respetando restricciones de rol, contraseña temporal y desactivación lógica.

## Fase 7 — Gestión de tareas administrador y auditor

Objetivo:

Implementar creación, listado, detalle, actualización y eliminación de tareas según permisos y estado.

## Fase 8 — Flujos de ejecutor

Objetivo:

Implementar listado de tareas propias, detalle, cambio de estado y comentarios sobre tareas vencidas.

## Fase 9 — Tests HTTP, hardening y documentación final

Objetivo:

Completar tests de integración HTTP mínimos, smoke tests, documentación de endpoints, validaciones finales y limpieza de deuda.

Cada fase debe quedar suficientemente detallada para que un agente futuro pueda trabajar solo en esa fase sin adelantarse a la siguiente.

---

# 16. Estado actual

Actualizar `docs/CURRENT_STATE.md`.

Debe reflejar:

* que el harness existe
* que el producto ya está especificado
* que las tecnologías ya están definidas
* que los ADRs iniciales ya existen
* que existe un plan por fases
* que todavía no hay implementación funcional
* que los endpoints están planificados, no implementados
* que los modelos están planificados, no implementados
* qué validaciones son posibles
* qué no se pudo validar

No vender humo.

No marcar como implementado algo que solo fue documentado.

---

# 17. Deuda técnica

Actualizar `docs/TECH_DEBT.md`.

Registrar solo deuda real o riesgos concretos.

No registrar como deuda cosas que son decisiones explícitas del proyecto.

La ausencia de migraciones SQL versionadas no debe registrarse como deuda mientras exista un ADR aceptado para usar AutoMigrate.

Riesgos sugeridos si aplican:

* Definir con precisión si eliminación de usuarios será lógica o física.
* Definir política exacta de contraseña temporal.
* Definir formato final de errores HTTP.
* Definir si se agregarán filtros de tareas para auditor.
* Definir si comentarios de tareas vencidas pueden ser múltiples.

Cada ítem debe tener:

* descripción
* impacto
* ubicación o área afectada
* recomendación
* fase sugerida si aplica

---

# 18. Smoke tests y checklist

Actualizar:

```text
docs/SMOKE_TESTS.md
docs/PRODUCTION_CHECKLIST.md
```

Deben quedar como pruebas/checklists esperados para fases futuras.

Marcar claramente lo no implementado como:

```text
PENDIENTE DE IMPLEMENTACIÓN
```

Smoke tests esperados al final del proyecto:

* levantar PostgreSQL
* levantar API
* login administrador
* crear usuario ejecutor
* cambiar contraseña ejecutor
* crear usuario auditor
* crear tarea asignada a ejecutor
* listar tareas como ejecutor
* ver detalle de tarea propia
* cambiar estado de tarea
* intentar cambiar estado de tarea vencida
* comentar tarea vencida
* listar tareas como auditor
* intentar modificar tarea como auditor y recibir rechazo
* cerrar sesión
* verificar que token con sesión revocada no funciona

---

# 19. AGENTS.md y AGENT_WORKFLOW.md

Actualizar para que los agentes futuros sepan:

* qué leer primero
* cómo elegir una fase
* cómo reportar resultados
* cuándo crear ADRs
* cuándo actualizar estado actual
* cuándo registrar deuda técnica
* cómo evitar scope creep
* que no deben implementar fuera de la fase activa
* que no deben mezclar documentación en inglés y español
* que no deben mencionar prueba, challenge, entrevista, evaluación o ejercicio

El workflow debe permitir ejecutar fases futuras con un prompt pequeño como:

```text
Trabaja solo en Fase X — <nombre exacto>
```

---

# 20. Validación de esta tarea

Al terminar:

* verificar que no se implementó código
* verificar que no se agregaron dependencias
* verificar que no se crearon archivos generados por frameworks
* verificar que la documentación está en español
* verificar que los ADRs iniciales existen
* verificar que `PLAN.md` tiene fases claras y no implementa nada
* verificar que `CURRENT_STATE.md` no vende humo
* verificar que `README.md` no promete funcionalidades implementadas inexistentes
* verificar que no hay referencias a prueba, challenge, entrevista, postulación, evaluación o ejercicio académico

No ejecutar build ni tests si no hay implementación.

---

# 21. Entrega final

Responder con:

1. resumen de especificación documentada
2. archivos actualizados
3. ADRs creados
4. plan por fases creado
5. decisiones pendientes
6. deuda o riesgos registrados
7. validación realizada
8. siguiente fase recomendada

No implementar código.

No responder solo con recomendaciones.

Haz los cambios directamente en el repositorio.
