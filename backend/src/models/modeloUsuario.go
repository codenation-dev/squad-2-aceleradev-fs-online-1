package models

import (
	"time"
)

// Usuario representa o cadastro de um usuário
type Usuario struct {
	Codigo        int       `json:"codigo"`
	Email         string    `json:"email"`
	Nome          string    `json:"nome"`
	Senha         string    `json:"senha"`
	DataCriacao   time.Time `json:"dataCriacao"`
	RecebeAlertas bool      `json:"recebeAlertas"`
}
