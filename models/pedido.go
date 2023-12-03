package models

import "gorm.io/gorm"

type Pedido struct {
	gorm.Model
	Valor             float64 `json:"valorTotal"`
	MetodoPagamentoID int
	MetodoPagamento   MetodoPagamento `json:"metodoPagamento"`
	FuncionarioID     int
	Funcionario       Funcionario `json:"feitoPor"`
	Pratos            []Prato     `gorm:"many2many:pedido_prato;" json:"pratos"`
}
