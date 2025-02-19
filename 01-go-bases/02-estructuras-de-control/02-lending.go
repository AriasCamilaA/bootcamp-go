// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
// Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
// Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
// tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará
// interés a los que posean un sueldo superior a $100.000.

// Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.

// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

package main

import "fmt"

func lending() {

	var age, old, cash int
	var work bool

	fmt.Println("Ingrese su edad:")
	fmt.Scan(&age)
	fmt.Println("¿Trabaja? 1:Si 0:No")
	fmt.Scan(&work)

	if work {
		fmt.Println("¿Hace cuánto tiempo trabaja en su empleo actual?:")
		fmt.Scan(&old)
		fmt.Println("¿Cuál es su sueldo?:")
		fmt.Scan(&cash)
	}

	if age > 22 && work && old > 1 {
		if cash > 100000 {
			fmt.Println("Prestamo otorgado sin interes")
		} else {
			fmt.Println("Prestamo otorgado")
		}
	} else {
		fmt.Println("No cumples con los requisitos")
	}
}

// func main() {
// 	lending()
// }
