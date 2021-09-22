package auxiliar

import (
	"fmt"
)

func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	//Caso importemos um pacote auxiliar dentro do mesmo pacote
	//não precisamos mudar o estado para public (usar letra
	//maiúscula)
	escrever2()
}
