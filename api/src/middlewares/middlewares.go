package middlewares

import "net/http"

// Autenticar verifica a autenticação da requisição
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//autenticando...
		next(w, r)
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//logando...
		next(w, r)
	}
}