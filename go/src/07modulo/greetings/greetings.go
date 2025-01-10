package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {

	if name == "" {
		// return "", errors.New("Nombre vacío")
		return name, errors.New("nombre vacío")

	}
	// Devuelve un saludo que incluye el nombre en un mensaje
	// message := fmt.Sprintf("¡Hola,  %v! ¡Bienvenido!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos devuelve un mapa que asocia cada uno de los nombres con un mensaje de saludo
// en un mapa
func Hellos(names []string) (map[string]string, error) {
	// Un mapa para almacenar los mensajes de saludo
	messages := make(map[string]string)
	// Iterar a través de los valores de la slice names
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// En el mapa, asociar el mensaje de saludo con el nombre
		messages[name] = message
	}
	return messages, nil
}

// randomFormat devuelve uno de varios formatos de mensajes de saludo
// se selecciona de fomra aleatoria
func randomFormat() string {
	// Una slice de mensajes de saludo
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Saludos, %v! ¡Bienvenido!",
		"¡Que onda, %v! ¡Bienvenido!",
	}

	// Devuelve uno de los mensajes seleccionados de forma aleatoria
	return formats[rand.Intn(len(formats))]
}
