// Ejercicio 3 - Declaración de variables

// Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.

// Necesita ayuda para:

// Detectar cuáles de estas variables que declaró el alumno son correctas.

// Corregir las incorrectas.

//    var 1firstName string -> debe llamarse var firstaName string // sin ponerle el 1 al principio

//    var lastName string -> // Esta bien

//    var int age -> debe ser var age int // esta al reves

//    1lastName := 6 -> lastName := "6" // se debe quitar el número al principio

//    var driver_license = true -> var driverLicense = true  // Se debe usar camelCase

//    var person height int -> var personHeight int // El nombre de la variable no deben ser dos palabras

//    childsNumber := 2 -> childrenNumber := 2 // Asi queda correcto en inglés

package main

import "fmt"

func main() {
    var firstName string
    var lastName string
    var age int
    lastName = "6"
    var driverLicense = true
    var personHeight int
    childrenNumber := 2

    fmt.Println(firstName, lastName, age, driverLicense, personHeight, childrenNumber)
}
