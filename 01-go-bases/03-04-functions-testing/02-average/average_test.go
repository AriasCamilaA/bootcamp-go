/*
	El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente, por lo que ahora nos corresponde realizar los test correspondientes:

	Calcular el promedio de las notas de los alumnos.
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name        string
		grades      []int
		expected    float64
		expectedErr bool
	}{
		{"No grades", []int{}, 0, true},
		{"Negative grade", []int{85, -90, 78}, 0, true},
		{"Single grade", []int{85}, 85, false},
		{"Multiple grades", []int{85, 90, 78, 92, 88}, 86.6, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Average(tt.grades...)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expected, got)
		})
	}
}
