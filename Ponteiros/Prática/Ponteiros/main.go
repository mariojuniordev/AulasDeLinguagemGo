package main

import (
	"fmt"
)

func main() {
	//var saldo float64
	//fmt.Print("Digite o seu saldo: ")
	//fmt.Scanf("%f\n", &saldo)

	//Para atribuir a criação de uma variável a um ponteiro
	//é necessário usar a seguinte syntax:
	//var <variavel> *<tipo> = new(<tipo>)
	var saldo *float64 = new(float64)
	fmt.Print("Digite o seu saldo: ")
	//Agora que a variável já aponta para a posição de memória
	//em sua inicalização, não precisamos mais do &
	//fmt.Scanf("%f\n", &saldo)
	fmt.Scanf("%f\n", saldo)
	//É necessário usar o & antes do nome do parâmetro para que
	//o sistema reconheça que o que está sendo passado é a
	//posição de memória da variável saldo. Seu valor já está
	//alocado por conta do Scanf()
	//calcularRendimento(&saldo)
	//Após fazer a referência à posição de memória na
	//inicialização da var não precisamos mais do &
	calcularRendimento(saldo)
	//É necessário colocar o * na frente da variavel saldo
	//para indicar que queremos o VALOR do ponteiro
	fmt.Printf("Seu saldo com rendimentos é de %.2f \n", *saldo)
}

//Para que esta função funcione corretamente é necessário
//referenciar a posição de memória da variável e não o seu valor
//diretamente, para isso, devemos usar o * antes dos tipos dos
//parâmetros e antes dos nomes de cada variável que queremos
//referenciar
func calcularRendimento(saldo *float64) {
	//o * na frente da variável saldo indica que queremos o
	//VALOR do nosso ponteiros
	*saldo += *saldo * 0.03
}
