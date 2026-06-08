# Testing backend

## Estado

`PLANIFICADO`

## Herramientas definidas

- `testing`.
- `testify`.
- `httptest` cuando corresponda.

## Cobertura esperada

- Tests unitarios de dominio para roles, estados y transiciones.
- Tests de casos de uso para permisos y restricciones.
- Tests de autenticación para JWT, `session_id`, sesiones revocadas y contraseñas temporales.
- Tests HTTP de flujos criticos cuando existan handlers.
- Tests de persistencia cuando exista infraestructura local.

## Pendiente de implementación

- Comandos reales de test.
- Fixtures.
- Helpers de test.
- Smoke tests ejecutables.
