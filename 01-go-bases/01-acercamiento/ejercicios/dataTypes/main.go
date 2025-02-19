// jercicio 4 - Tipos de datos

// Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go, pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:

// Verificar su código y realizar las correcciones necesarias.

//    var lastName string = "Smith"

//    var age int = "35"

//    boolean := "false"

//    var salary string = 45857.90

//    var firstName string = "Mary"

package main

import "fmt"

func main() {
    var lastName string = "Smith"  // Correcto
    var age int = 35               // Corregido, ahora es un int
    boolean := false               // Corregido, ahora es un bool
    var salary float64 = 45857.90  // Corregido, ahora es un float64
    var firstName string = "Mary"  // Correcto

    // Imprimir las variables para verificar que todo funciona correctamente
    fmt.Println(lastName, age, boolean, salary, firstName)
}
