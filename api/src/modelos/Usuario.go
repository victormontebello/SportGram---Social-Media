package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID               uint64    `json:"id,omitempty"`
	Nome             string    `json:"nome,omitempty"`
	Nick             string    `json:"nick,omitempty"`
	Email            string    `json:"email,omitempty"`
	Senha            string    `json:"senha"`
	CriadoEm         time.Time `json:"criadoEm,omitempty"`
	Esporte          string    `json:"esporte,omitempty"`
	AnosExperiencia  int       `json:"anosExperiencia,omitempty"`
	PossuiPatrocinio bool      `json:"possuiPatrocinio,omitempty"`
}

func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()

	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	if usuario.Senha == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	if usuario.Esporte == "" {
		return errors.New("o esporte é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Esporte = strings.TrimSpace(usuario.Esporte)
}
