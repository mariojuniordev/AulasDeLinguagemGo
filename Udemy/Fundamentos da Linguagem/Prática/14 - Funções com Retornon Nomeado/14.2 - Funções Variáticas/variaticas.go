package main

import (
	"fmt"
)

//Para que uma função receba n parâmetros podemos usar o operador ...
//antes do tipo do parâmetro declaradado
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4)
	fmt.Println(totalDaSoma)
	escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
}
