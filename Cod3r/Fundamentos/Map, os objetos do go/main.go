package main

import "fmt"

func main() {
	// var aprovados map[int]string
	//mapas devem ser inicalizados usa o método make()
	aprovados :=  make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[91023192839] = "Pedro"
	aprovados[12312312315] = "Ana"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF:%d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12312312315])
	delete(aprovados, 12312312315)
	fmt.Println(aprovados[12312312315])
}