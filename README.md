# Backend REST API en Go (Gin Framework)

Este repositorio contiene un backend REST API desarrollado en **Go**, utilizando el framework **Gin** para manejar las solicitudes HTTP. El objetivo principal de este proyecto es implementar una arquitectura robusta y escalable aplicando principios sólidos de desarrollo.

Para ir al repositorio del frontend haga click aqui. https://github.com/EdgarDev17/board-management-frontend
## Características

- **Principios SOLID**: La aplicación sigue los principios de diseño SOLID, asegurando que el código sea modular, mantenible y fácil de extender.
- **CQRS (Command Query Responsibility Segregation)**: La lógica de la aplicación está dividida en dos partes principales:
  - **Lectura**: Utiliza **MongoDB** para realizar las consultas y obtener los datos.
  - **Escritura**: Utiliza **PostgreSQL** para manejar las operaciones de escritura en la base de datos.
- **Patrón de Diseño Repositorio**: Este patrón asegura que la lógica de acceso a los datos está desacoplada de la lógica de negocio.
- **Inyección de Dependencias**: Las dependencias se inyectan en las diferentes partes de la aplicación para mantener una arquitectura limpia y flexible.

## Estructura del Proyecto

El proyecto está organizado de la siguiente manera:

```
project-root/
├── cmd/
│   └── api/
│       └── main.go           # Punto de entrada de la aplicación
├── internal/
│   ├── domain/              # Entidades de negocio e interfaces
│   │   ├── models/          # Modelos de dominio/entidades
│   │   └── repositories/    # Interfaces de repositorios
│   ├── infrastructure/      # Implementaciones externas
│   │   ├── database/        # Conexiones a bases de datos y sus implementaciones
│   │   └── http/            # Configuración del servidor HTTP
│   ├── usecases/            # Lógica de negocio de la aplicación
│   │   └── services/        # Implementaciones de los servicios
│   └── interfaces/          # Adaptadores de interfaz
│       ├── handlers/        # Controladores HTTP
│       └── middleware/      # Middleware HTTP
├── pkg/                     # Paquetes reutilizables
│   ├── logger/              # Utilidades de registro
│   └── validator/           # Utilidades de validación
├── config/                  # Archivos de configuración
├── migrations/              # Migraciones de bases de datos
├── scripts/                 # Scripts de construcción y despliegue
├── tests/                   # Pruebas de integración y end-to-end
├── .env                     # Variables de entorno
├── .gitignore
├── go.mod
└── README.md

```

# Documentación de Endpoints de la API

## Tableros (Boards)

### Obtener Todos los Tableros

- **Método:** GET
- **Endpoint:** `/api/v1/boards`
- **Descripción:** Recupera una lista de todos los tableros.

### Obtener Tablero por ID

- **Método:** GET
- **Endpoint:** `/api/v1/boards/:id`
- **Descripción:** Recupera un tablero específico por su ID.

### Crear Tablero

- **Método:** POST
- **Endpoint:** `/api/v1/boards`
- **Descripción:** Crea un nuevo tablero.

### Actualizar Tablero

- **Método:** PUT
- **Endpoint:** `/api/v1/boards`
- **Descripción:** Actualiza un tablero existente.

### Eliminar Tablero

- **Método:** DELETE
- **Endpoint:** `/api/v1/boards/:id`
- **Descripción:** Elimina un tablero específico por su ID.

### Obtener Tareas por ID de Tablero

- **Método:** GET
- **Endpoint:** `/api/v1/boards/tasks/:id`
- **Descripción:** Recupera todas las tareas asociadas a un tablero específico.

## Tareas (Tasks)

### Obtener Todas las Tareas

- **Método:** GET
- **Endpoint:** `/api/v1/tasks`
- **Descripción:** Recupera una lista de todas las tareas.

### Obtener Tarea por ID

- **Método:** GET
- **Endpoint:** `/api/v1/tasks/:id`
- **Descripción:** Recupera una tarea específica por su ID.

### Crear Tarea

- **Método:** POST
- **Endpoint:** `/api/v1/tasks`
- **Descripción:** Crea una nueva tarea.

### Actualizar Tarea

- **Método:** PUT
- **Endpoint:** `/api/v1/tasks`
- **Descripción:** Actualiza una tarea existente.

### Eliminar Tarea

- **Método:** DELETE
- **Endpoint:** `/api/v1/tasks/:id`
- **Descripción:** Elimina una tarea específica por su ID.

### Actualizar Estado de Tarea

- **Método:** PATCH
- **Endpoint:** `/api/v1/tasks/state`
- **Descripción:** Actualiza el estado de una tarea existente.
