package main

import (
	"fmt"
	"time"
)

// Channel (canal) - É uma forma de comunicação entre go routines
//Channel é um tipo de dado

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}



func main() {
	//Canais servem para executar determinadas linhas de código quando a recepção de determinados dados estiver
	//pronta
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <- c, <- c //recebendo dados do canal

	fmt.Println(a,b)

	fmt.Println(<-c)

}