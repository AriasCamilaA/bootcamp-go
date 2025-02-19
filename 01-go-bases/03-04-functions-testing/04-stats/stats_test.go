/*
	Los profesores de la universidad de Colombia, entraron al programa de análisis de datos  de Google, el cual premia a los mejores estadísticos de la región. Por ello los profesores nos solicitaron comprobar el correcto funcionamiento de nuestros cálculos estadísticos. Se solicita la siguiente tarea:

	Realizar test para calcular el mínimo de calificaciones.
	Realizar test para calcular el máximo de calificaciones.
	Realizar test para calcular el promedio de calificaciones.
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	t.Run("minimum", func(t *testing.T) {
		minFunc, err := operation(minimum)
		assert.NoError(t, err)
		assert.Equal(t, 2, minFunc(2, 3, 3, 4, 10, 2, 4, 5))
	})

	t.Run("average", func(t *testing.T) {
		averageFunc, err := operation(average)
		assert.NoError(t, err)
		assert.Equal(t, 3, averageFunc(2, 3, 3, 4, 1, 2, 4, 5))
	})

	t.Run("maximum", func(t *testing.T) {
		maxFunc, err := operation(maximum)
		assert.NoError(t, err)
		assert.Equal(t, 10, maxFunc(2, 3, 3, 4, 10, 2, 4, 5))
	})

	t.Run("undefined operation", func(t *testing.T) {
		_, err := operation("undefined")
		assert.Error(t, err)
	})
}
