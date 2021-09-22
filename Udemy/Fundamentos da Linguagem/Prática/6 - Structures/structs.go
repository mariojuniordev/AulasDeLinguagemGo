package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	//Para instanciar uma struct colocamos seu nome como tipo da variavel
	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	//Também é possível instanciar structs usando a syntax

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	//declarar apenas 1 atributo
	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
