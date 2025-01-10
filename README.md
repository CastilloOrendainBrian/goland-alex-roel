# Curso Golang Alex Roel

Aprende lenguaje de Go, Desarrollo Web con Go, Manejo de HTML, CSS, API RESTful com MySQL y ORM con Go

## Instalación docker

```bash
docker build -t ubuntu-golang .
docker run -it --name ubuntu-golang-container -v $(pwd)/go:/root/go ubuntu-golang
docker start -ai ubuntu-golang-container