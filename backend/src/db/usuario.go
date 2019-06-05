package db

import (
	"models"
)

// Gravar efetua a gravação do usuário.
func Gravar(usuario *models.Usuario) error {

	sql := `INSERT INTO USUARIO(USU_EMAIL, USU_NOME, USU_SENHA, USU_RECEBERALERTA)
	VALUES($1, $2, $3, $4)
	RETURNING (USU_ID, USU_DATACRIACAO, "")`

	con, err := ObtenhaConexao()
	if err != nil {
		return err
	}
	defer con.Close()

	rows := con.QueryRow(sql, usuario.Email, usuario.Nome, usuario.Senha, usuario.RecebeAlertas)
	err = rows.Scan(&usuario.ID, &usuario.DataCriacao, &usuario.Senha)
	if err != nil {
		return err
	}

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
