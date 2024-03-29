# Aplicación de Tareas con Go y Docker

Este repositorio contiene una aplicación simple de lista de tareas (TODO list) implementada en Go y empaquetada en un contenedor Docker.

## Requisitos

Asegúrate de tener los siguientes requisitos instalados en tu sistema:

- Go: Versión 1.21.3 o superior.
- Docker: Para construir y ejecutar la aplicación en un contenedor.

1 - Instalacion GoLang (https://go.dev/dl/):
 - Descargar version de windows. (https://go.dev/dl/go1.21.3.windows-amd64.msi)
 - Ejecutar el archivo go1.21.3.windows-amd64.msi y seguir los pasos de instalacion.
 - En este caso se realizo la instalacion por defecto.
   (go version go1.21.3 windows/amd64)

2 - Instalacion Docker (https://www.docker.com/get-started/):
 - Descargar version de windows. (https://desktop.docker.com/win/main/amd64/Docker%20Desktop%20Installer.exe?utm_source=docker&utm_medium=webreferral&utm_campaign=dd-smartbutton&utm_location=module&_gl=1*1xi0fco*_ga*MTk3MTEyMDcwMC4xNjk3MTQxMTcw*_ga_XJWPQMJYHQ*MTY5NzE0MTE3MC4xLjEuMTY5NzE0MTE3My41Ny4wLjA.)
 -  Ejecutar archivo "Docker Desktop Installe.exe".
 -  Seguir los  pasos de instalacion.
 -  En este caso se realizo la instalacion por defecto.
   (Docker version 24.0.6, build ed223bc)


## Instrucciones

### Construir y ejecutar la aplicación localmente

1. Clona este repositorio:

   ```bash
   git clone https://github.com/mess1ahs/todo-app.git
   cd todo-app

2. Ejecuta el build del proyecto y ejecuta el container

    ```bash
    docker compose up --build

## Documentacion:

1. Endpoints de la API (Open API)
Archivo: open-api.yml

La aplicación expone los siguientes endpoints de la API:

POST /tasks: Crea una nueva tarea.

GET /tasks: Obtiene la lista de tareas existentes.

GET /tasks/{id}: Obtiene una tarea específica por ID.

PUT /tasks/: Actualiza una tarea existente por ID.

DELETE /tasks/{id}: Elimina una tarea por ID.

Documentacion en la carpeta docs/ (index.html)


2. Coleccion Postman

- Archivo: todo-app.postman_collection.json

3. Estructura del proyecto

- main.go: El archivo principal que inicia el servidor de la aplicación.
- models/task.go: Define el modelo de tarea y las funciones para interactuar con la base de datos.
- Dockerfile: El archivo de Docker para construir una imagen de contenedor para la aplicación.