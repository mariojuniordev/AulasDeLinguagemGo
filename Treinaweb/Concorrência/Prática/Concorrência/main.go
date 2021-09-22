package main

import (
	"fmt"
)

func main() {
	var limite int
	//Os canais são usados para transmitir informações de dentro
	//das Go Routines para fora e também enviar informações de
	//fora para dentro das Go Routines

	//os channels (chan) são canais de comunicação através dos
	//quais as go routines podem receber valores ou repassar
	//valores
	//Obs.: Canais são SÍNCRONOS
	canal1 := make(chan int, 2000)
	canal2 := make(chan int, 500)
	fmt.Print("Informe um limite: ")
	fmt.Scanf("%d", &limite)
	//Para que uma função em Go seja executada como assíncrona
	//é necessário usar a palavra reservada "go" na frente de
	//sua chamada
	go conteAte(limite, canal1)
	go func(n int) {
		resultado := 0
		for i := 0; i <= n*10; i++ {
			fmt.Printf(" - [conteAte] O número %d \n", i)
			resultado = i * 10
		}
		canal2 <- resultado
		//para atribuir valores a canais é necessário usar
		//a <-
	}(limite)
	for i := 0; i <= limite; i++ {
		fmt.Printf(" - [main] O número é %d \n", i)
	}
	//para encapsular um canal dentro de uma variável, é neces-
	//sário usar a syntax abaixo:
	//<nomeDaVariavel> := <- <nomeDoCanal>
	resultadoCanal1 := <-canal1
	resultadoCanal2 := <-canal2

	//Também é possível fazer a chamada de canais usando a
	//syntax abaixo
	fmt.Printf("Canal 1 = %d, Canal 2 = %d \n", resultadoCanal1, resultadoCanal2)
	//Obs.: Funçõe assícronas são funções q são executadas
	//lado a lado
	//Funções paralelas são funções que dividem o hardware
	//da máquina em paralelo para executar um algortimo
	//O método nativo .Sleep() da library time atrasa em um
	//determinado tempo a execução de uma determinada função
	//O .Second da library time representa 1 segundo)
	//time.Sleep(10 * time.Second)
}

func conteAte(limite int, canal chan int) {
	resultado := 0
	for i := 0; i <= limite*20; i++ {
		fmt.Printf(" - [conteAte] O número %d \n", i)
		resultado = i * 20
	}
	canal <- resultado
}
