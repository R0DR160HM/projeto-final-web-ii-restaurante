package models

import "gorm.io/gorm"

type Ingrediente struct {
	gorm.Model
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
}
