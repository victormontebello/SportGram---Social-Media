package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao : representa uma publicação feita por um usuário
type Publicacao struct {
	ID           uint64 `json:"id,omitempty"`
	Titulo       string `json:"titulo,omitempty"`
	Conteudo     string `json:"conteudo,omitempty"`
	AutorID      uint64 `json:"autorId,omitempty"`
	AutorNick    string `json:"autorNick,omitempty"`
	Curtidas     uint64 `json:"curtidas"`
	CriadaEm     time.Time `json:"criadaEm,omitempty"`
	Midia 	   []byte `json:"midia,omitempty"`
}

// Preparar : prepara a publicação para ser salva no banco de dados
func (p *Publicacao) Preparar() error {
	if erro := p.validar(); erro != nil {
		return erro
	}

	p.formatar()

	return nil
}

func (p *Publicacao) validar() error {
	if p.Titulo == "" {
		return errors.New("o título é obrigatório e não pode estar em branco")
	}

	if p.Conteudo == "" {
		return errors.New("o conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (p *Publicacao) formatar() {
	p.Titulo = strings.TrimSpace(p.Titulo)
	p.Conteudo = strings.TrimSpace(p.Conteudo)

	p.CriadaEm = time.Now()

	if p.Curtidas == 0 {
		p.Curtidas = 0
	}

	if p.Midia == nil {
		p.Midia = []byte{}
	}
}


