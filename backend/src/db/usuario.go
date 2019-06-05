package db

import (
	"log"
	"models"
)

// Gravar efetua a gravação do usuário.
func Gravar(usuario *models.Usuario) {

	sql := `INSERT INTO USUARIO(USU_EMAIL, USU_NOME, USU_SENHA, USU_RECEBERALERTA)
	VALUES("$1", "$2", "$3", $4)
	RETURNING (USU_ID, USU_DATACRIACAO, "")`

	con := ObtenhaConexao()
	defer con.Close()

	rows := con.QueryRow(sql, usuario.Email, usuario.Nome, usuario.Senha, usuario.RecebeAlertas)
	err := rows.Scan(usuario.ID, usuario.DataCriacao, usuario.Senha)
	if err != nil {
		log.Fatal(err)
	}
}

// Remover efetua a remoção de um usuário.
func Remover(id string) {
	sql := `DELETE FROM USUARIO WHERE USU_ID = "$1"`

	con := ObtenhaConexao()
	defer con.Close()

	_, err := con.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}
