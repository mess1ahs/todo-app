# Aplicación de Tareas con Go y Docker

Este repositorio contiene una aplicación simple de lista de tareas (TODO list) implementada en Go y empaquetada en un contenedor Docker.

## Requisitos

Asegúrate de tener los siguientes requisitos instalados en tu sistema:

- Go: Versión 1.21.3 o superior.
- Docker: Para construir y ejecutar la aplicación en un contenedor.

## Instrucciones

### Construir y ejecutar la aplicación localmente

1. Clona este repositorio:

   ```bash
   git clone <URL_DEL_REPOSITORIO>
   cd todo-app

2. Ejecuta el build del proyecto y ejecuta el container

    ```bash
    docker compose up --build

## Documentacion:

1. Endpoints de la API (Open API)
Archivo: 

La aplicación expone los siguientes endpoints de la API:

POST /tasks: Crea una nueva tarea.
GET /tasks: Obtiene la lista de tareas existentes.
GET /tasks/{id}: Obtiene una tarea específica por ID.
PUT /tasks/: Actualiza una tarea existente por ID.
DELETE /tasks/{id}: Elimina una tarea por ID.


2. Coleccion Postman

Archivo: todo-app.postman_collection.json

3. Estructura del proyecto

main.go: El archivo principal que inicia el servidor de la aplicación.
models/task.go: Define el modelo de tarea y las funciones para interactuar con la base de datos.
Dockerfile: El archivo de Docker para construir una imagen de contenedor para la aplicación.