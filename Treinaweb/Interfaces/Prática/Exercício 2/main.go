package main

import (
	"fmt"
)

type Buzinador interface {
	buzinar() string
}

type Carro struct {
	marca string
}

func (carro *Carro) buzinar(barulho string) string {
	if barulho == nil {
		return " - Buzina: BiBi"
	}
	return " - Buzina: " + barulho
}

func main() {
	carro := Carro{}
	fmt.Println(chamarBuzina(&carro))
}

func chamarBuzina(buzinador Buzinador) string {
	return buzinador.buzinar()
}
