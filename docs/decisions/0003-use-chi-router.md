# 0003 — Usar chi como router HTTP

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

La API REST necesita un router HTTP liviano, idiomatico y compatible con middleware.

## Decisión

Usar `chi` como router HTTP.

## Consecuencias

Permite rutas claras y middlewares composables. No debe contener reglas de negocio en handlers.

## Alternativas consideradas

Router estándar `net/http` puro y frameworks mas amplios. Se elige `chi` por equilibrio entre simplicidad y ergonomia.

## Relación con otros ADRs

Se integra con la arquitectura hexagonal como adapter inbound HTTP.
