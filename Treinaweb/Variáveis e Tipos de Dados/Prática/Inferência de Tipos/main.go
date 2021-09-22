package main

import (
	"fmt"
)

func main() {
	//Por padrão as variáveis que não recebem tipagem são tipadas como string (e ainda são fortemente tipadas)
	nome := "Mário";
	nome = 2;
	fmt.Println(nome);
}