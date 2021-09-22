package main

import (
	"fmt"
)

var n int

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}

//A função init() é uma função q é executada antes da função main()
func init() {
	n = 10
	fmt.Println("Função init sendo executada")
}
