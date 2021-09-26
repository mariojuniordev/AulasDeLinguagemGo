package main

import "fmt"

//Funções Variáticas são funções que recebem parâmetros variáveis
func media(numeros ...float64) float64 {

	total := 0.0
	for _, num := range numeros{
		total += num
	}

	return total / float64(len(numeros))
	
}

func main(){
	fmt.Printf("Média: %.2f", media(7,4,5,10))
}