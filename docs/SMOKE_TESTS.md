# Smoke tests

## Estado

`PENDIENTE DE IMPLEMENTACION`

No existen smoke tests ejecutables porque todavia no hay aplicacion.

## Validacion documental disponible

- Confirmar que existe `README.md`.
- Confirmar que existe `PLAN.md`.
- Confirmar que existe `AGENTS.md`.
- Confirmar que existe `docs/CURRENT_STATE.md`.
- Confirmar que existen ADRs iniciales en `docs/decisions/`.
- Confirmar que existe `docs/backend/`.
- Confirmar que no existe codigo de aplicacion.
- Confirmar que no existe carpeta `migrations/`.

## Smoke tests futuros

- Login exitoso.
- Logout revoca sesion.
- Token con sesion revocada es rechazado.
- Usuario con contrasena temporal solo puede cambiar contrasena y cerrar sesion.
- `ADMIN` no puede crear otro `ADMIN`.
- `EXECUTOR` no puede ver tareas ajenas.
- `AUDITOR` no puede escribir tareas.
