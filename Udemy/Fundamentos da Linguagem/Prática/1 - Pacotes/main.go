package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquvio modulo")
	auxiliar.Escrever()

	//Para referenciar um pacote devemos utilizar o nome que
	//vem depois da última barra, nesse caso, o checkmail
	//erro := checkmail.ValidateFormat("devbook@gmail.com")
	//fmt.Println(erro)
}
