package main

import (
	"fmt"
	"time"
)

func main(){
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <- canal //a variável aberto será usada para verificar se o canal está aberto
		if !aberto{
			break //Encerra o loop
		}
	
		fmt.Println(mensagem)
	}

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string){
	for i:= 0; i < 5; i++{
		//Para atribuir valor a um canal usamo sa syntax abaixo
		canal <- texto //Quando o Go chega na execução de um canal ele "encerra" a aplicação
		time.Sleep(time.Second)
	}
	//A função nativa close() faz com que o canal não envie nem receba mais dados
	close(canal)
}