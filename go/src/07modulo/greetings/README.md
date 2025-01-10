# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personales

## Instalación
Ejecuta el siguiente comando para instalar el paquete:
```bash
go get -u github.com/CastilloOrendainBrian/goland-alex-roel/tree/main/go/src/07modulo/greetings
```

## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código

```go
import (
	"fmt"

	"github.com/CastilloOrendainBrian/goland-alex-roel/tree/main/go/src/07modulo/greetings"
)

func main() {
	messages, err := greetings.Hellos("Juan")

	if err != nil {
		fmt.Println("Ocurrió un error:", err)
        return
	}

	fmt.Println(messages)
}

```
Este ejemplo importa el paquete github.com/CastilloOrendainBrian/goland-alex-roel/tree/main/go/src/07modulo/greetings y llama a la función Hello saludo personalizado, Si ocurre un error, se imprime un mensaje de error.