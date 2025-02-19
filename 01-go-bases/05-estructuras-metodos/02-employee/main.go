/*
	Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará
	a gestionar correctamente dichos empleados. Los objetivos son:

	1. Crear una estructura Person con los campos ID, Name, DateOfBirth.
	2. Crear una estructura Employee con los campos: ID, Position y una composición con la estructura Person.
	3. Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión
		de los campos de un empleado.
	4. Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último
		ejecutar el método PrintEmployee().

	Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.
*/

package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("The employee data is: \nID: %d\nPosition: %s\nName: %s\nDateOfBirth: %s\n", e.ID, e.Position, e.Person.Name, e.Person.DateOfBirth)
}

func main() {
	person := Person{111111, "Camila Arias", "14-08-2003"}
	employee := Employee{111111, "Developer", person}

	employee.PrintEmployee()
}
