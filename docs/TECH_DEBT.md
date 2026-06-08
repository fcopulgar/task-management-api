# Deuda técnica

## Crítica

Sin deuda registrada.

## Alta

Sin deuda registrada.

## Media

Sin deuda registrada.

## Baja

- **Duración de token/sesión hardcodeada**: `DefaultTokenDurationHours = 24` y `DefaultSessionDuration = 24 * time.Hour` estan definidos como constantes en `internal/application/auth_usecase.go`. Seria deseable moverlos a configuración de entorno para mayor flexibilidad operativa.
- **Salida JSON en PascalCase**: `UserOutput`, `TaskOutput` y `CommentOutput` usan los nombres de campo Go por defecto (PascalCase) en las respuestas JSON. Para consistencia con APIs REST convencionales, considerar migrar a snake_case en una version futura. Esto es un cambio breaking que requiere coordinacion.

## Deuda futura / no bloqueante

- Evaluar migraciones SQL versionadas en una etapa futura si `AutoMigrate` deja de ser suficiente para evolucion controlada del esquema. No aplica a la etapa actual.
- Definir estrategia de despliegue, secretos, respaldo y recuperacion antes de producción.
- Agregar filtros avanzados para consultas de auditoria (por estado, rango de fechas, asignatario).

## Resuelta

Sin deuda registrada.
