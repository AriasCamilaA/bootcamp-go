package main

import (
	"errors"
	"fmt"
)

func MonthlySalary(hours int, amountPerHour float64) (salary float64, err error) {
	if hours < 80 {
		err = errors.New("Error: the worker cannot have worked less than 80 hours per month")
		return
	}

	salary = float64(hours) * amountPerHour

	if salary >= 150000 {
		salary -= salary * 0.1
	} else {
		err = fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %.2f", salary)
	}
	return
}

func main() {
	hours := 100
	amountPerHour := 1600.0

	salary, err := MonthlySalary(hours, amountPerHour)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The calculated salary is: %.2f\n", salary)
	}
}
