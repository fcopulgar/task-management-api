En mi experiencia como ingeniero de software, me he dado cuenta de que la IA todavía no se está utilizando correctamente en muchos equipos: se le pide código sin contexto, se gasta de más, se corrige demasiado a mano y se termina produciendo menos de lo que realmente se podría. En este proyecto quise mostrar un enfoque distinto: aplicar criterios reales de ingeniería de software antes de implementar, creando un harness con documentación, reglas, ADRs, guardrails, fases y validaciones para que la IA no improvise, sino que ejecute sobre una estructura clara. Con esa base, la implementación completa usando agentes costó aproximadamente **USD $0.50** en API.

---

## Qué se intentó demostrar

La idea no fue simplemente usar IA para programar una API.

La idea fue demostrar que el mayor valor no está en pedirle al modelo “haz todo”, sino en preparar bien el contexto para que cualquier agente pueda trabajar con control.

En vez de partir directo con código, primero se construyó una base de trabajo:

* reglas para agentes;
* documentación de requisitos;
* arquitectura objetivo;
* ADRs versionados;
* guardrails de dominio;
* plan por fases;
* criterios de aceptación;
* deuda técnica visible;
* smoke tests;
* estado actual del proyecto.

Ese harness hace que el proyecto no dependa del historial de un chat ni de un proveedor de IA específico.

---

## Enfoque usado

El trabajo se separó en dos grandes partes.

### 1. Planificación y harness

Primero se diseñó una estructura para que el agente no trabajara a ciegas.

El objetivo fue que el repositorio tuviera memoria propia:

```text
requisitos
decisiones técnicas
arquitectura
reglas de dominio
plan activo
estado actual
deuda técnica
validaciones
```

Esto permite que un agente nuevo pueda entrar al proyecto, leer los archivos correctos y continuar sin inventar contexto.

### 2. Implementación por fases

Después se ejecutó la implementación usando agentes, pero no con un prompt gigante de “hazme todo el proyecto”.

La implementación se trabajó por fases desde `PLAN.md`, manteniendo alcance, validaciones y documentación actualizada.

El flujo buscado fue:

```text
planificar -> documentar -> ejecutar -> validar -> corregir
```

---

## Herramientas utilizadas

Para planificación, diseño del harness, revisión de arquitectura, prompts y análisis crítico se utilizó:

```text
ChatGPT 5.5
```

Para implementación se utilizó:

```text
OpenCode + DeepSeek V4 PRO
```

Costo aproximado de API para la implementación:

```text
USD $0.50
```

El costo bajo no fue el punto de partida; fue una consecuencia de tener buen contexto, buenas restricciones y fases claras.

---

## Prompts principales

El proceso partió separando dos responsabilidades que normalmente se mezclan mal.

1. Crear el harness del proyecto.
2. Especificar el producto y planificar la implementación.

Los prompts principales fueron:

* [Prompt 1 — Bootstrap universal de harness](./01-initial-prompt-structure.md)
* [Prompt 2 — Especificación del proyecto y plan por fases](./02-requirements-prompt.md)

El primer prompt no implementa producto ni asume tecnología. Solo crea la estructura documental y las reglas de trabajo.

El segundo prompt toma requisitos, tecnologías y decisiones técnicas, y deja el proyecto listo para ser implementado por fases.

---

## Qué aporta el harness

El harness evita que el agente trabaje como si estuviera en una conversación aislada.

Le define:

* qué leer primero;
* qué decisiones respetar;
* qué reglas de dominio no romper;
* qué fase ejecutar;
* qué está fuera de alcance;
* cuándo crear ADRs;
* cuándo registrar deuda técnica;
* cómo validar;
* cómo reportar resultados.

La IA sigue escribiendo código, pero lo hace dentro de un sistema de trabajo más parecido a un proyecto real.

---

## ADRs y trazabilidad

El proyecto usa ADRs versionados en:

```text
docs/decisions/
```

La razón es simple: las decisiones técnicas tienen historia.

Si mañana se cambia PostgreSQL por MongoDB, no se borra la decisión anterior. Se crea un nuevo ADR, se explica el cambio y se marca el anterior como reemplazado o deprecado.

Esto permite entender no solo qué se decidió, sino también por qué se decidió.

---

## Rol humano

La IA ayudó con gran parte de la ejecución, pero el criterio técnico no se delega.

El trabajo humano estuvo en:

* definir el enfoque;
* separar planificación de implementación;
* decidir la arquitectura;
* corregir el rumbo;
* detectar sobreimplementación;
* revisar inconsistencias;
* priorizar mejoras;
* controlar el alcance;
* validar si el resultado era mantenible.

La IA acelera, pero necesita dirección.