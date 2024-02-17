package rotas

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Rota : Representa todas as rotas da aplicação
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar : Inicializa as rotas da aplicação
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	} 
	return r
}
