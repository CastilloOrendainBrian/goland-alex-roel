# Dockerfile para un contenedor basado en Ubuntu con Golang

# Usar una imagen base de Ubuntu
FROM ubuntu:latest

# Establecer variables de entorno
ENV DEBIAN_FRONTEND=noninteractive \
    LANG=C.UTF-8 \
    PATH="/usr/local/go/bin:$PATH"

# Actualizar y preparar el sistema
RUN apt-get update && apt-get install -y \
    wget \
    curl \
    build-essential \
    git \
    vim \
    && rm -rf /var/lib/apt/lists/*

# Descargar e instalar Go
ENV GO_VERSION=1.21.1
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Crear un directorio de trabajo
WORKDIR /workspace

# Comando por defecto al iniciar el contenedor
CMD ["/bin/bash"]