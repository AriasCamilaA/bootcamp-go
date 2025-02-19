// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener
// cada una de las letras por separado para deletrearla. Para eso tendrán que:

// Crear una aplicación que reciba una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
// Luego, imprimir cada una de las letras.

package main

import "fmt"

func lettersInWord() {
	var word string
	fmt.Println("Ingrese una palabra:")
	fmt.Scan(&word)
	fmt.Println("La palabra tiene", len(word), "letras.")
	fmt.Println("Las letras de la palabra son:")
	for i := 0; i < len(word); i++ {
		fmt.Println("  ", i+1, "->", string(word[i]))
	}
}

// func main() {
// 	lettersInWord()
// }
