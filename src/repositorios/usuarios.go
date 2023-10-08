package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

//Usuarios é classe de usuário do repositório
type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuarios cria um repositório do tipo conexão com o banco
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//Criar permite criar um usuário
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios(nome, nick, email, senha) values (?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

//BUscar  encontra usuários que contenham as nome o nick semelhante ao valor digitado
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {

		var usuario modelos.Usuario

		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

//BuscarPorID encontra o usuário dono do ID informado
func (repositorio Usuarios) BuscarPorID(usuarioID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"Select id, nome, nick, email, criadoEm from usuarios where id = ?", usuarioID)

	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()
	var usuario modelos.Usuario

	if linhas.Next() {
		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}

//Atualizar permite atualizar os dados de um usuário
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {

	stament, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer stament.Close()

	if _, erro := stament.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

//Deletar exclui um usuário por ID
func (repositorio Usuarios) Deletar(ID uint64) error {

	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil

}

//BuscarPorEmail busca o email de um usuário
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro := linhas.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

//Seguir permite que um usuário siga outro
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("insert into seguidores(usuario_id, seguidor_id) values(?,?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

//PararDeSeguir permite o usuário deixar de seguir outro usuário
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {

	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

//BuscarSeguindo verifica que o usuário está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(`	
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

//BuscarSenha serve para buscar a senha de um usuário no banco
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linhas, erro := repositorio.db.Query("select  senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro := linhas.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil
}

//AtualizarSenha serve para atualizar a senha de um usuário
func (repostorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	stament, erro := repostorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	if _, erro := stament.Exec(senha, usuarioID); erro != nil {
		return erro
	}
	return nil
}
