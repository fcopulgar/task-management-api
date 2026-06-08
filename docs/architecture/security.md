# Seguridad

## Estado

`PLANIFICADO`

## Autenticación

- Login con credenciales.
- JWT como token de acceso.
- Claims deben incluir usuario, perfil y `session_id`.
- Usuarios inactivos no pueden iniciar sesión.

## Sesiones revocables

- Login valido debe crear o asociar una sesión persistida.
- Logout debe marcar `revoked_at`.
- Middleware de autenticación debe rechazar sesiones revocadas.
- `expires_at` define vigencia persistida de sesión.

## Autorización

- `ADMIN`: administra usuarios y tareas bajo reglas de estado.
- `EXECUTOR`: opera solo sobre tareas propias.
- `AUDITOR`: solo lee tareas.

## Contraseñas

- Usar bcrypt para hashing.
- Nunca retornar `password_hash` por API.
- Usuarios con `must_change_password = true` solo pueden cambiar contraseña y cerrar sesión.

## Datos sensibles

- No registrar contraseñas en logs.
- No exponer hashes en DTOs de salida.
- La politica final de secretos esta `PENDIENTE DE DEFINICION`.
