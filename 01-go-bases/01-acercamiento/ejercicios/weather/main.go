// Ejercicio 2 - Clima


// Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares.

// Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
// Imprimí los valores de las variables en consola.
// ¿Qué tipo de dato le asignarías a las variables?

package main

import "fmt"

func main() {
    var temp float32 = 16.5
    var hum int = 60
    var pres float32 = 1013.25 

    fmt.Printf("Temperatura: %.2f°C\n", temp)
    fmt.Printf("Humedad: %d%%\n", hum)
    fmt.Printf("Presión atmosférica: %.2f hPa\n", pres)
}
