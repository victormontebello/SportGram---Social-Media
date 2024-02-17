package controllers

import "net/http"

// Login : Autentica o usuário
func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}