package models

import "gorm.io/gorm"

type Endereco struct {
	gorm.Model
	CEP    string `json:"cep" gorm:"size:9"`
	Rua    string `json:"rua"`
	Numero int64  `json:"numero"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`
	UF     string `json:"uf" gorm:"size:2"`
}
