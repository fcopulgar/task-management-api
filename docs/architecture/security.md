# Seguridad

## Estado

`PLANIFICADO`

## Autenticacion

- Login con credenciales.
- JWT como token de acceso.
- Claims deben incluir usuario, perfil y `session_id`.
- Usuarios inactivos no pueden iniciar sesion.

## Sesiones revocables

- Login valido debe crear o asociar una sesion persistida.
- Logout debe marcar `revoked_at`.
- Middleware de autenticacion debe rechazar sesiones revocadas.
- `expires_at` define vigencia persistida de sesion.

## Autorizacion

- `ADMIN`: administra usuarios y tareas bajo reglas de estado.
- `EXECUTOR`: opera solo sobre tareas propias.
- `AUDITOR`: solo lee tareas.

## Contrasenas

- Usar bcrypt para hashing.
- Nunca retornar `password_hash` por API.
- Usuarios con `must_change_password = true` solo pueden cambiar contrasena y cerrar sesion.

## Datos sensibles

- No registrar contrasenas en logs.
- No exponer hashes en DTOs de salida.
- La politica final de secretos esta `PENDIENTE DE DEFINICION`.
