/*
	Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
	se implemente “errors.New()”.
*/

package main

import (
	"errors"
	"fmt"
)

var ErrSalaryTooLow = errors.New("Error: salary is less than 10000")

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
