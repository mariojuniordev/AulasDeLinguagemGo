package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var waitGroup sync.WaitGroup

	//waitGroup.Add(<número de Go Routines que vão rodar ao mesmo tempo>)
	waitGroup.Add(2)

	//Função anônima
	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1 - Tira 1 go routine do contador waitGroup.Add()
	}()

	//Função anônima
	go func(){
		escrever("Programando em Go!")
		waitGroup.Done() // -1 - Tira 1 go routine do contador waitGroup.Add()
	}()
	
	//Função anônima
	go func(){
		escrever("Go Routine 3!")
		waitGroup.Done() // -1 - Tira 1 go routine do contador waitGroup.Add()
	}()

			//Função anônima
	go func(){
		escrever("Go Routine 4!")
		waitGroup.Done() // -1 - Tira 1 go routine do contador waitGroup.Add()
	}()

	waitGroup.Wait()

}

func escrever(texto string){
	for i:= 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}