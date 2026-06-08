# Flujo de datos

## Estado

`PLANIFICADO`

## DISEÑO OBJETIVO — Login

```mermaid
sequenceDiagram
    participant C as Cliente
    participant H as Handler HTTP
    participant U as Caso de uso Login
    participant R as Repositorio Usuarios
    participant S as Sesiones
    C->>H: POST /auth/login
    H->>U: credenciales
    U->>R: buscar usuario activo
    U->>S: crear sesion
    U-->>H: JWT con session_id
    H-->>C: token
```

## DISEÑO OBJETIVO — Logout

```mermaid
sequenceDiagram
    participant C as Cliente
    participant M as Middleware Auth
    participant U as Caso de uso Logout
    participant S as Sesiones
    C->>M: POST /auth/logout con JWT
    M->>S: validar session_id no revocado
    M->>U: sesion actual
    U->>S: marcar revoked_at
    U-->>C: sesion cerrada
```

## DISEÑO OBJETIVO — Cambio obligatorio de contrasena

```mermaid
flowchart TD
    A[Request protegido] --> B{must_change_password?}
    B -- si --> C{Ruta permitida?}
    C -- password/logout --> D[Permitir]
    C -- otra ruta --> E[Rechazar]
    B -- no --> D
```

## DISEÑO OBJETIVO — Tareas por administrador

```mermaid
flowchart TD
    A[ADMIN] --> B[Crear tarea]
    B --> C{Assignee es EXECUTOR?}
    C -- no --> R[Rechazar]
    C -- si --> D[Guardar ASSIGNED]
    A --> E[Actualizar o eliminar]
    E --> F{Estado ASSIGNED?}
    F -- no --> R
    F -- si --> G[Aplicar cambio]
```

## DISEÑO OBJETIVO — Tareas por ejecutor

```mermaid
flowchart TD
    A[EXECUTOR] --> B[Consultar tarea]
    B --> C{Es propia?}
    C -- no --> R[Rechazar]
    C -- si --> D[Permitir lectura]
    A --> E[Cambiar estado]
    E --> F{Vencida?}
    F -- si --> R
    F -- no --> G{Transicion valida?}
    G -- si --> H[Actualizar estado]
    G -- no --> R
```

## DISEÑO OBJETIVO — Auditor

```mermaid
flowchart TD
    A[AUDITOR] --> B[Listar tareas]
    A --> C[Ver detalle]
    A --> D[Intentar escribir]
    D --> E[Rechazar]
```

## DISEÑO OBJETIVO — Estados de tarea

```mermaid
stateDiagram-v2
    [*] --> ASSIGNED
    ASSIGNED --> STARTED
    STARTED --> WAITING
    STARTED --> FINISHED_SUCCESS
    STARTED --> FINISHED_ERROR
    WAITING --> WAITING
    WAITING --> FINISHED_SUCCESS
    WAITING --> FINISHED_ERROR
```
