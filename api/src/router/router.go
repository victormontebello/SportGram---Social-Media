package router

import (
	"modulo/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar : Retorna um roteador com as rotas configuradas
func Gerar() *mux.Router{
	r := mux.NewRouter()
	return rotas.Configurar(r)
}