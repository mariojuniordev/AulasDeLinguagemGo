package main

import (
	"fmt"
)

//Funções Recursivas são funções que "chamam" a si mesmas
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	fmt.Println("Funções Recursivas")

	//uint() converte um número para o tipo uint
	posicao := uint(10)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
