# 0005 — Usar GORM como ORM

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

La persistencia debe implementarse en Go sobre PostgreSQL sin exponer detalles de base de datos al dominio.

## Decisión

Usar GORM como ORM encapsulado solo en adapters outbound de persistencia.

## Consecuencias

Reduce codigo repetitivo de persistencia. Puede acoplar si se filtra fuera del adapter, por lo que dominio y aplicación no deben importar GORM.

## Alternativas consideradas

SQL manual y sqlc. Se descartan en esta etapa para priorizar velocidad de implementación y simplicidad.

## Relación con otros ADRs

Depende de ADR 0004 y se complementa con ADR 0006.
