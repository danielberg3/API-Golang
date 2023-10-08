package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:              "/publicacoes",
		Metodo:           http.MethodPost,
		Funcao:           controllers.CriarPublicacao,
		RequerAutenticao: true,
	},
	{
		URI:              "/publicacoes",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarPublicacoes,
		RequerAutenticao: true,
	},
	{
		URI:              "/publicacoes/{publicacaoID}",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarPublicacao,
		RequerAutenticao: true,
	},
	{
		URI:              "/publicacoes/{publicacaoID}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.AtualizarPublicacao,
		RequerAutenticao: true,
	},
	{
		URI:              "/publicacoes/{publicacaoID}",
		Metodo:           http.MethodDelete,
		Funcao:           controllers.DeletarPublicacao,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioID}/publicacoes",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarPublicacoesPorUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/publicacoes/{publicacaoID}/curtir",
		Metodo:           http.MethodPost,
		Funcao:           controllers.CurtirPublicacao,
		RequerAutenticao: true,
	},
	{
		URI:              "/publicacoes/{publicacaoID}/descurtir",
		Metodo:           http.MethodPost,
		Funcao:           controllers.DescurtirPublicacao,
		RequerAutenticao: true,
	},
}
