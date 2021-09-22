package main

import (
	"fmt"
)

func main() {

	//ARRAYS
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "3", "4", "5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//SLICES

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	//Adicionar valores a um slice
	//<nome da var> = append(<nome do slice>, <elemento adicionado>)
	slice = append(slice, 18)

	//Retornar valores especificos de um slice
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	//Arrays Internos
	//<nome> := make([]<tipo>, <tamanho>, <capacidade máxima>)
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length (tamanho)
	fmt.Println(cap(slice3)) //capicity (capacidade)

	//Quando ultrapassamos a capacidade máxima de um slice em Go
	//o Go cria um novo slice com a capacidade (+1) dobrada
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
