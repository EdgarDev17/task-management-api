# Backend REST API en Go (Gin Framework)

Este repositorio contiene un backend REST API desarrollado en **Go**, utilizando el framework **Gin** para manejar las solicitudes HTTP. El objetivo principal de este proyecto es implementar una arquitectura robusta y escalable aplicando principios sólidos de desarrollo.

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
