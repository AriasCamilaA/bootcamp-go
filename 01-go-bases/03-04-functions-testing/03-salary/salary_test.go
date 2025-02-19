/*
	La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos de los salarios, por ello nos piden realizar una serie de tests sobre nuestro programa. Necesitaremos realizar las siguientes pruebas en nuestro código:

	Calcular el salario de la categoría “A”.
	Calcular el salario de la categoría “B”.
	Calcular el salario de la categoría “C”.
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSalary(t *testing.T) {

	minutes := 10000

	tests := []struct {
		minutesWorked int
		category      string
		expected      float64
	}{
		{minutes, "A", (float64(minutes)/60)*3000 + ((float64(minutes)/60)*3000)*0.5},
		{minutes, "B", (float64(minutes)/60)*1500 + ((float64(minutes)/60)*1500)*0.2},
		{minutes, "C", (float64(minutes) / 60) * 1000},
	}

	for _, test := range tests {
		result := CalculateSalary(test.minutesWorked, test.category)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
