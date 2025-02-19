/*
	En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

	Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: the salary entered does not reach the taxable minimum" y lanzalo en caso de que “salary” sea menor a 150.000. De lo contrario, tendrás que imprimir por consola el mensaje “Must pay tax”.
*/

package main

import (
	"fmt"
)

type TaxesError struct {
	msg string
}

func (e TaxesError) Error() string {
	return e.msg
}

func main() {
	salary := 120000

	if salary < 150000 {
		err := &TaxesError{msg: "Error: the salary entered does not reach the taxable minimum"}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Must pay tax")
	}

}
