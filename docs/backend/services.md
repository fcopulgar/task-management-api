# Servicios backend

## Estado

`PLANIFICADO`

## Casos de uso esperados

- Login.
- Logout.
- Cambio de contraseña.
- Gestión de usuarios por administrador.
- Gestión de tareas por administrador.
- Consulta de tareas por auditor.
- Consulta y actualización de tareas propias por ejecutor.
- Comentario de tareas vencidas propias.

## Servicios tecnicos esperados

- Servicio de hashing bcrypt.
- Servicio de JWT.
- Servicio de sesiones revocables.
- Repositorios de usuarios, sesiones, tareas y comentarios.

## Reglas de organización

- Los servicios de aplicación coordinan casos de uso.
- Las reglas de negocio viven en dominio y aplicación.
- Los handlers HTTP no contienen reglas de negocio.
- Los repositorios no contienen reglas de negocio.
