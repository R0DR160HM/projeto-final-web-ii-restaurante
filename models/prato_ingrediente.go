package models

type PratoIngrediente struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	Ingrediente Ingrediente `json:"ingrediente"`
	Prato       Prato       `json:"prato"`
}
