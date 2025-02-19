/*
	El refugio de animales envió una queja ya que el cálculo total de alimento a comprar no fue el correcto y compraron menos alimento del que necesitaban. Para mantener satisfecho a nuestro cliente deberemos realizar los test necesarios para verificar que todo funcione correctamente:

	Verificar el cálculo de la cantidad de alimento para los perros.
	Verificar el cálculo de la cantidad de alimento para los gatos.
	Verificar el cálculo de la cantidad de alimento para los hamster.
	Verificar el cálculo de la cantidad de alimento para las tarántulas.

	Ejemplo:

	func TestDog(t *testing.T)

	func TestCat(t *testing.T)
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDog(t *testing.T) {
	animalFunc, msg := animal(dog)
	assert.Equal(t, "", msg)
	assert.NotNil(t, animalFunc)
	assert.Equal(t, 100.0, animalFunc(10))
}

func TestCat(t *testing.T) {
	animalFunc, msg := animal(cat)
	assert.Equal(t, "", msg)
	assert.NotNil(t, animalFunc)
	assert.Equal(t, 50.0, animalFunc(10))
}

func TestHamster(t *testing.T) {
	animalFunc, msg := animal(hamster)
	assert.Equal(t, "", msg)
	assert.NotNil(t, animalFunc)
	assert.Equal(t, 2.5, animalFunc(10))
}

func TestTarantula(t *testing.T) {
	animalFunc, msg := animal(tarantula)
	assert.Equal(t, "", msg)
	assert.NotNil(t, animalFunc)
	assert.Equal(t, 1.5, animalFunc(10))
}

func TestUnknownAnimal(t *testing.T) {
	animalFunc, msg := animal("unknown")
	assert.Nil(t, animalFunc)
	assert.Equal(t, "Animal not found", msg)
}
