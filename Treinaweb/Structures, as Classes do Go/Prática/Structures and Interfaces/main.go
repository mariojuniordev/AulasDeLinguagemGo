package main

import (
	"errors"
	"fmt"
)

//As estruturas equivalentes às Classes no Go são as structures (struct)

type Carro struct {
	modelo     string
	marca      string
	velocidade float32
}

//Para criar métodos relacionados a estruturas em Go é necessário
//criar as funções (métodos) correspondentes fora do escopo
//de bloco, passando como parâmetro o a struct que aciona o
//método a partir do nome do método mais o ponteiro para a
//struct Carro (*Carro), segue syntax:
//func(<nome da struct minusculo> *<nome da struct>) <nome do método>(){}

//Caso a função contenha erro, é necessário colocar error logo
//após o nome da função (método)

func (carro *Carro) acelerar() error {
	if carro.velocidade < 15 {
		carro.velocidade += 5
		return nil
	} else {
		//O pacote errors do Go deve receber apenas letras minúsculas
		return errors.New("o carro já está em sua velocidade máxima")
	}
}

func (carro *Carro) frear() error {
	if carro.velocidade > 0 {
		carro.velocidade -= 5
		return nil
	} else {
		return errors.New("o carro já está parado")
	}
}

func main() {
	//Em Go não é necessário usar classes construtoras
	//basta encapsular o instância da structure dentro
	//de uma variável
	carro := Carro{modelo: "Palio", marca: "Fiat"}
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
			err = carro.acelerar()
		case 2:
			err = carro.frear()

		}

		if err != nil {
			fmt.Printf(" ** Não foi possível executar a ação: %s ** \n", err)
		} else if opcao != 3 {
			fmt.Printf("O carro %s da marca %s está a %.2f km/h \n", carro.modelo, carro.marca, carro.velocidade)
		}

		fmt.Println("Fim da execução...")

	}
}
