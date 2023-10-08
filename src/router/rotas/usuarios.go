package rotas

import (
	"api/src/controllers"
	"net/http"
)

var RotasUsuario = []Rota{
	{
		URI:              "/usuarios",
		Metodo:           http.MethodPost,
		Funcao:           controllers.CriarUsuario,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuarios",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuarios,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioID}",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioID}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.AtualizarUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioID}",
		Metodo:           http.MethodDelete,
		Funcao:           controllers.DeletarUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioID}/seguir",
		Metodo:           http.MethodPost,
		Funcao:           controllers.SeguirUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioID}/parar-de-seguir",
		Metodo:           http.MethodPost,
		Funcao:           controllers.PararDeSeguirUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioID}/seguindo",
		Metodo:           http.MethodPost,
		Funcao:           controllers.BuscarSeguindo,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioID}/atualizar-senha",
		Metodo:           http.MethodPost,
		Funcao:           controllers.AtualizarSenha,
		RequerAutenticao: true,
	},
}
