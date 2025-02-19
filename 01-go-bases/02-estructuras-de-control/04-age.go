// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

//   var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado también es necesario:

// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del mapa.

package main

import "fmt"

func employees() {

	employees := map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])

	over21 := 0
	for _, age := range employees {
		if age > 21 {
			over21++
		}
	}

	fmt.Println("Cantidad de empleados mayores a 21 años:", over21)

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println(employees)
}

// func main() {
// 	employees()
// }
