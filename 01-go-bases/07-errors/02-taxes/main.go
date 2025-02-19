/*
	En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

	Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: salary
	is less than 10000" y lanzalo en caso de que “salary” sea menor o igual a  10000. La validación
	debe ser hecha con la función “Is()” dentro del “main”.
*/

package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
	message string
}

func (e *SalaryError) Error() string {
	return e.message
}

var ErrSalaryTooLow = &SalaryError{message: "Error: salary is less than 10000"}

func main() {
	salary := 9000

	if salary <= 10000 {
		err := ErrSalaryTooLow
		if errors.Is(err, ErrSalaryTooLow) {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Salary is sufficient")
	}
}
