/*
	Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

	Categoría C, su salario es de $1.000 por hora.
	Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
	Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.

	Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.
*/

package main

import "fmt"

func CalculateSalary(minutesWorked int, category string) float64 {
	hoursWorked := float64(minutesWorked) / 60
	var hourlyRate float64
	var bonusPercentage float64

	switch category {
	case "A":
		hourlyRate = 3000
		bonusPercentage = 0.50
	case "B":
		hourlyRate = 1500
		bonusPercentage = 0.20
	case "C":
		hourlyRate = 1000
		bonusPercentage = 0.0
	default:
		return 0
	}

	salary := hoursWorked * hourlyRate
	bonus := salary * bonusPercentage
	totalSalary := salary + bonus

	return totalSalary
}

func main() {
	minutesWorked := 9600
	category := "A"

	salary := CalculateSalary(minutesWorked, category)
	fmt.Printf("The salary for category %s with %d minutes worked is: $%.2f\n", category, minutesWorked, salary)
}
