package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Nome           string    `json:"nome"`
	CPF            string    `json:"cpf"`
	DataNascimento time.Time `json:"dataNascimento"`
	Email          string    `json:"email"`
	Celular        string    `json:"celular"`
	EnderecoId     int
	Endereco       Endereco `json:"endereco"`
}
