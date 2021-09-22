package main

import (
	"container/list"
	"fmt"
)

func main() {
	numeros := list.New();
	//O método PushBack() insere um elemento na parte de trás (fim da lista)
	el := numeros.PushBack(1);
	//O método PushFront() insere um elemento na parte da frente da lista (início da lista)
	numeros.PushFront(0);
	numeros.PushBack(2);
	//Saída: 012

	//O equivalente de null em Go é nil
	//O método .Front() retorna o primeiro elemento da lista
	//O método .Next() pega o próximo elmento
	//O método .Prev() pega o elemento anterior
	//O método .Value pega o próprio elemento 
	for e := numeros.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value);
	}

	fmt.Println("---------------------------")

	//O método .Remove() remove elementos de uma lista
	numeros.Remove(el);
	for e := numeros.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value);
	}
}