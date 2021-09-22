package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	//Syntax dos maps:
	//<variavel> := map[<tipo das chaves>]<tipo dos valores>{}
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	//Deletar uma chave de um map
	delete(usuario2, "nome")

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
