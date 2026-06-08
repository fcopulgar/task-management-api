# Requisitos funcionales

## Estado

`PLANIFICADO`

Los requisitos estan especificados para implementacion futura. Ningun endpoint esta implementado todavia.

### RF-001 — Login de usuario

**Estado:** PLANIFICADO

**Descripcion:** El sistema debe autenticar usuarios activos mediante credenciales validas y entregar un token.

**Reglas:**

- El token debe identificar usuario, perfil y sesion.
- El token debe incluir `session_id`.
- Un usuario inactivo no puede iniciar sesion.

**Criterios de aceptacion:**

- Credenciales validas de usuario activo generan token.
- Credenciales invalidas son rechazadas.
- Usuarios inactivos son rechazados.

### RF-002 — Logout de usuario

**Estado:** PLANIFICADO

**Descripcion:** El sistema debe cerrar la sesion actual del usuario autenticado.

**Reglas:**

- Logout debe revocar la sesion persistida.
- Un token con sesion revocada no debe ser aceptado.

**Criterios de aceptacion:**

- Logout marca la sesion como revocada.
- El token previo queda inutilizable para endpoints protegidos.

### RF-003 — Cambio de contrasena

**Estado:** PLANIFICADO

**Descripcion:** Cualquier perfil autenticado debe poder cambiar su contrasena.

**Reglas:**

- La nueva contrasena debe almacenarse hasheada con bcrypt.
- No se debe exponer `password_hash` por API.

**Criterios de aceptacion:**

- La contrasena cambia correctamente.
- El hash no aparece en respuestas HTTP.

### RF-004 — Restriccion por contrasena temporal

**Estado:** PLANIFICADO

**Descripcion:** Usuarios creados con contrasena temporal deben cambiarla en el primer ingreso.

**Reglas:**

- Un usuario con `must_change_password = true` solo puede cambiar contrasena y cerrar sesion.
- Al cambiar la contrasena, la marca debe quedar resuelta.

**Criterios de aceptacion:**

- El acceso a otras funcionalidades queda bloqueado mientras la marca esta activa.
- Cambio exitoso habilita el flujo normal segun perfil.

### RF-005 — Gestion de usuarios por administrador

**Estado:** PLANIFICADO

**Descripcion:** `ADMIN` debe poder crear, listar, ver detalle, actualizar y desactivar usuarios.

**Reglas:**

- Solo `ADMIN` puede gestionar usuarios.
- La eliminacion debe preferir desactivacion logica.

**Criterios de aceptacion:**

- Perfiles no administradores no acceden a gestion de usuarios.
- La desactivacion deja `active = false`.

### RF-006 — Restriccion de creacion de administradores

**Estado:** PLANIFICADO

**Descripcion:** Un `ADMIN` no debe poder crear otros usuarios `ADMIN`.

**Reglas:**

- `POST /users` solo permite crear `EXECUTOR` o `AUDITOR`.
- Los usuarios creados por `ADMIN` nacen con contrasena temporal.

**Criterios de aceptacion:**

- Intentar crear `ADMIN` es rechazado.
- Usuarios creados quedan con `must_change_password = true`.

### RF-007 — Gestion de tareas por administrador

**Estado:** PLANIFICADO

**Descripcion:** `ADMIN` debe poder crear, listar, ver detalle, actualizar y eliminar tareas segun reglas.

**Reglas:**

- Una tarea debe tener titulo, descripcion y fecha de vencimiento.
- Una tarea nueva nace en `ASSIGNED`.

**Criterios de aceptacion:**

- `ADMIN` crea tareas validas.
- Tareas nuevas quedan en `ASSIGNED`.

### RF-008 — Asignacion de tareas solo a ejecutores

**Estado:** PLANIFICADO

**Descripcion:** Las tareas solo pueden asignarse a usuarios `EXECUTOR`.

**Reglas:**

- No se puede asignar una tarea a `ADMIN`.
- No se puede asignar una tarea a `AUDITOR`.

**Criterios de aceptacion:**

- Asignaciones a `EXECUTOR` son aceptadas.
- Asignaciones a otros perfiles son rechazadas.

### RF-009 — Restriccion de actualizacion/eliminacion por estado

**Estado:** PLANIFICADO

**Descripcion:** `ADMIN` solo puede actualizar o eliminar tareas en estado `ASSIGNED`.

**Reglas:**

- Tareas en estados distintos de `ASSIGNED` no pueden ser actualizadas por `ADMIN`.
- Tareas en estados distintos de `ASSIGNED` no pueden ser eliminadas por `ADMIN`.

**Criterios de aceptacion:**

- Operaciones sobre `ASSIGNED` son aceptadas si cumplen permisos.
- Operaciones sobre otros estados son rechazadas.

### RF-010 — Listado de tareas propias para ejecutor

**Estado:** PLANIFICADO

**Descripcion:** `EXECUTOR` debe poder listar solo sus tareas asignadas.

**Reglas:**

- No puede ver tareas de otros ejecutores.

**Criterios de aceptacion:**

- La lista contiene solo tareas con `assignee_id` del usuario autenticado.

### RF-011 — Detalle de tarea propia para ejecutor

**Estado:** PLANIFICADO

**Descripcion:** `EXECUTOR` debe poder ver detalle solo de tareas propias.

**Reglas:**

- Acceso a tareas ajenas debe ser rechazado.

**Criterios de aceptacion:**

- Tarea propia retorna detalle.
- Tarea ajena no retorna informacion sensible.

### RF-012 — Actualizacion de estado por ejecutor

**Estado:** PLANIFICADO

**Descripcion:** `EXECUTOR` debe poder modificar el estado de tareas propias segun flujo permitido.

**Reglas:**

- No puede modificar datos generales de la tarea.
- Solo puede aplicar transiciones validas.

**Criterios de aceptacion:**

- Transiciones validas son aceptadas.
- Transiciones invalidas son rechazadas.

### RF-013 — Bloqueo de actualizacion de tarea vencida

**Estado:** PLANIFICADO

**Descripcion:** `EXECUTOR` no puede cambiar estado de una tarea vencida.

**Reglas:**

- Si `due_at` esta vencido, el cambio de estado debe rechazarse.

**Criterios de aceptacion:**

- Tarea vencida no cambia de estado por ejecutor.

### RF-014 — Comentario sobre tarea vencida

**Estado:** PLANIFICADO

**Descripcion:** `EXECUTOR` debe poder comentar tareas vencidas propias.

**Reglas:**

- Solo ejecutores pueden comentar tareas propias vencidas.
- El comentario queda asociado a tarea y usuario.

**Criterios de aceptacion:**

- Comentario valido queda persistido.
- Comentario sobre tarea ajena o no vencida es rechazado.

### RF-015 — Visualizacion de tareas por auditor

**Estado:** PLANIFICADO

**Descripcion:** `AUDITOR` debe poder listar y ver tareas de cualquier usuario.

**Reglas:**

- `AUDITOR` no puede crear, modificar ni eliminar usuarios o tareas.

**Criterios de aceptacion:**

- `AUDITOR` accede a consultas de tareas.
- Operaciones de escritura son rechazadas.

### RF-016 — Control de transiciones de estado

**Estado:** PLANIFICADO

**Descripcion:** El sistema debe controlar el flujo de estados de tareas.

**Reglas:**

- `ASSIGNED -> STARTED`.
- `STARTED -> WAITING`, `FINISHED_SUCCESS` o `FINISHED_ERROR`.
- `WAITING -> WAITING`, `FINISHED_SUCCESS` o `FINISHED_ERROR`.
- Estados terminales: `FINISHED_SUCCESS` y `FINISHED_ERROR`.

**Criterios de aceptacion:**

- Las transiciones permitidas son aceptadas.
- Las transiciones no listadas son rechazadas.
