package main

import (
	"fmt"
	"go-sandbox/cfernan/dial"
)

// Ejercicio dial
func main() {
	// Ordenación con el algoritmo de Dial
	numeros := []int{3, 9, 2, 1, 8, 9, 13, 17, 0, 4}
	fmt.Println(numeros)
	if err := dial.Sort(numeros); err != nil {
		fmt.Println(" error durante la ordenación")
	} else {
		fmt.Println(numeros)
	}
}
