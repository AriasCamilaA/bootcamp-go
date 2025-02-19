/*
	El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.
	Los datos requeridos son:
		- File
		- Name
		- ID
		- Phone number
		- Home

		-> Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los datos de un array. En caso de que esté repetido, debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
			1.- generar un panic;
			2.- lanzar por consola el mensaje: “Error: client already exists”, y continuar con la ejecución del programa normalmente.

		-> Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por parámetro algún valor cero (recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).

		-> Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “End of execution” y “Several errors were detected at runtime”. Utilizá defer para cumplir con este requerimiento.

		Requerimientos generales:

		- Utilizá recover para recuperar el valor de los panics que puedan surgir
		- Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
		- Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente para el caso de error retornado).
*/

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Estructura que representa un cliente
type Client struct {
	File  string
	Name  string
	ID    string
	Phone string
	Home  string
}

// Función para verificar si el cliente ya existe
func clientExists(clients []Client, client Client) bool {
	for _, existingClient := range clients {
		if existingClient.ID == client.ID {
			return true
		}
	}
	return false
}

// Función para validar si todos los campos de un cliente tienen valor distinto de cero
func validateClient(client Client) (bool, error) {
	if client.File == "" || client.Name == "" || client.ID == "" || client.Phone == "" || client.Home == "" {
		return false, errors.New("all fields must be filled")
	}
	return true, nil
}

// Función para guardar un nuevo cliente en el archivo
func saveClientToFile(path string, client Client) error {
	// Abrir el archivo en modo "append" (agregar datos al final del archivo)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Escribir los datos del cliente en el archivo
	clientInfo := fmt.Sprintf("%s | %s | %s | %s | %s\n", client.File, client.Name, client.ID, client.Phone, client.Home)
	_, err = file.WriteString(clientInfo)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Array de clientes, podrías cargarlo desde un archivo o base de datos.
	clients := []Client{
		{"file1", "Alice", "1", "123456789", "Street 1"},
		{"file2", "Bob", "2", "987654321", "Street 2"},
	}

	// Defer para garantizar que siempre se imprima el mensaje de finalización
	defer fmt.Println("\nEnd of execution")

	// Simulamos un cliente nuevo a registrar
	newClient := Client{"file3", "Charlie", "3", "555123456", "Street 3"}

	// Recuperar panics
	defer func() {
		if r := recover(); r != nil {
			// Mensaje en caso de panic
			fmt.Println("\nSeveral errors were detected at runtime")
			fmt.Println("Error:", r)
		}
	}()

	// Tarea 1: Verificamos si el cliente ya existe
	if clientExists(clients, newClient) {
		// Si el cliente existe, se lanza un panic pero se maneja con recover
		panic("Error: client already exists")
	}

	// Tarea 2: Validamos si todos los datos del cliente son distintos de cero
	isValid, err := validateClient(newClient)
	if err != nil {
		// Si los datos son inválidos, se lanza un panic con el mensaje de error
		panic(err)
	}

	// Si llegamos aquí, es porque el cliente es válido y no existe previamente
	if isValid {
		// Procedemos con la acción de registrar el cliente, por ejemplo agregándolo a la lista
		clients = append(clients, newClient)
		fmt.Println("New client added successfully:", newClient.Name)

		// Guardamos el nuevo cliente en el archivo
		err = saveClientToFile("customers.txt", newClient)
		if err != nil {
			// Si ocurre un error al guardar en el archivo, lanzamos un panic
			panic("Error saving client to file")
		}
	}

	// Leemos el archivo de clientes para mostrar su información
	path := "customers.txt"
	textbar := "-------------------------------"

	// Abrimos el archivo
	file, err := os.Open(path)
	if err != nil {
		// Si hay un error al abrir el archivo, hacemos panic
		panic("The indicated file was not found or is damaged")
	}

	// Asegurarnos de cerrar el archivo al final
	defer file.Close()

	// Leemos el contenido del archivo
	content, err := io.ReadAll(file)
	if err != nil {
		// Si ocurre un error al leer el archivo, también hacemos panic
		panic("The indicated file was not found or is damaged")
	}

	// Si llegamos aquí, significa que se pudo leer el archivo correctamente
	fmt.Printf("Customers information found:\n%s\n%s\n%s", textbar, string(content), textbar)
}
