# Metadata — Construccion del proyecto

Este directorio contiene los prompts originales usados para construir el proyecto de forma incremental con agentes de IA.

## Orden de ejecucion

El proyecto se construyo en tres etapas principales:

### 1. Estructura documental base

**Prompt:** [`01-initital-prompt-structure.md`](01-initital-prompt-structure.md)

Crea el harness documental del repositorio: `README.md`, `AGENTS.md`, `PLAN.md`, y toda la estructura bajo `docs/` (requirements, architecture, decisions, plans, standards). Establece las reglas de trabajo para agentes, convenciones de documentacion en espanol, plantillas de ADRs y el sistema de fases. No define tecnologias ni implementa codigo.

### 2. Especificacion del producto

**Prompt:** [`02-requeriments-prompt.md`](02-requeriments-prompt.md)

Define el producto `task-management-api`, el stack tecnologico (Go, chi, GORM, PostgreSQL, JWT, bcrypt), los requisitos funcionales (RF-001 a RF-016) y no funcionales (RNF-001 a RNF-010), crea los 8 ADRs iniciales, documenta la arquitectura hexagonal objetivo, los endpoints planificados, y genera el `PLAN.md` con 9 fases de implementacion. Sigue sin implementar codigo.

### 3. Implementacion por fases

**Prompt:** definido en [`docs/AGENT_WORKFLOW.md`](../docs/AGENT_WORKFLOW.md)

Cada fase se ejecuto individualmente usando el siguiente formato:

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

**Fases ejecutadas:**

| Fase | Nombre | Resultado |
|------|--------|-----------|
| 1 | Bootstrap tecnico del proyecto | `go.mod`, `chi`, servidor minimo, Docker, Makefile |
| 2 | Dominio y reglas centrales | Entidades, roles, estados, transiciones, errores, 36 tests |
| 3 | Puertos de aplicacion y contratos internos | Interfaces de repositorios, DTOs, mocks, 10 tests |
| 4 | Persistencia GORM y AutoMigrate | Modelos GORM, mappers, repositorios, PostgreSQL, 13 tests |
| 5 | Autenticacion, sesiones y cambio de contrasena | Login, JWT, bcrypt, middleware, 33 tests |
| 6 | Gestion de usuarios administrador | CRUD usuarios, autorizacion, desactivacion logica, 18 tests |
| 7 | Gestion de tareas administrador y auditor | CRUD tareas, asignacion EXECUTOR, restriccion ASSIGNED, 14 tests |
| 8 | Flujos de ejecutor | Tareas propias, transiciones, comentarios en vencidas, 16 tests |
| 9 | Tests HTTP, hardening y documentacion final | Smoke tests, API docs, checklist, deuda tecnica |

Cada fase se ejecuto de forma aislada, respetando estrictamente su alcance sin adelantar trabajo de fases futuras.
