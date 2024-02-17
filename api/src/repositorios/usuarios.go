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

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, senha, criadoEm, esporte, anosExperiencia, possuiPatrocinio from usuarios where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick,
	)
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
	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm, esporte, anosExperiencia, possuiPatrocinio from usuarios where id = ?", ID)
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

	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ?, esporte = ?, anosExperiencia = ?, possuiPatrocinio = ? where id = ?",
	)
	if erro != nil {
		fmt.Println("Erro ao preparar a query de atualização de usuário: ", erro)
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Esporte, usuario.AnosExperiencia, usuario.PossuiPatrocinio, usuario.ID); erro != nil {
		fmt.Println("Erro ao executar a query de atualização2 de usuário: ", erro)
		return erro
	}

	return nil
}