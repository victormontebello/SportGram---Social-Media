package modelos

import (
	"errors"
	"modulo/src/seguranca"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa) ; erro != nil {
		return erro
	} 

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

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o e-mail inserido está em um formato inválido")
	}

	if usuario.Senha == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	if usuario.Esporte == "" {
		return errors.New("o esporte é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Esporte = strings.TrimSpace(usuario.Esporte)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)

		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}
	return nil
}
