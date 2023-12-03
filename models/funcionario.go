package models

import (
	"gorm.io/gorm"
)

type Funcionario struct {
	gorm.Model
	Nome           string   `json:"nome"`
	CPF            string   `json:"cpf"`
	DataNascimento string   `json:"dataNascimento"`
	CargoID        int      `json:"cargoId"`
	Cargo          Cargo    `json:"cargo"`
	EnderecoID     int      `json:"enderecoId"`
	Endereco       Endereco `json:"endereco"`
}
