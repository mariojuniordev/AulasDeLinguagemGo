package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	fmt.Println(config.StringConexaoBanco)

	fmt.Println("Rodando API!")

	//Gera um novo Router e encapsula numa variável
	r := router.Gerar()

	fmt.Println(config.SecretKey)

	//log.Fatal() loga o um Fatal Error
	//o método ListenAndServe() do objeto http inicializa um servidor HTTP e recebe como parâmetro a porta
	//onde será executada a requisição e um router ou nil
	//http.ListenAndServe("<endereço da porta>", <router>)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}