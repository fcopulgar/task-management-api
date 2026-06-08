# Checklist de produccion

## Estado

`PENDIENTE DE IMPLEMENTACION`

No existe implementacion desplegable todavia.

## Seguridad

- [ ] Configurar gestion de secretos.
- [ ] Validar expiracion y revocacion de sesiones.
- [ ] Confirmar que `password_hash` no se expone por API.
- [ ] Revisar politica de CORS si aplica.

## Base de datos

- [ ] Definir estrategia de respaldo y recuperacion.
- [ ] Definir monitoreo de PostgreSQL.
- [ ] Revisar uso de `AutoMigrate` antes de produccion.

## Operacion

- [ ] Definir entorno de despliegue.
- [ ] Definir health checks.
- [ ] Definir logging operacional.
- [ ] Definir monitoreo y alertas.

## Pendiente de definicion

- Responsables operativos.
- Politicas finales de secretos.
- Estrategia de despliegue.
