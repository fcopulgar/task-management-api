# 0007 — Usar JWT con sesiones revocables

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

La API necesita tokens para autenticación y logout efectivo. Un JWT puro no permite invalidacion inmediata sin estado adicional.

## Decisión

Usar JWT con `session_id` en claims y sesiones persistidas revocables en base de datos.

## Consecuencias

Permite validar tokens y cerrar sesión revocando la sesión persistida. Requiere consultar estado de sesión en endpoints protegidos.

## Alternativas consideradas

JWT sin estado y sesiones opacas. JWT sin estado se descarta porque logout no revocaria tokens ya emitidos.

## Relación con otros ADRs

Se relaciona con seguridad, middleware de autenticación y persistencia PostgreSQL.
