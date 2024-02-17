package router

import (
	"github.com/gorilla/mux"
)

// Gerar : Retorna um roteador com as rotas configuradas
func Gerar() *mux.Router{
	return mux.NewRouter()
}