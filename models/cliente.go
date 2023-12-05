package models

import (
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Nome           string   `json:"nome"`
	CPF            string   `json:"cpf" gorm:"size:14; unique"`
	DataNascimento string   `json:"dataNascimento" gorm:"size:10"`
	Email          string   `json:"email" gorm:"unique"`
	Celular        string   `json:"celular" gorm:"unique"`
	EnderecoId     int      `json:"enderecoId"`
	Endereco       Endereco `json:"endereco"`
}
