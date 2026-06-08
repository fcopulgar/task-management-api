# Validación

## Estado

`PLANIFICADO`

Las validaciones ejecutables completas dependen de la implementación futura. En la etapa actual aplica validación documental.

## Reglas

- No inventar comandos de validación.
- Documentar que se valido.
- Documentar que no se pudo validar.
- Registrar bloqueos reales.
- Actualizar esta guia cuando existan comandos reales.

## Validación documental actual

- Verificar existencia de documentos obligatorios.
- Verificar ADRs iniciales.
- Verificar que no exista codigo de aplicación.
- Verificar que no exista carpeta `migrations/`.
- Verificar que la documentación no prometa funcionalidades implementadas cuando estan planificadas.

## Validación futura

- `go test ./...`: `PENDIENTE DE IMPLEMENTACION`.
- Build de aplicación: `PENDIENTE DE IMPLEMENTACION`.
- Smoke tests HTTP: `PENDIENTE DE IMPLEMENTACION`.
