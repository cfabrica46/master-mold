# Configurar la imagen base de Alpine
FROM golang:alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo go.mod y go.sum y descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo el código fuente del proyecto al contenedor
COPY . .

# Instalar herramientas adicionales necesarias para compilar y ejecutar la aplicación
RUN apk add --no-cache git

# Compilar la aplicación
RUN go build -o app .

# Exponer el puerto que utiliza la aplicación
EXPOSE 8080

# Establecer el comando por defecto para ejecutar la aplicación
CMD ["./app"]
