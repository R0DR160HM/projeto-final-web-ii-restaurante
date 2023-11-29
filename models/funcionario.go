package models

import (
	"time"

	"gorm.io/gorm"
)

type Funcionario struct {
	gorm.Model
	Nome           string    `json:"nome"`
	CPF            string    `json:"cpf"`
	DataNascimento time.Time `json:"dataNascimento"`
	Cargo          Cargo     `json:"cargo"`
	Endereco       Endereco  `json:"endereco"`
}
