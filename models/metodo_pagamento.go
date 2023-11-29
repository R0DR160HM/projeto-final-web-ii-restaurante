package models

import "gorm.io/gorm"

type MetodoPagamento struct {
	gorm.Model
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}
