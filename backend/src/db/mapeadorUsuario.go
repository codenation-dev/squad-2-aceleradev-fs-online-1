package db

import (
	"bytes"
	"models"
	"time"
)

var sqlSelect = "SELECT USU_CODIGO, USU_EMAIL, USU_NOME, USU_SENHA, USU_DATACRIACAO, USU_RECEBERALERTA FROM USUARIO"

// GravarUsuario efetua a gravação do usuário.
func GravarUsuario(usuario *models.Usuario) error {

	sql := `INSERT INTO USUARIO(USU_EMAIL, USU_NOME, USU_SENHA, USU_RECEBERALERTA)
	VALUES($1, $2, crypt($3, gen_salt("bf")), $4)
	RETURNING (USU_CODIGO, USU_DATACRIACAO)`

	con, err := ObtenhaConexao()
	if err != nil {
		return err
	}
	defer con.Close()

	var codigo int
	var dataCriacao time.Time
	rows := con.QueryRow(sql, usuario.Email, usuario.Nome, usuario.Senha, usuario.RecebeAlertas)
	err = rows.Scan(&codigo, &dataCriacao)
	if err != nil {
		return err
	}

	usuario.Codigo = codigo
	usuario.DataCriacao = dataCriacao
	usuario.Senha = ""

	return nil
}

// AlterarUsuario efetua a alteração do usuário.
func AlterarUsuario(usuario *models.Usuario) error {

	sql := `
	UPDATE USUARIO
	SET USU_EMAIL = $2,
		USU_NOME = $3,
		USU_SENHA = crypt($4, gen_salt("bf")),
		USU_RECEBERALERTA = $5
	WHERE USU_CODIGO = $1`

	con, err := ObtenhaConexao()
	if err != nil {
		return err
	}
	defer con.Close()

	_, err = con.Exec(sql, usuario.Codigo, usuario.Email, usuario.Nome, usuario.Senha, usuario.RecebeAlertas)
	if err != nil {
		return err
	}

	usuario.Senha = ""

	return nil
}

// RemoverUsuario efetua a remoção de um usuário.
func RemoverUsuario(codigo int) error {

	sql := `DELETE FROM USUARIO WHERE USU_CODIGO = $1`

	con, err := ObtenhaConexao()
	if err != nil {
		return err
	}
	defer con.Close()

	_, err = con.Exec(sql, codigo)
	if err != nil {
		return err
	}

	return nil
}

// ObtenhaUsuario obtém um usuário pelo código
func ObtenhaUsuario(codigo int) (models.Usuario, error) {

	usuario := models.Usuario{}

	con, err := ObtenhaConexao()
	if err != nil {
		return usuario, err
	}
	defer con.Close()

	sql := bytes.NewBufferString(sqlSelect)
	sql.WriteString(" WHERE USU_CODIGO = $1")

	rows := con.QueryRow(sql.String(), codigo)
	err = rows.Scan(&usuario.Codigo, &usuario.Nome, &usuario.Senha, &usuario.RecebeAlertas)
	if err != nil {
		return usuario, err
	}

	usuario.Senha = ""

	return usuario, nil
}

// ObtenhaListaUsuario obtém todos os usuários.
func ObtenhaListaUsuario() ([]models.Usuario, error) {

	var usuarios []models.Usuario

	con, err := ObtenhaConexao()
	if err != nil {
		return usuarios, err
	}
	defer con.Close()

	rows, err := con.Query(sqlSelect)
	if err != nil {
		return usuarios, err
	}

	defer rows.Close()
	for rows.Next() {
		var usuAux = models.Usuario{}
		rows.Scan(&usuAux.Codigo, &usuAux.Email, &usuAux.Nome, &usuAux.Senha, &usuAux.DataCriacao, &usuAux.RecebeAlertas)
		usuAux.Senha = ""
		usuarios = append(usuarios, usuAux)
	}

	return usuarios, nil
}

// SenhaEstaCorreta verifica se a senha informada para o e-mail está correta.
func SenhaEstaCorreta(email string, senha string) (bool, error) {

	con, err := ObtenhaConexao()
	if err != nil {
		return false, err
	}
	defer con.Close()

	sql := "SELECT COUNT(*) FROM USUARIO WHERE USU_EMAIL = $1 AND USU_SENHA = crypt($2, USU_SENHA)"

	row := con.QueryRow(sql, email, senha)
	var quantidade int
	err = row.Scan(&quantidade)
	if err != nil {
		return false, err
	}

	return quantidade > 0, nil
}
