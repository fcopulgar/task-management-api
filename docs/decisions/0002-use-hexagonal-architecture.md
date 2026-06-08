# 0002 — Usar arquitectura hexagonal minimalista

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

El sistema debe separar reglas de negocio de frameworks, HTTP y persistencia sin sobredimensionar la solucion.

## Decision

Usar arquitectura hexagonal minimalista con dominio, aplicacion, adapters inbound y adapters outbound.

## Consecuencias

El dominio y la aplicacion quedan testeables y desacoplados. Exige mantener limites de dependencias y evitar que handlers o repositorios concentren reglas.

## Alternativas consideradas

Arquitectura por capas tradicional y arquitectura limpia completa. Se prefiere una variante hexagonal minimalista para reducir ceremonia.

## Relacion con otros ADRs

Condiciona ADRs de GORM, chi y separacion dominio/infraestructura.
