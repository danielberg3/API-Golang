package banco

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//Conectar retorna uma conexão com o banco de dados
func Conectar() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.StringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
