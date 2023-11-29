package models

import "gorm.io/gorm"

type Cargo struct {
	gorm.Model
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}
