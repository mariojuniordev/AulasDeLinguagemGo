package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	//Iicilizando um valor de forma lietral no map
	funcsESalarios["Rafael Filho"] = 1350.0
	//Tentando excluir um elemento q não existe
	delete(funcsESalarios, "Inexistente")

	//Imprime o nome e o salário
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}