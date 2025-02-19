/*
	A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio. Ahora que el archivo sí existe, el panic no debe ser lanzado.

	1. Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
	2. Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga. En el caso de no poder leerlo, se debe lanzar un “panic”.

	Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer” que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	textbar := "-------------------------------"
	path := "customers.txt"

	// Defer para garantizar que siempre se imprima el mensaje de finalización
	defer fmt.Println("\nejecución finalizada")

	// Abrir el archivo
	file, err := os.Open(path)
	if err != nil {
		// Si hay un error al abrir el archivo, hacer panic
		fmt.Printf("New error: %s", err)
		panic("The indicated file was not found or is damaged")
	}

	// Asegurarse de cerrar el archivo al final
	defer file.Close()

	// Leer el contenido del archivo
	content, err := io.ReadAll(file)
	if err != nil {
		// Si ocurre un error al leer el archivo, también hacemos panic
		fmt.Printf("New error: %s", err)
		panic("The indicated file was not found or is damaged")
	}

	// Si llegamos aquí, significa que se pudo leer el archivo correctamente
	fmt.Printf("Customers information found:\n%s\n%s\n%s", textbar, string(content), textbar)
}
