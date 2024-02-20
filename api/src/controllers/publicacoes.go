package controllers

import (
	"encoding/json"
	"io/ioutil"
	"modulo/src/autenticacao"
	"modulo/src/banco"
	"modulo/src/modelos"
	"modulo/src/repositorios"
	"modulo/src/respostas"
	"net/http"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request){
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	erro = json.Unmarshal(corpoDaRequisicao, &publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request){

}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request){

}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request){

}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request){

}
