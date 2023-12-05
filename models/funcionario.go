package models

import (
	"gorm.io/gorm"
)

type Funcionario struct {
	gorm.Model
	Nome           string   `json:"nome"`
	CPF            string   `json:"cpf" gorm:"size:14; unique"`
	DataNascimento string   `json:"dataNascimento" gorm:"size:10"`
	CargoID        int      `json:"cargoId"`
	Cargo          Cargo    `json:"cargo"`
	EnderecoID     int      `json:"enderecoId"`
	Endereco       Endereco `json:"endereco"`
}
