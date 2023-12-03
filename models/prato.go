package models

import "gorm.io/gorm"

type Prato struct {
	gorm.Model
	Nome         string        `json:"nome"`
	Descricao    string        `json:"descricao"`
	Valor        float64       `json:"valor"`
	Disponivel   bool          `json:"disponivel"`
	Ingredientes []Ingrediente `gorm:"many2many:prato_ingrediente;" json:"ingredientes"`
}
