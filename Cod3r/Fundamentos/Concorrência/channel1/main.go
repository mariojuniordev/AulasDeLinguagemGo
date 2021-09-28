package main

import "fmt"

func main() {
	//Syntax p/ criação de channels:
	//<nome> := make(chan <tipo>, <buffer>)
	ch := make(chan int, 1)

	ch <- 1 // ch encapsula o valor 1 - ENTRADA DE DADOS
	<-ch    // SAÍDA DE DADOS - Determinado valor está saindo do canal

	ch <- 2
	fmt.Println(<-ch)

}