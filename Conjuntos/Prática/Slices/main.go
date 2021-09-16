package main

import (
	"fmt"
)

func main() {
	//O Slice é inicializado com um determinado número de posições, 
	//mas pode ser expandido de forma dinâmica
	//Syntax: <nome da variavel> := make([]<tipo dos elementos>, <numero de posições inicial>)
	amigos := make([]string, 3);
	nome := "";
	i := 0;

	//loop for com estrutura equivalente ao while
	for nome != "q" {
		fmt.Print("Digite o nome de um amigo ou 'q' para parar: ")
		fmt.Scanf("%s", &nome);
		if nome != "q" {
			if i < 3 {
				amigos[i] = nome;
			} else {
				//a função nativa do Go, append(), expande o tamanho do Slice
				//Syntax: apprend(<numoe do Slice>, <nome dos elementos do Slice>)
				amigos = append(amigos, nome);
			}
			i++;
		}
	}
	fmt.Println(amigos);
	fmt.Printf("Você tem %d amigos \n", len(amigos));
	for _, nomeAmigo := range amigos {
		fmt.Printf(" - %s \n", nomeAmigo);
	}

	//Para acessar elementos específcios dentro de um slice 
	//podemos usar a syntax a seguir p/ acessar estes dados
	//<nome do novo Slice> := <slice desejado>[<posição inicial |inclusa|> : <posição final |não-inclusa|>]
	outroSlice := amigos[1:3];
	fmt.Println(outroSlice);

	//Ao omitir o índice do primeiro elemento de um slice
	//, o índice será automaticamente subentendido como 0
	maisUmSlice := amigos[:3];
	fmt.Println(maisUmSlice);
}