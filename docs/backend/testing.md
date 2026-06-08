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
- Tests de autenticacion para JWT, `session_id`, sesiones revocadas y contrasenas temporales.
- Tests HTTP de flujos criticos cuando existan handlers.
- Tests de persistencia cuando exista infraestructura local.

## Pendiente de implementacion

- Comandos reales de test.
- Fixtures.
- Helpers de test.
- Smoke tests ejecutables.
