package main

import (
	"fmt"
)

func main() {
	//ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)

	//OPERADORES LÓGIOS
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	//OPERADORES UNÁRIOS
	numero := 10
	numero++     //numero = numero + 1
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--     //numero = numero -1
	numero -= 20 //numero = numero -20

	numero *= 3 //numero = numero * 3

	fmt.Println(numero)
}
