# Guardrails de dominio

## Proposito

Este documento contiene reglas que futuros agentes no deben romper.

## Estado

`PLANIFICADO`

## Usuarios

- Solo `ADMIN` puede crear usuarios.
- `ADMIN` solo puede crear usuarios `EXECUTOR` o `AUDITOR`.
- `ADMIN` no puede crear otros `ADMIN`.
- Todo usuario creado por `ADMIN` nace con contrasena temporal o con `must_change_password = true`.
- Un usuario inactivo no puede iniciar sesion.
- Nunca retornar `password_hash` por API.
- `email` debe ser unico.
- La desactivacion logica se prefiere sobre eliminacion fisica.

## Autenticacion

- Login valido debe crear o asociar una sesion.
- JWT debe incluir `session_id`.
- Logout debe revocar la sesion.
- Un token con sesion revocada no debe ser aceptado.
- Usuarios con `must_change_password = true` solo pueden cambiar contrasena y cerrar sesion.
- Las contrasenas deben almacenarse con bcrypt.

## Tareas

- Una tarea nueva nace en `ASSIGNED`.
- Solo se puede asignar una tarea a un usuario `EXECUTOR`.
- `ADMIN` solo puede actualizar o eliminar tareas en `ASSIGNED`.
- `EXECUTOR` solo puede listar y ver tareas propias.
- `EXECUTOR` solo puede cambiar estado de tareas propias.
- `EXECUTOR` no puede cambiar estado de una tarea vencida.
- `EXECUTOR` solo puede comentar tareas vencidas propias.
- `AUDITOR` solo puede leer tareas.
- Las transiciones de estado deben respetar el flujo definido.

## Estados de tarea

- `ASSIGNED`.
- `STARTED`.
- `WAITING`.
- `FINISHED_SUCCESS`.
- `FINISHED_ERROR`.

## Transiciones permitidas

- `ASSIGNED -> STARTED`.
- `STARTED -> WAITING`.
- `STARTED -> FINISHED_SUCCESS`.
- `STARTED -> FINISHED_ERROR`.
- `WAITING -> WAITING`.
- `WAITING -> FINISHED_SUCCESS`.
- `WAITING -> FINISHED_ERROR`.

## Arquitectura

- Dominio sin dependencias externas.
- Dominio no importa GORM.
- Application no importa GORM.
- Handlers no contienen reglas de negocio.
- Repositories no contienen reglas de negocio.
- GORM vive solo en adapters outbound.
- DTOs HTTP no deben ser usados como entidades de dominio.
- Modelos de dominio y modelos GORM deben mantenerse separados.

## Persistencia

- Usar PostgreSQL.
- Usar GORM.
- Usar `AutoMigrate`.
- No usar migraciones SQL versionadas en esta etapa.
- No crear carpeta `migrations/`.
