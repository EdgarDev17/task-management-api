# Usa una imagen oficial de Go como base
FROM golang:latest AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum para aprovechar el cache de dependencias
COPY go.mod go.sum ./

# Descarga las dependencias del proyecto
RUN go mod tidy

# Copia todo el código fuente al contenedor
COPY . .

# Exponer el puerto que usa tu app de Gin
EXPOSE 8080

# Establece el comando por defecto para correr el servidor
CMD ["go", "run", "/cmd/api/main.go"]
