package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 100
	fmt.Println(numero)

	//o tipo uint só suporta números positivos
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123123123123.54
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//FIM NÚMEROS REAIS

	//STRINGS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//As variáveis do tipo char retornam o valor da letra
	//digitada na tabela ASC
	char := 'A'
	fmt.Println(char)

	//FIM STRINGS

	//Em Go todo tipo de dado é inicializado com 0
	//e para strings é inicializado com string vazia
	var texto string
	fmt.Println(texto)

	//O valor 0 de bool é false
	var booleano1 bool
	fmt.Println(booleano1)

	//o valor 0 de error é nil
	var erro error
	fmt.Println(erro)

	//para criar erros em Go é necessário importar o pacote na-
	//tivo erros
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
