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
	CargoID        int
	Cargo          Cargo `json:"cargo"`
	EnderecoID     int
	Endereco       Endereco `json:"endereco"`
}
