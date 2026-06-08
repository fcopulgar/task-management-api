# Decisiones arquitectonicas

## Estado

`IMPLEMENTADO`

## Que es un ADR

Un ADR registra una decision arquitectonica relevante, su contexto, consecuencias y alternativas consideradas.

## ADRs vigentes

- `0001-use-go.md`: usar Go como lenguaje principal.
- `0002-use-hexagonal-architecture.md`: usar arquitectura hexagonal minimalista.
- `0003-use-chi-router.md`: usar `chi` como router HTTP.
- `0004-use-postgresql.md`: usar PostgreSQL.
- `0005-use-gorm.md`: usar GORM.
- `0006-use-gorm-automigrate.md`: usar GORM `AutoMigrate`.
- `0007-use-jwt-with-revocable-sessions.md`: usar JWT con sesiones revocables.
- `0008-use-bcrypt-for-password-hashing.md`: usar bcrypt para hash de contrasenas.

## Numeracion

Usar numeracion secuencial de cuatro digitos:

```text
0001-titulo-de-la-decision.md
0002-otra-decision.md
```

## Cambio de decisiones

Los ADRs son historial tecnico. No se borran cuando una decision cambia. Si una decision deja de estar vigente, se crea un nuevo ADR y el anterior se marca como `Reemplazado` o `Deprecado`.
