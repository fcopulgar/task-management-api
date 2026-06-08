# 0004 — Usar PostgreSQL como base de datos

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

El sistema necesita persistencia relacional para usuarios, sesiones, tareas y comentarios.

## Decisión

Usar PostgreSQL como base de datos principal.

## Consecuencias

Entrega integridad relacional y soporte robusto para evolucion futura. Requiere configuración local y operativa adecuada.

## Alternativas consideradas

SQLite y MySQL. Se descartan para mantener una base relacional robusta desde el inicio.

## Relación con otros ADRs

Base para ADR 0005 sobre GORM y ADR 0006 sobre AutoMigrate.
