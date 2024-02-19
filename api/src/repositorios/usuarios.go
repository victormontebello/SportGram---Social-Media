package repositorios

import (
	"database/sql"
	"fmt"
	"modulo/src/modelos"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha, esporte, anosExperiencia, possuiPatrocinio) values (?, ?, ?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha, usuario.Esporte, usuario.AnosExperiencia, usuario.PossuiPatrocinio)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuários que atendem um filtro
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	const query = "select id, nome, nick, email, senha, criadoEm, esporte, anosExperiencia, possuiPatrocinio from usuarios where nome like ? or nick like ?"

	linhas, erro := repositorio.db.Query(query,nomeOuNick, nomeOuNick,)
	if erro != nil {
		return nil, erro
	}
	
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Senha,
			&usuario.CriadoEm, 
			&usuario.Esporte, 
			&usuario.AnosExperiencia, 
			&usuario.PossuiPatrocinio,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz um usuário do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	const semID = 0
	const query = "select id, nome, nick, email, criadoEm, esporte, anosExperiencia, possuiPatrocinio from usuarios where id = ?"
	linhas, erro := repositorio.db.Query(query, ID)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm, 
			&usuario.Esporte, 
			&usuario.AnosExperiencia, 
			&usuario.PossuiPatrocinio,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	if(usuario.ID == semID) {
		return modelos.Usuario{}, fmt.Errorf("nenhum usuário encontrado com o id %d", ID)
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(usuario modelos.Usuario) error {
	const query = "update usuarios set nome = ?, nick = ?, email = ?, esporte = ?, anosExperiencia = ?, possuiPatrocinio = ? where id = ?"
	statement, erro := repositorio.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Esporte, usuario.AnosExperiencia, usuario.PossuiPatrocinio, usuario.ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar remove um usuário do banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	const query = "delete from usuarios where id = ?"
	statement, erro := repositorio.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	const query = "select id, senha from usuarios where email = ?"
	linhas, erro := repositorio.db.Query(query, email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	const query = "insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)"
	statement, erro := repositorio.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	const query = "delete from seguidores where usuario_id = ? and seguidor_id = ?"
	statement, erro := repositorio.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}