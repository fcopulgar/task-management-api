# 0007 — Usar JWT con sesiones revocables

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

La API necesita tokens para autenticacion y logout efectivo. Un JWT puro no permite invalidacion inmediata sin estado adicional.

## Decision

Usar JWT con `session_id` en claims y sesiones persistidas revocables en base de datos.

## Consecuencias

Permite validar tokens y cerrar sesion revocando la sesion persistida. Requiere consultar estado de sesion en endpoints protegidos.

## Alternativas consideradas

JWT sin estado y sesiones opacas. JWT sin estado se descarta porque logout no revocaria tokens ya emitidos.

## Relacion con otros ADRs

Se relaciona con seguridad, middleware de autenticacion y persistencia PostgreSQL.
