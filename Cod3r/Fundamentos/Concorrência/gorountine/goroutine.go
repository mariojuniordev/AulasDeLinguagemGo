package main

import (
	"fmt"
	"time"
)

//func fale(<var1>, <var2> <tipo var 1 e var2>, qtde int){}
func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second) //delay
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main(){
	//Exemplo de concorrência
	//fale("Maria", "Pq vc não fala comigo?", 3)
	//fale("João", "Só posso falar depois de vc!", 1)

	//go fale("Maria", "Ei...", 500)
	//go fale("João", "Opa...", 500)

	go fale("Maria", "Entendi!!!", 10) //As Go Routines fazem a função ser executada de forma independente
	fale("João", "Parabéns", 5)
}