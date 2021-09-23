package main

import (
	"fmt"
)

func main(){
	//Para especificar uma capacidade para o canal, usamos a syntax abaixo:
	//<nome> := make(chan <tipo>, <capacidade>)
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	canal <- "Terceiro Valor!"

	mensagem := <- canal
	mensagem2 := <- canal
	mensagem3 := <- canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}