package db

import (
	"models"
	"time"
)

// Gravar efetua a gravação do usuário.
func Gravar(usuario *models.Usuario) error {

	sql := `INSERT INTO USUARIO(USU_EMAIL, USU_NOME, USU_SENHA, USU_RECEBERALERTA)
	VALUES($1, $2, crypt($3, gen_salt("bf")), $4)
	RETURNING (USU_ID, USU_DATACRIACAO)`

	con, err := ObtenhaConexao()
	if err != nil {
		return err
	}
	defer con.Close()

	var id string
	var dataCriacao time.Time
	rows := con.QueryRow(sql, usuario.Email, usuario.Nome, usuario.Senha, usuario.RecebeAlertas)
	err = rows.Scan(&id, &dataCriacao)
	if err != nil {
		return err
	}

	usuario.ID = id
	usuario.DataCriacao = dataCriacao
	usuario.Senha = ""

	return nil
}

// Remover efetua a remoção de um usuário.
func Remover(id string) error {
	sql := `DELETE FROM USUARIO WHERE USU_ID = $1`

	con, err := ObtenhaConexao()
	if err != nil {
		return err
	}
	defer con.Close()

	_, err = con.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}
