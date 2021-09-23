package main

import (
	"fmt"
	"time"
)

func main(){
	//CONCORRÊNCIA != PARALELISMO
	//As Go Routines fazem que 2 funções sejam executadas lado a lado
	go escrever("Hello World!") //goroutine
	escrever("Proramando em Go!")
}

func escrever(texto string){
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}