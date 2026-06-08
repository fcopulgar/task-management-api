# Limites de modulos

## Estado

`PLANIFICADO`

## Modulos previstos

- `domain`: reglas centrales sin dependencias externas.
- `application`: casos de uso y puertos.
- `adapters/inbound/http`: handlers y rutas HTTP.
- `adapters/outbound/persistence`: GORM y PostgreSQL.
- `adapters/outbound/security`: bcrypt y JWT.
- `config`: configuración técnica.

## Dependencias permitidas

- HTTP puede depender de aplicación.
- Aplicación puede depender de dominio y puertos.
- Adapters outbound implementan puertos de aplicación.
- Infraestructura puede depender de librerias externas.

## Dependencias prohibidas

- Dominio no depende de HTTP.
- Dominio no depende de GORM.
- Aplicación no depende de GORM.
- Handlers no acceden directamente a base de datos.
- Repositories no contienen reglas de negocio.
- DTOs HTTP no reemplazan entidades de dominio.

## Contratos entre modulos

- La capa de aplicación expone casos de uso.
- Los adapters inbound traducen HTTP a inputs de casos de uso.
- Los adapters outbound traducen puertos a implementaciones concretas.
- Los mappers separan modelos GORM de entidades de dominio.
