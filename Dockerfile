# Etapa de construcción
FROM golang:1.24.1-alpine

# Establecemos el directorio de trabajo
WORKDIR /app

# Instalamos las dependencias necesarias para Go (git, gcc, musl-dev)
RUN apk add --no-cache git gcc musl-dev

# Copiamos los archivos go.mod y go.sum para que se descarguen las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiamos todo el código fuente al contenedor
COPY . .

# Construimos la aplicación Go y generamos el binario 'app'
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

# Comando para ejecutar la aplicación
CMD ["./app"]
