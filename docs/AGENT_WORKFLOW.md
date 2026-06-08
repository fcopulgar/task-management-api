# Flujo operativo para agentes

## Estado

`IMPLEMENTADO`

## Inicio obligatorio

Antes de trabajar, leer las fuentes de verdad en el orden definido por `AGENTS.md`.

## Formato para iniciar una fase futura

```text
Quiero trabajar solo en:

Fase X — <nombre exacto de PLAN.md>

Reglas:
- lee primero README.md, PLAN.md, AGENTS.md y docs/CURRENT_STATE.md
- trabaja solo dentro del alcance de la fase
- no adelantes fases futuras
- no hagas refactors fuera de scope
- no inventes requisitos
- no inventes tecnologias
- si cambia una decision arquitectonica, crea un nuevo ADR
- actualiza PLAN.md
- actualiza CURRENT_STATE.md
- actualiza TECH_DEBT.md si corresponde
- reporta validacion realizada
```

## Cierre obligatorio

Todo cierre debe reportar:

- archivos modificados
- tareas del plan completadas
- validacion realizada
- documentacion actualizada
- deuda o pendientes detectados
