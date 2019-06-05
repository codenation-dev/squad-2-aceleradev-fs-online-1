package models

import (
	"time"
)

// Usuario representa o cadastro de um usuário
type Usuario struct {
	ID            string
	Email         string
	Nome          string
	Senha         string
	DataCriacao   time.Time
	RecebeAlertas bool
}
