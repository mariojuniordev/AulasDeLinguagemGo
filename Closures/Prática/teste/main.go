package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estou abrindo a conexão com o banco de dados");
	//o defer, assim como o finally só éxecutado depois que todas as linhas de código da função são executadas

	executarConsulta();
}

func executarConsulta(){
	fmt.Println("Estou executando a consulta ao banco de dados...");
	defer fmt.Println("Estou fechando a conexão com o banco de dados!");
	//Se houver mais de 1 defer, o defer a ser executado primeiro é o que for o último difer na ordem das linhas
	//Seguindo o princípio LIFO (Last In Frist Out)
	defer fmt.Println("Mais um Defer");
}