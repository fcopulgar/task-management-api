# 0006 — Usar GORM AutoMigrate

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

La etapa inicial necesita inicializar esquema sin mantener migraciones SQL versionadas.

## Decision

Usar GORM `AutoMigrate` para inicializacion de esquema en esta etapa. No crear carpeta `migrations/` ni herramientas `golang-migrate` o `goose`.

## Consecuencias

Simplifica el bootstrap. Puede ser insuficiente para evolucion controlada de produccion, por lo que una decision futura podria reemplazarla con migraciones versionadas.

## Alternativas consideradas

Migraciones SQL versionadas con golang-migrate o goose. Se descartan en esta etapa por restriccion explicita.

## Relacion con otros ADRs

Depende de ADR 0005. Si cambia, crear nuevo ADR y marcar este como reemplazado o deprecado.
