# Checklist de producción

## Estado

`IMPLEMENTADO` (parcial — items operativos requieren entorno de despliegue)

## Seguridad

- [x] Configurar gestión de secretos — `JWT_SECRET` via variable de entorno.
- [x] Validar expiración y revocación de sesiones — sesiones revocables con `revoked_at`, middleware verifica validez.
- [x] Confirmar que `password_hash` no se expone por API — `UserToOutput` excluye el campo.
- [ ] Revisar politica de CORS si aplica — no se requiere para API backend sin frontend.
- [x] Contraseñas hasheadas con bcrypt.
- [x] JWT incluye `session_id` en claims.
- [x] Usuarios con `must_change_password` restringidos a cambio de contraseña y logout.
- [x] Autorización por roles (`RequireRole`, `RequireAnyRole`).
- [x] `ADMIN` no puede crear otros `ADMIN`.

## Base de datos

- [x] PostgreSQL como base de datos principal.
- [x] GORM `AutoMigrate` para inicialización de esquema.
- [x] UUIDs generados por PostgreSQL (`gen_random_uuid()`).
- [x] Conexión configurable via variables de entorno (`DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSLMODE`).
- [x] Pool de conexiones configurado (25 max open, 5 max idle, 5 min lifetime).
- [ ] Definir estrategia de respaldo y recuperacion — `PENDIENTE DE DEFINICION`.
- [ ] Definir monitoreo de PostgreSQL — `PENDIENTE DE DEFINICION`.
- [x] Revisar uso de `AutoMigrate` antes de producción — adecuado para etapa actual; evaluar migraciones SQL versionadas en etapa futura.

## Operación

- [ ] Definir entorno de despliegue — `PENDIENTE DE DEFINICION`.
- [x] Definir health checks — `GET /health` retorna `{"status":"ok"}`.
- [x] Definir logging operacional — chi `Logger` middleware incluido.
- [ ] Definir monitoreo y alertas — `PENDIENTE DE DEFINICION`.
- [x] Docker Compose para desarrollo local.
- [x] `Dockerfile` multi-stage optimizado.

## Testing

- [x] Tests unitarios de dominio (36 tests).
- [x] Tests unitarios de aplicación (38 tests).
- [x] Tests de contratos y DTOs (9 tests).
- [x] Tests de persistencia con PostgreSQL (13 tests, requieren BD).
- [x] Tests de seguridad y JWT (8 tests).
- [x] Tests HTTP de handlers (36 tests).
- [x] Smoke tests documentados (19 escenarios).
- [x] `go build`, `go vet`, `go test` automatizados.

## Pendiente de definición

- Responsables operativos.
- Politicas finales de secretos (rotacion, almacenamiento).
- Estrategia de despliegue.
- Filtros avanzados para auditoria.
- Migraciones SQL versionadas (cuando `AutoMigrate` sea insuficiente).
