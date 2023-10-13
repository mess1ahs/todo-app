# Primera etapa: Compilación
FROM golang:1.21.3 AS build

# Establece CGO_ENABLED=1
ENV CGO_ENABLED=1

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el código fuente y el archivo go.mod/go.sum
COPY . .

# Compila la aplicación
RUN go build -o todo-app

# Segunda etapa: Etapa final
FROM golang:1.21.3

# Instala las dependencias necesarias en la etapa final
RUN apt-get update && apt-get install -y ca-certificates

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el binario compilado desde la etapa de compilación
COPY --from=build /app/todo-app .

# Ejecuta la aplicación
CMD ["./todo-app"]
