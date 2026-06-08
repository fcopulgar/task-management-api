# Alcance del producto

## Proposito del producto

`task-management-api` es una API REST para gestionar usuarios y tareas con autenticacion, autorizacion por perfiles, cambio obligatorio de contrasena temporal, cierre de sesion y control de estados de tareas.

## Problema que resuelve

Permite centralizar la administracion de tareas asignadas a ejecutores, controlar quien puede modificar informacion y entregar visibilidad de estado a auditores sin exponer capacidades de escritura.

## Usuarios y actores

- `ADMIN`: administra usuarios y tareas.
- `EXECUTOR`: trabaja sobre tareas propias asignadas.
- `AUDITOR`: consulta tareas y estados sin modificar informacion.

## Alcance inicial

- Login, logout y cambio de contrasena.
- Sesiones revocables con JWT que incluye `session_id`.
- Usuarios con roles `ADMIN`, `EXECUTOR` y `AUDITOR`.
- Creacion y gestion de usuarios por `ADMIN`.
- Tareas asignadas solo a usuarios `EXECUTOR`.
- Flujo de estados de tareas.
- Comentarios de ejecutores en tareas vencidas propias.
- Consulta de tareas por auditor.

## Fuera de alcance

- Frontend.
- Mobile.
- Workers.
- CLI.
- Migraciones SQL versionadas.
- Integraciones externas.
- Notificaciones.
- Reporteria avanzada.

## Supuestos

- La primera version es una API REST backend.
- PostgreSQL sera la base de datos principal.
- La inicializacion de esquema se realizara con GORM `AutoMigrate`.
- Las tareas se asignan a un unico ejecutor mediante `assignee_id`.

## Pendientes

- Definir politicas finales de despliegue.
- Definir filtros especificos de consulta para auditoria.
- Resolver preguntas abiertas en `open-questions.md`.
