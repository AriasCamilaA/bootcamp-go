// Realizar una aplicación que reciba  una variable con el número del mes.

// Según el número, imprimir el mes que corresponda en texto.
// ¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
// Ej: 7, Julio.

// Nota: Validar que el número del mes, sea correcto.

package main

import "fmt"

// func main() {
func month() {
	var months = map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	var month int
	fmt.Println("Ingrese el número del mes:")
	fmt.Scan(&month)

	if month < 1 || month > 12 {
		fmt.Println("El número del mes es incorrecto.")
		return
	}

	fmt.Println("El mes es:", months[month])
}

// func main() {
// 	month()
// }
