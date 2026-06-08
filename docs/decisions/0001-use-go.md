# 0001 — Usar Go como lenguaje principal

**Estado:** Aceptado  
**Fecha:** 2026-06-07

## Contexto

El proyecto requiere un backend claro, eficiente y mantenible para exponer una API REST.

## Decisión

Usar Go como lenguaje principal del backend.

## Consecuencias

Go entrega tipado estatico, binarios simples y buen soporte para APIs HTTP. El equipo debera mantener convenciones claras para evitar acoplamiento innecesario.

## Alternativas consideradas

Node.js, Java y Python. Se descartan en esta etapa para mantener una base simple y alineada con el stack definido.

## Relación con otros ADRs

Base para ADRs de router, arquitectura y testing.
