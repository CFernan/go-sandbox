package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sumar todos los int entre 0 y 1 millón que no sean múltiplos de 3 y 5")
	suma := 0
	for i := range 1000001 {
		if i%3 > 0 && i%5 > 0 {
			suma += i
		}
	}
	fmt.Printf("Resultado: %d\n", suma)
}
