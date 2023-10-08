package rotas

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

//Rota representa todas as rotas da API
type Rota struct {
	URI              string
	Metodo           string
	Funcao           func(w http.ResponseWriter, r *http.Request)
	RequerAutenticao bool
}

//Configurar configura as rotas
func Configurar(r *mux.Router) *mux.Router {

	rotas := RotasUsuario
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {
		if rota.RequerAutenticao == true {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	return r
}
