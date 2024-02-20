package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"modulo/src/autenticacao"
	"modulo/src/banco"
	"modulo/src/modelos"
	"modulo/src/repositorios"
	"modulo/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

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
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.Buscar(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
	
}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request){
	var params = mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPorID(publicacaoID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

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

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, errors.New("publicação não encontrada"))
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicação que não seja sua"))
		return
	}

	publicacao.ID = publicacaoID

	erro = repositorio.Atualizar(publicacao)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	publicacaoID, erro := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacao.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar uma publicação que não seja sua"))
		return
	}

	erro = repositorio.Deletar(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func BuscarPublicacoesDoUsuario(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.BuscarPorUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if(publicacoes == nil){
		erro := errors.New("nenhuma publicação encontrada")
		respostas.Erro(w, http.StatusNotFound, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	publicacaoID, erro := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	_, erro = autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	erro = repositorio.Curtir(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}