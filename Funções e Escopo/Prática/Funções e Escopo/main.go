package main

import (
	"fmt"
)

func main() {
	var numero1, numero2 int;
	var operacao string;

	fmt.Print("Digite o primeiro número: ");
	fmt.Scanf("%d\n", &numero1);
	fmt.Print("Digite a operação (+, -, *, /, $): ");
	fmt.Scanf("%s\n", &operacao);
	fmt.Print("Digite o segundo número: ");
	fmt.Scanf("%d\n", &numero2);

	if operacao == "+" {
		somar(numero1, numero2);
	} else if operacao == "-" {
		resultado := subtrair(numero1, numero2);
		fmt.Printf("%d - %d = %d\n", numero1, numero2, resultado);
	} else if operacao == "*" {
		resultado := multiplicar(numero1, numero2);
		fmt.Printf("%d * %d = %d\n", numero1, numero2, resultado);
	} else if operacao == "/" {
		resultado, resto := dividir(numero1, numero2);
		fmt.Printf("QUOCIENTE = %d; RESTO = %d\n", resultado, resto);
	} else if operacao == "$" {
		inc, dec := incrementoDecremento(numero1, numero2);
		fmt.Printf("O incremento é: %d e o decremento é: %d\n", inc, dec);
	} else {
		fmt.Println("Operação Inválida!!!")
	}

}

//Função Void (Sem Retorno)
func somar(n1 int, n2 int) {
	fmt.Printf("%d + %d = %d", n1, n2, n1+n2);
}

//Em funções com retorno em Go é necessário colocar o tipo do dado q será retornado após os parâmetros
//Syntax: func <nome da função>(<parâmetros com tipos>) <tipo do dado retornado> {<oque a função faz>} 
func subtrair(n1 int, n2 int) int {
	return n1 - n2;
}

//Também podemos utilizar a seguinte syntax para funções com retorno
//Syntax: func <nome da função>(<parâmetros>) (<nome da variável de retorno e tipo>){<oque a função faz>}
func multiplicar(n1 int, n2 int) (resultado int) {
	resultado = n1 * n2;
	return
}

//Também é possível criar funções com 2 retornos utilizando a mesma syntax q vimos anteriormente
func dividir(n1 int, n2 int)(int, int){
	quociente := n1 / n2;
	resto := n1%n2;

	return quociente, resto;

}

func incrementoDecremento(n1 int, n2 int)(inc int, dec int){
	inc = n1 + n2;
	dec = n1 - n2;

	return
}