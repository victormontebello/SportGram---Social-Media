package repositorios

import (
	"database/sql"
	"modulo/src/modelos"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	stmt, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id, midia) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer stmt.Close()

	resultado, erro := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID, publicacao.Midia)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}