package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	//Para dizermos que os atributos de um struct devem ficar em uma determinada chave JSON, usamos a syntax
	//abaixo
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	//[]byte(<nome do JSON a ser convertido>) - Converte uma estrutura de dados em Slice de bytes
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Toby","raca":"Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}