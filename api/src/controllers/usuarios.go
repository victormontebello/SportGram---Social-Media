package controllers

import (
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criar usuário"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários"))
}

func BuscarUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário por id"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar usuário"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar usuário"))
}
