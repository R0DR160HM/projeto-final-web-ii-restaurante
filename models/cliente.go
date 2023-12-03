package models

import (
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Nome           string   `json:"nome"`
	CPF            string   `json:"cpf"`
	DataNascimento string   `json:"dataNascimento"`
	Email          string   `json:"email"`
	Celular        string   `json:"celular"`
	EnderecoId     int      `json:"enderecoId"`
	Endereco       Endereco `json:"endereco"`
}
