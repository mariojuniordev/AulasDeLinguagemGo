package main

import (
	"fmt"
)

func main(){
	fmt.Println("Bem-vindo! :)");
	defer func(){
		//A função nativa recover() recupera o último erro ocorrido dentro da função main()
		//Logo após isso escrevemos ; e a condição propriamente dita do IF
		//Syntax: if <nome da variavel> := recover(); <condição> {<execução>}
		if ex := recover(); ex != nil {
			fmt.Printf("Desculpe, ocrreu um erro: %s\n", ex);
		}
		fmt.Println("Obrigado por utilizar o nosso software!!!")
	}() //É necessário abrir e fechar parênteses logo após uma função defer para que ela seja executada
	//corretamente, isso garante que o defer seja chamado e que a execução da função seja um defer(finally)
	var numero int;
	fmt.Print("Digite um número mairo que 10: ");
	fmt.Scanf("%d", &numero)
	if numero <= 10 {
		panic("número é inválido!")
	} else {
		fmt.Println("Você digitou um bom número! :)")
	}
}