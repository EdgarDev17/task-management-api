# backend/Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

# Copiar go.mod y go.sum
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar la aplicación
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]

# docker-compose.yml
version: '3.8'

services:
  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
    depends_on:
      - backend
    environment:
      - REACT_APP_API_URL=http://localhost:8080

  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    ports:
      - "8080:8080"