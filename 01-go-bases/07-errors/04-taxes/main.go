/*
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba
por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por
consola deberá decir: “Error: the minimum taxable amount is 150,000 and the salary entered is: [salary]”,
siendo [salary] el valor de tipo int pasado por parámetro).
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
		err := fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
		fmt.Println(err)
	} else {
		fmt.Println("Salary is sufficient")
	}
}
