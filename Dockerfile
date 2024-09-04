# Dockerfile

# Etapa de compilación
FROM golang:1.20.6 as builder

WORKDIR /app

# Copia los archivos de gestión de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente del proyecto
COPY . .

# Compila el ejecutable
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway ./cmd/server/main.go

# Etapa de ejecución
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia el ejecutable desde la etapa de compilación
COPY --from=builder /app/gateway .

# Expone el puerto que tu aplicación utiliza
EXPOSE 8081

# Define la variable de entorno para que Fiber sepa que está en producción
ENV FIBER_PREFORK=true

# Comando para ejecutar la aplicación
CMD ["./gateway"]
