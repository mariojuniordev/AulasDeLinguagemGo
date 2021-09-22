package main

import (
	"fmt"
)

func main(){
	/*var amigos [5]string;
	fmt.Println(amigos);
	for i := 0; i < len(amigos); i++ {
		fmt.Print("Digite o nome de um dos seus amigos: \n");
		fmt.Scanf("%s", &amigos[i]);
	}
	fmt.Println(amigos);
	fmt.Println("Seus amigos são: ");
	//Equivalente do forEach()
	//forEach amigo in amigos
	for _, amigo := range amigos {
		fmt.Printf(" - %s \n", amigo)
;	}*/

	//<nome do array> := [<numero de posições>]<tipo dos dados>{x[0], x[1]..., x[n]}
	arrayInicializado := [3]int{2, 4, 6};

	//array multidimensional
	fmt.Println(arrayInicializado);
	var matriz [3][3]int;
	for i:= 0; i < len(matriz); i++ {
		for j:=0; j < len(matriz); j++ {
			fmt.Print("Digite um número: ");
			fmt.Scanf("%d", &matriz[i][j]);
		}
	}
	fmt.Println(matriz);
}