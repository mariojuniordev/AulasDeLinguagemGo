package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:    "/usuarios", // URI são como "pedaços" de URLS
		Metodo: http.MethodPost, //o método .MethodPost envia dados de um cliente HTTP para o servidor
		//O método .ResponseWriter é usado para configurar uma reposta HTTP e enviar os dados AO CLIENTE
		//O método .Request é a FORMA COM QUE plataformas de comunicação com a internet (BROWSERS) PEDEM 
		//AO SERVIDOR INFORMAÇÕES NECESSÁRIAS PARA CARREGAR UMA PÁGINA WEB
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		//GET sem userId
		URI:    "/usuarios", 
		Metodo: http.MethodGet, //o método .MethodGet BUSCA dados de um cliente HTTP no servidor
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		//GET com userId
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodGet, 
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		//PUT
		URI:    "/usuarios/{usuarioId}", 
		Metodo: http.MethodPut, //O método .MethodPut faz requisição para armazenar a entidade-corpo num local
		//espicificado pela URL
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		//DELETE
		URI:    "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete, //O método .MethodDelete é usado para DELETAR um recurso do servidor
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuario/{usuarioId}/seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuario/{usuarioId}/parar-de-seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/seguidores",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/seguindo",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/atualizar-senha",
		Metodo: http.MethodPost,
		Funcao: controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}