package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de un contacto
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guarda contactos en una archivo JSON
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

// Carga contactos desde un archivo JSON
func loadContactsFromFile(contacs *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacs)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Slice de contactos
	var contacts []Contact

	// Cargar contactos existentes desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos:", err)
	}

	// Crear instancia de bufio
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostar opciones al usuario
		fmt.Print("==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar contacto\n",
			"2. Mostrar contactos\n",
			"3. Salir\n",
			"Elige una opción:")

		// Leer la opción del usuario
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
			return
		}

		// Manejar la opción del usuario
		switch option {
		case 1:
			// ingresar y crear un nuevo contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')

			// Agregar el contacto al slice
			contacts = append(contacts, c)

			// Guardar los contactos en el archivo
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar los contactos:", err)
			}

		case 2:
			// Mostrar los contactos
			fmt.Println("==== CONTACTOS ====")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("==== FIN DE CONTACTOS ====")

		case 3:
			// Salir del programa
			return

		default:
			fmt.Println("Opción no válida")
		}
	}
}
