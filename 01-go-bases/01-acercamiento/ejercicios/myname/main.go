// Ejercicio 1 - Imprimí tu nombre
// Tendrás que:

// Crear una aplicación donde tengas como variable tu nombre y dirección.
// Imprimir en consola el valor de cada una de las variables.

package main

import "fmt"

var name string = "Camila Arias"
var adress string = "Carrera 11a #4-33 Sur, Bogotá"

func main() {
    fmt.Printf("Mi nombre es %s y vivo en %s ", name, adress)
}
