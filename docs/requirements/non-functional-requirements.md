# Requisitos no funcionales

## Estado

`PLANIFICADO`

### RNF-001 — Arquitectura hexagonal minimalista

**Estado:** PLANIFICADO

**Descripción:** El backend debe organizarse con una arquitectura hexagonal simple, separando dominio, aplicación, adapters inbound y adapters outbound.

**Criterios de aceptacion:**

- Las dependencias apuntan hacia dominio y aplicación.
- La infraestructura queda encapsulada en adapters.

### RNF-002 — Dominio desacoplado de infraestructura

**Estado:** PLANIFICADO

**Descripción:** El dominio no debe depender de GORM, HTTP, base de datos ni frameworks.

**Criterios de aceptacion:**

- Paquetes de dominio no importan GORM ni `net/http`.
- Reglas criticas pueden probarse sin infraestructura.

### RNF-003 — Persistencia con PostgreSQL y GORM

**Estado:** PLANIFICADO

**Descripción:** La persistencia principal debe usar PostgreSQL mediante GORM.

**Criterios de aceptacion:**

- Repositorios concretos usan GORM solo en adapters outbound.
- Los modelos GORM estan separados de entidades de dominio.

### RNF-004 — Inicialización con AutoMigrate

**Estado:** PLANIFICADO

**Descripción:** La inicialización de esquema debe usar GORM `AutoMigrate` en esta etapa.

**Criterios de aceptacion:**

- No existe carpeta `migrations/`.
- No se agregan herramientas `golang-migrate` ni `goose`.
- No existen comandos `make migrate-up` ni `make migrate-down`.

### RNF-005 — Seguridad de contraseñas con bcrypt

**Estado:** PLANIFICADO

**Descripción:** Las contraseñas deben almacenarse exclusivamente como hashes bcrypt.

**Criterios de aceptacion:**

- Ninguna respuesta API expone `password_hash`.
- La verificacion de credenciales usa comparacion bcrypt.

### RNF-006 — Autenticación con JWT y sesiones revocables

**Estado:** PLANIFICADO

**Descripción:** La autenticación debe usar JWT con `session_id` y validar que la sesión persistida no este revocada.

**Criterios de aceptacion:**

- El token contiene `session_id`.
- Logout revoca la sesión.
- Sesiones revocadas no autorizan endpoints protegidos.

### RNF-007 — Documentación en español

**Estado:** PLANIFICADO

**Descripción:** Toda la documentación permanente del repositorio debe mantenerse en español.

**Criterios de aceptacion:**

- Nuevos documentos se escriben en español.
- Estados usan las marcas definidas por `AGENTS.md`.

### RNF-008 — ADRs versionados

**Estado:** PLANIFICADO

**Descripción:** Las decisiones tecnicas relevantes deben registrarse como ADRs secuenciales.

**Criterios de aceptacion:**

- Cada decisión relevante tiene ADR.
- Si una decisión cambia, se crea un nuevo ADR y no se borra el anterior.

### RNF-009 — Tests de reglas criticas

**Estado:** PLANIFICADO

**Descripción:** Las reglas de seguridad, permisos, estados y contraseñas deben tener tests.

**Criterios de aceptacion:**

- Reglas de dominio tienen tests unitarios.
- Flujos HTTP criticos tienen tests con `httptest` cuando corresponda.

### RNF-010 — Handlers delgados

**Estado:** PLANIFICADO

**Descripción:** Los handlers HTTP deben limitarse a transporte, validación de entrada basica, llamada a casos de uso y respuesta.

**Criterios de aceptacion:**

- Reglas de negocio viven en dominio o aplicación.
- Handlers no acceden directamente a GORM.

### RNF-011 — Claridad de errores

**Estado:** PLANIFICADO

**Descripción:** Los errores deben ser claros para clientes sin filtrar detalles sensibles.

**Criterios de aceptacion:**

- Errores de autenticación no revelan si email o contraseña fue el dato incorrecto.
- Errores de validación indican campos invalidos cuando sea seguro hacerlo.

### RNF-012 — Separacion de responsabilidades

**Estado:** PLANIFICADO

**Descripción:** DTOs HTTP, modelos de dominio y modelos GORM deben mantenerse separados.

**Criterios de aceptacion:**

- DTOs HTTP no se usan como entidades de dominio.
- Modelos GORM no contienen reglas de negocio.
