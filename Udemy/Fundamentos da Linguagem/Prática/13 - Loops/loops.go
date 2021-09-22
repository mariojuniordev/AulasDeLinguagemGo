package main

import (
	"fmt"
)

func main() {

	/*fmt.Println("Loops!")
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		//O método nativo .Sleep() do objeto time atrasa a execução da
		//próxima linha de código em um determinado tempo
		//o .Second do objeto time representa 1 segundo
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j")
		time.Sleep(time.Second)
	}*/

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
