# Validacion

## Estado

`PLANIFICADO`

Las validaciones ejecutables completas dependen de la implementacion futura. En la etapa actual aplica validacion documental.

## Reglas

- No inventar comandos de validacion.
- Documentar que se valido.
- Documentar que no se pudo validar.
- Registrar bloqueos reales.
- Actualizar esta guia cuando existan comandos reales.

## Validacion documental actual

- Verificar existencia de documentos obligatorios.
- Verificar ADRs iniciales.
- Verificar que no exista codigo de aplicacion.
- Verificar que no exista carpeta `migrations/`.
- Verificar que la documentacion no prometa funcionalidades implementadas cuando estan planificadas.

## Validacion futura

- `go test ./...`: `PENDIENTE DE IMPLEMENTACION`.
- Build de aplicacion: `PENDIENTE DE IMPLEMENTACION`.
- Smoke tests HTTP: `PENDIENTE DE IMPLEMENTACION`.
