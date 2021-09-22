package main

import (
	"fmt"
	"strconv"
)

func main() {

	var texto string;
	fmt.Print("Digite um número\n");
	fmt.Scanf("%s", &texto);
	//var numero int;
	//numero, _ = strconv.Atoi(texto);

	//Syntax: strconv.ParseInt(x, <decimal(10), octa(8) ou hexa(16)>, <precisão em bits (32 ou 64)>)

	numero, _ := strconv.ParseInt(texto, 10, 32); //o 10 como segundo parâmetro é p/ converter p/ decimal
	
	
	var conv = float64(numero);
	fmt.Println(numero);
	fmt.Println(conv);
	
}