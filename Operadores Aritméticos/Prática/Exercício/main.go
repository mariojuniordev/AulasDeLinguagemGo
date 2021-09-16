package main

import (
	"fmt"
	"strconv"
)

func main() {
	inteiro, _ := strconv.ParseInt("2", 10, 32);
	fmt.Println(inteiro);
	fmt.Println(11/2); //por padrão, retorna valores inteiros, portanto, retornará 5

}