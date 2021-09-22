package main

import (
	"fmt"

	estrutura "treinaweb.com.br/mariojuniordev/carros"
)

//Assim como na criação de Structures, nas interfaces iniciamos
//com a palavra reservada type, depois o nome da interface e,
//em seguida, o seu tipo (interface), segue Syntax:
//type <nome da interface> interface {<métodos>}

func main() {
	//Em Go não é necessário usar classes construtoras (new)
	//basta encapsular a instância da structure dentro
	//de uma variável
	carro := estrutura.Carro{Modelo: "Palio", Marca: "Fiat"}
	opcao := 0
	for opcao != 3 {
		fmt.Println("O que você deseja fazer?")
		fmt.Println(" 1 - Acelerar")
		fmt.Println(" 2 - Frear")
		fmt.Println(" 3 - Sair")
		fmt.Scanf("%d\n", &opcao)

		var err error = nil

		switch opcao {
		case 1:
			//Para referenciarmos o método de uma inteface é
			//necessário usar o ponteiro (endereço de memória)
			//como parâmetro do método (usar &)
			err = fazerAcelerar(&carro)
		case 2:
			err = fazerFrear(&carro)

		}

		if err != nil {
			fmt.Printf(" ** Não foi possível executar a ação: %s ** \n", err)
		} else if opcao != 3 {
			fmt.Printf("O carro %s da marca %s está a %.2f km/h \n", carro.Modelo, carro.Marca, carro.Velocidade)
		}

		fmt.Println("Fim da execução...")

	}
}

//Para implementarmos uma interface a um método, precisamos
//criar uma função que retorna um erro, que recebe como parametro
//o nome do veículo seguido do nome da interface. Syntax:
//func <nome do método>(<nome do veiculo><interface>) error{}
func fazerAcelerar(veiculo estrutura.Acelerador) error {
	return veiculo.Acelerar()
}

//Obs.: Interfaces não ficam alocadas em memória, por isso
//não é necessário usar * antes de sua declaração como parâmetro

func fazerFrear(veiculo estrutura.Freio) error {
	return veiculo.Frear()
}
