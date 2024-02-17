package rotas

import (
	"net/http"
)

// Rota : Representa todas as rotas da aplicação
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar : Inicializa as rotas da aplicação
