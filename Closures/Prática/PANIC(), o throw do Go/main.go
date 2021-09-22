package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bem-vindo!");
	defer fmt.Println("Obrigado por utilizar o nosso software!!!");
	fmt.Print("Digite um número maior que 10: ");
	var numero int;
	fmt.Scanf("%d", &numero);
	if numero <= 10 {
		//panic() é o equivalente ao throw das demais linguagens
		//É utilizado para informar situações de erro ao usuário
		//A execução do panic (throw) não impacta a execução do defer(finally)
		//O defer é executado normalmente
		panic("Número inválido!!!")
	} else {
		fmt.Println("Você digitou um bom número! :)");
	}
}