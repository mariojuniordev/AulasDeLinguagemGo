package main

import (
	"fmt"
)

func main() {
	//Funções anônimas são funções "sem nome"
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}
