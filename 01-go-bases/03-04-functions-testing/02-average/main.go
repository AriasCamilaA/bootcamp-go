/*
	Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/

package main

import (
	"errors"
	"fmt"
)

func Average(grades ...int) (float64, error) {
	if len(grades) == 0 {
		return 0, errors.New("no grades provided")
	}

	var sum int
	for _, grade := range grades {
		if grade < 0 {
			return 0, errors.New("negative grades are not allowed")
		}
		sum += grade
	}

	return float64(sum) / float64(len(grades)), nil
}

func main() {
	grades := []int{85, 90, 78, 92, 88}
	avg, err := Average(grades...)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("The average grade is %.2f\n", avg)
}
