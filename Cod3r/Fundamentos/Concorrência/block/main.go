package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação de block
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println(<-c) //operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) //deadlock - Não há dado a ser enviado para esse canal pq todas as Go Routines já "morreram"
	fmt.Println("Fim")
}
