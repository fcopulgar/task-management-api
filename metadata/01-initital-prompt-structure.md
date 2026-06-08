# Bootstrap universal de harness para proyecto nuevo

Quiero inicializar un repositorio vacío con un harness profesional para desarrollar un proyecto de software de forma incremental usando agentes de IA.

Este prompt NO debe implementar el producto.

Este prompt NO debe asumir tecnologías.

Este prompt NO debe crear código fuente.

Este prompt solo debe crear la estructura documental, reglas de trabajo, convenciones, plantillas y esqueleto de gobernanza técnica del proyecto.

---

# 1. Datos mínimos del proyecto

## Nombre del proyecto

```text
<PROJECT_NAME>
```

---

# 2. Objetivo de esta tarea

Crear una base de documentación y reglas para que el proyecto pueda ser desarrollado por fases posteriormente.

El resultado debe permitir que cualquier agente futuro entienda:

* cómo trabajar en el repositorio
* dónde documentar decisiones
* cómo usar planes activos
* cómo archivar planes antiguos
* cómo registrar deuda técnica
* cómo documentar estado actual
* cómo crear ADRs versionados
* cómo evitar inventar arquitectura, tecnologías o requisitos
* cómo continuar el proyecto sin depender del contexto de un chat

---

# 3. Reglas estrictas

No implementar código.

No crear archivos propios de una tecnología específica.

No crear:

```text
src/
app/
cmd/
internal/
pkg/
backend/
frontend/
mobile/
workers/
package.json
go.mod
pyproject.toml
requirements.txt
Cargo.toml
pom.xml
build.gradle
docker-compose.yml
Dockerfile
```

No crear documentación específica de una tecnología todavía.

No inventar requisitos.

No inventar arquitectura.

No crear un plan de implementación concreto todavía.

No crear tareas técnicas específicas de una tecnología.

No agregar dependencias.

No agregar comandos de build, test o run todavía.

No crear documentación en inglés.

Toda la documentación debe estar en español.

Excepciones permitidas:

* nombres de archivos técnicos estándar
* nombres de carpetas técnicas estándar
* términos ampliamente usados como ADR, README, API, CI/CD, CLI

Si algo todavía no está definido, marcarlo como:

```text
PENDIENTE DE DEFINICIÓN
```

Si algo todavía no está implementado, marcarlo como:

```text
PENDIENTE DE IMPLEMENTACIÓN
```

Si algo no puede inferirse, marcarlo como:

```text
NO CONFIRMADO
```

---

# 4. Estructura que debes crear

Crear esta estructura base:

```text
README.md
AGENTS.md
PLAN.md

docs/
  README.md
  CURRENT_STATE.md
  DOMAIN_GUARDRAILS.md
  AGENT_WORKFLOW.md
  TECH_DEBT.md
  PRODUCTION_CHECKLIST.md
  SMOKE_TESTS.md

  requirements/
    README.md
    product-scope.md
    functional-requirements.md
    non-functional-requirements.md
    open-questions.md

  architecture/
    README.md
    overview.md
    module-boundaries.md
    data-flow.md
    security.md

  decisions/
    README.md
    TEMPLATE.md

  plans/
    README.md
    archive/
      .gitkeep
    future/
      .gitkeep

  standards/
    README.md
    documentation-style.md
    definition-of-done.md
    validation.md
```

No crear carpetas adicionales.

No crear estructura de código.

---

# 5. `README.md`

Crear un `README.md` breve y profesional.

Debe incluir:

* nombre del proyecto
* estado inicial
* propósito del repositorio
* estructura documental
* cómo deben trabajar futuros agentes
* dónde están las fuentes de verdad
* qué falta definir

No debe incluir stack técnico.

No debe incluir comandos de instalación.

No debe incluir instrucciones de ejecución.

No debe prometer funcionalidades.

Debe dejar claro que el proyecto aún está en fase de definición inicial.

---

# 6. `AGENTS.md`

Crear `AGENTS.md` como el archivo principal para cualquier agente.

Debe incluir:

## Propósito

Explicar que este archivo define reglas obligatorias de trabajo para agentes.

## Fuentes de verdad

Orden inicial:

1. `README.md`
2. `PLAN.md`
3. `AGENTS.md`
4. `docs/CURRENT_STATE.md`
5. `docs/requirements/`
6. `docs/DOMAIN_GUARDRAILS.md`
7. `docs/architecture/`
8. `docs/decisions/`
9. `docs/TECH_DEBT.md`

## Reglas de trabajo

Incluir reglas como:

* trabajar solo dentro del alcance pedido
* no adelantar fases futuras
* no inventar requisitos
* no inventar tecnologías
* no inventar arquitectura
* no implementar código si la tarea es documental
* no modificar decisiones históricas sin crear nuevo ADR
* no borrar ADRs
* no leer planes archivados salvo indicación explícita
* registrar deuda técnica real
* mantener documentación en español

## Resumen obligatorio después de cada tarea

Incluir exactamente esta sección:

```md
## Resumen obligatorio después de cada tarea

Siempre reportar:

- archivos modificados
- tareas del plan completadas
- validación realizada
- documentación actualizada
- deuda o pendientes detectados
```

---

# 7. `PLAN.md`

Crear un `PLAN.md` inicial, pero no debe contener un plan técnico real todavía.

Debe indicar que el proyecto está en fase de definición.

Debe incluir:

```md
# Plan activo

## Estado

`PENDIENTE DE ESPECIFICACIÓN`

## Objetivo actual

Completar la especificación del proyecto, tecnologías, decisiones técnicas y requisitos antes de iniciar implementación.

## Fase activa

### Fase 0 — Definición del proyecto

**Estado:** `PENDIENTE ⬜`

#### Objetivo

Completar requisitos, tecnologías, decisiones técnicas y arquitectura inicial.

#### Alcance exacto

- Documentar alcance funcional
- Documentar alcance no funcional
- Documentar tecnologías elegidas
- Crear ADRs iniciales
- Crear plan de implementación por fases
- Actualizar estado actual del proyecto

#### Tareas

- [ ] Completar requisitos funcionales
- [ ] Completar requisitos no funcionales
- [ ] Completar decisiones técnicas iniciales
- [ ] Crear ADRs iniciales
- [ ] Crear plan de implementación real
- [ ] Actualizar documentación vigente

#### Criterios de aceptación

- Los requisitos principales están documentados
- Las tecnologías principales están documentadas
- Las decisiones arquitectónicas iniciales tienen ADR
- Existe un plan de implementación por fases
- No se ha implementado código todavía

#### Entregables

- `docs/requirements/`
- `docs/architecture/`
- `docs/decisions/`
- `PLAN.md` actualizado
```

No agregar fases técnicas concretas todavía.

---

# 8. `docs/CURRENT_STATE.md`

Crear el archivo con esta estructura:

```md
# Estado actual

## Resumen

El proyecto fue inicializado con un harness documental y operativo. Aún no tiene requisitos, tecnologías ni implementación definidos.

## Implementado

- Harness documental base.
- Reglas de trabajo para agentes.
- Estructura inicial de documentación.
- Convención para planes.
- Convención para ADRs versionados.

## Parcialmente implementado

- PENDIENTE DE DEFINICIÓN

## No implementado todavía

- Código de aplicación.
- Stack tecnológico.
- Arquitectura concreta.
- Requisitos funcionales.
- Requisitos no funcionales.
- Plan técnico de implementación.

## Decisiones confirmadas

- La documentación se mantiene en español.
- Se usan ADRs versionados para decisiones arquitectónicas relevantes.
- El desarrollo se organiza mediante `PLAN.md`.
- Los planes archivados no se leen salvo necesidad explícita.

## Pendiente de definición

- Nombre funcional del producto.
- Requisitos funcionales.
- Requisitos no funcionales.
- Tecnologías.
- Arquitectura concreta.
- Estrategia de testing.
- Estrategia de ejecución local.
- Estrategia de despliegue.

## Riesgos técnicos

- No hay riesgos técnicos evaluables todavía porque el stack no está definido.

## Validación disponible

- Validación documental de estructura creada.

## Qué no se pudo validar

- Build.
- Tests.
- Ejecución local.
- Integraciones.
```

Adaptar solo el nombre del proyecto.

---

# 9. `docs/DOMAIN_GUARDRAILS.md`

Crear un archivo de guardrails, pero sin inventar dominio.

Debe incluir:

```md
# Guardrails de dominio

## Propósito

Este documento contiene reglas de dominio que los agentes no deben romper.

## Estado actual

`PENDIENTE DE DEFINICIÓN`

El dominio funcional todavía no fue especificado.

## Reglas conocidas

- No inventar entidades de dominio.
- No inventar reglas de negocio.
- No inventar permisos.
- No inventar estados.
- No inventar integraciones.
- Si falta una regla, marcarla como `PENDIENTE DE DEFINICIÓN`.

## Pendiente de definición

- Entidades principales.
- Actores.
- Permisos.
- Estados.
- Flujos críticos.
- Reglas de datos sensibles.
- Reglas de borrado.
- Reglas de auditoría.
```

---

# 10. `docs/requirements/`

Crear documentación base para requisitos.

## `docs/requirements/README.md`

Explicar que esta carpeta contiene alcance y requisitos.

## `product-scope.md`

Debe contener:

* propósito del producto
* usuarios objetivo
* problema que resuelve
* alcance inicial
* fuera de alcance
* pendientes

Todo como `PENDIENTE DE DEFINICIÓN`.

## `functional-requirements.md`

Debe contener una plantilla clara para requisitos funcionales.

## `non-functional-requirements.md`

Debe contener una plantilla para requisitos no funcionales.

## `open-questions.md`

Debe contener preguntas abiertas pendientes.

---

# 11. `docs/architecture/`

Crear documentación base sin inventar arquitectura.

Cada archivo debe tener contenido útil, pero marcar como pendiente lo no definido.

## `overview.md`

Debe indicar que la arquitectura concreta está pendiente.

## `module-boundaries.md`

Debe explicar que los límites de módulos se definirán cuando existan tecnologías y requisitos.

## `data-flow.md`

Debe indicar que los flujos de datos están pendientes.

## `security.md`

Debe indicar que la estrategia de seguridad está pendiente.

No inventar diagramas técnicos.

Puedes incluir placeholders Mermaid vacíos solo si aportan, pero deben estar marcados como `PENDIENTE DE DEFINICIÓN`.

---

# 12. `docs/decisions/`

Crear estructura de ADRs versionados.

## `docs/decisions/README.md`

Debe explicar:

* qué es un ADR
* cuándo crear un ADR
* cómo numerarlo
* cómo cambiar una decisión
* por qué no se deben borrar ADRs antiguos

Debe incluir esta regla:

```md
Los ADRs son historial técnico. No se borran cuando una decisión cambia. Si una decisión deja de estar vigente, se crea un nuevo ADR y el anterior se marca como `Reemplazado` o `Deprecado`.
```

## `docs/decisions/TEMPLATE.md`

Crear plantilla:

```md
# NNNN — Título de la decisión

**Estado:** Propuesto  
**Fecha:** YYYY-MM-DD

## Contexto

...

## Decisión

...

## Consecuencias

...

## Alternativas consideradas

...

## Relación con otros ADRs

...
```

No crear ADRs concretos todavía, porque no hay decisiones técnicas definidas.

---

# 13. `docs/plans/`

Crear documentación de planes.

## `docs/plans/README.md`

Debe explicar:

* `PLAN.md` es el plan activo
* `docs/plans/archive/` contiene planes cerrados
* `docs/plans/future/` contiene planes futuros no activos
* los agentes deben leer primero `PLAN.md`
* los planes archivados solo se leen cuando la fase activa lo pida explícitamente

---

# 14. `docs/TECH_DEBT.md`

Crear archivo de deuda técnica.

No inventar deuda.

Debe indicar que aún no existe deuda técnica evaluable porque el proyecto no tiene implementación.

Usar estructura:

```md
# Deuda técnica

## Crítica

Sin deuda registrada.

## Alta

Sin deuda registrada.

## Media

Sin deuda registrada.

## Baja

Sin deuda registrada.

## Deuda futura / no bloqueante

Sin deuda registrada.

## Resuelta

Sin deuda registrada.
```

---

# 15. `docs/AGENT_WORKFLOW.md`

Crear un prompt operativo genérico para fases futuras.

Debe incluir:

```text
Quiero trabajar solo en:

Fase X — <nombre exacto de PLAN.md>

Reglas:
- lee primero README.md, PLAN.md, AGENTS.md y docs/CURRENT_STATE.md
- trabaja solo dentro del alcance de la fase
- no adelantes fases futuras
- no hagas refactors fuera de scope
- no inventes requisitos
- no inventes tecnologías
- si cambia una decisión arquitectónica, crea o actualiza ADRs
- actualiza PLAN.md
- actualiza CURRENT_STATE.md
- actualiza TECH_DEBT.md si corresponde
- reporta validación realizada
```

---

# 16. `docs/standards/`

Crear estándares básicos:

## `documentation-style.md`

Debe indicar:

* documentación en español
* textos breves y accionables
* diferenciar implementado, pendiente y no confirmado
* no vender humo
* no prometer funcionalidades inexistentes
* no mezclar planes futuros con estado actual

## `definition-of-done.md`

Debe definir criterios generales para cerrar una fase:

* alcance completado
* validación ejecutada
* documentación actualizada
* deuda registrada
* pendientes explícitos
* resumen final entregado

## `validation.md`

Debe explicar que las validaciones concretas dependen del stack futuro.

No inventar comandos.

---

# 17. Validación de esta tarea

Al terminar:

* listar estructura creada
* verificar que no se creó código fuente
* verificar que no se crearon archivos de tecnologías específicas
* verificar que la documentación está en español
* verificar que `PLAN.md` no contiene implementación técnica concreta
* verificar que no se crearon ADRs concretos sin decisiones técnicas

No ejecutar build ni tests porque no hay tecnología definida.

---

# 18. Entrega final

Responder con:

1. resumen de la base creada
2. estructura creada
3. archivos principales generados
4. confirmación de que no se implementó código
5. confirmación de que no se asumieron tecnologías
6. validación realizada
7. siguiente paso recomendado

No implementar producto.

No crear código.

Haz los cambios directamente en el repositorio.
