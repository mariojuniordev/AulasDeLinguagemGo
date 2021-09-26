package main

import "fmt"

// imprimirAprovados recebe como parâmetro nomes de aprovados e os retorna de forma ordenada
func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}

}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Mário"}
	//Passando um slice como parâmetro de uma função
	//É necessário usar os ... depois do nome do slice, para que sejam passados todos os seus valores
	//como parâmtero de uma determinada função
	imprimirAprovados(aprovados...)
}