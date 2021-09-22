package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

//Para criar métodos para uma structure usamos a syntax abaixo
//func (<parâmetro><nome da struct>) <nome do método>() {}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//Para atualizar dados do usuário fora do método é necessário usar
//ponteiros nos parâmetros do método
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 15}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
