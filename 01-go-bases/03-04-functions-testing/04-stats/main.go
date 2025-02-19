/*
	Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

	Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.

	Ejemplo:

	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	...
	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)
	...
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/

package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation(opType string) (func(...int) int, error) {
	switch opType {
	case minimum:
		return func(nums ...int) int {
			if len(nums) == 0 {
				return 0
			}
			min := nums[0]
			for _, num := range nums {
				if num < min {
					min = num
				}
			}
			return min
		}, nil
	case average:
		return func(nums ...int) int {
			if len(nums) == 0 {
				return 0
			}
			sum := 0
			for _, num := range nums {
				sum += num
			}
			return sum / len(nums)
		}, nil
	case maximum:
		return func(nums ...int) int {
			if len(nums) == 0 {
				return 0
			}
			max := nums[0]
			for _, num := range nums {
				if num > max {
					max = num
				}
			}
			return max
		}, nil
	default:
		return nil, errors.New("operation not defined")
	}
}

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Minimum: %d\n", minValue)
	fmt.Printf("Average: %d\n", averageValue)
	fmt.Printf("Maximum: %d\n", maxValue)
}
