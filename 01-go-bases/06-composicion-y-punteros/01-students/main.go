/*
	Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el
	detalle de los datos de cada uno de ellos/as, de la siguiente manera:


	Name: [Nombre del alumno]
	Apellido: [Apellido del alumno]
	DNI: [DNI del alumno]
	Fecha: [Fecha ingreso alumno]


	Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.

	Para ello es necesario generar una estructura Alumnos con las variables Nombre, A
	pellido, DNI, Fecha y que tenga un método detalle
*/

package main

import (
	"fmt"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Printf("Name: %s\nApellido: %s\nDNI: %s\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	alumno := Alumno{
		Nombre:   "Juan",
		Apellido: "Perez",
		DNI:      "12345678",
		Fecha:    "01/01/2020",
	}
	alumno.detalle()

	alumno2 := Alumno{
		Nombre:   "Maria",
		Apellido: "Gomez",
		DNI:      "87654321",
		Fecha:    "02/02/2021",
	}
	alumno2.detalle()
}
