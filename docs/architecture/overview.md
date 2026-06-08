# Vision general de arquitectura

## Estado

`PLANIFICADO`

La arquitectura objetivo es hexagonal minimalista para mantener dominio y aplicación desacoplados de frameworks, HTTP y persistencia.

## Capas objetivo

- Dominio: entidades, reglas, roles, estados, transiciones y errores de negocio.
- Aplicación: casos de uso, puertos internos y coordinacion de reglas.
- Adapters inbound: handlers HTTP con `chi`.
- Adapters outbound: persistencia PostgreSQL con GORM, hashing bcrypt y emisión/validación JWT.

## DISEÑO OBJETIVO — Arquitectura hexagonal

```mermaid
flowchart LR
    HTTP[HTTP chi] --> APP[Casos de uso]
    APP --> DOMAIN[Dominio]
    APP --> PORTS[Puertos outbound]
    PORTS --> GORM[Adapter GORM]
    PORTS --> SECURITY[Adapter seguridad]
    GORM --> PG[(PostgreSQL)]
```

## Persistencia objetivo

- PostgreSQL como base de datos.
- GORM como ORM.
- `AutoMigrate` para inicialización de esquema.
- Modelos GORM separados de modelos de dominio.
- GORM solo en adapters outbound.

## Seguridad objetivo

- JWT con `session_id`.
- Sesiones persistidas y revocables.
- Logout mediante revocación de sesión.
- bcrypt para hash de contraseñas.
