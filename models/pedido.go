package models

import "gorm.io/gorm"

type Pedido struct {
	gorm.Model
	Valor             float64         `json:"valorTotal"`
	MetodoPagamentoID int             `json:"metodoPagamentoId"`
	MetodoPagamento   MetodoPagamento `json:"metodoPagamento"`
	FuncionarioID     int             `json:"funcionarioId"`
	Funcionario       Funcionario     `json:"feitoPor"`
	Pratos            []Prato         `gorm:"many2many:pedido_prato;" json:"pratos"`
}
