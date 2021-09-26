package main

import "fmt"

func main() {
	//homogênea ( mesmo tipo) e estática (fixo)

	var notas [3]float64
	//Por padrão, os valores de um array inicalizado são zeros
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.10

	fmt.Println(notas)

	total := 0.0
	for i:= 0; i< len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("A sua média foi %.2f", media)

}