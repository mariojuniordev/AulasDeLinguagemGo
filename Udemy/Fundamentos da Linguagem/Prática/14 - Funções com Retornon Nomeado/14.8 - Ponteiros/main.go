package main

import (
	"fmt"
)

func inverterSinal(numero int) int {
	return numero * -1
}

//O * antes da variavel faz passar o ponteiro e não necessariaemten a
//variável em si
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
