package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else {
		fmt.Println("É menor que zero")
	}

}
