/*
	La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados al momento de depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos de estos impuestos están correctos. Para esto nos encargaron el trabajo de realizar los test correspondientes para:

	Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
	Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
	Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTaxSalary(t *testing.T) {

	t.Run("Empleado gane por debajo de $50.000", func(t *testing.T) {
		// ARRANGE
		in := float64(45000)
		expected := float64(0)
		// ACT
		result := TaxSalary(in)
		// ASSERT
		require.Equal(t, expected, result)
	})

	t.Run("Empleado gane por encima de $50.000", func(t *testing.T) {
		// ARRANGE
		in := float64(52000)
		expected := in * 0.17
		// ACT
		result := TaxSalary(in)
		// ASSERT
		require.Equal(t, expected, result)
	})

	t.Run("Empleado gane por encima de $150.000", func(t *testing.T) {
		// ARRANGE
		in := float64(152000)
		expected := in * 0.27
		// ACT
		result := TaxSalary(in)
		// ASSERT
		require.Equal(t, expected, result)
	})

}
