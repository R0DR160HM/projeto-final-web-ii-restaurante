package models

import "gorm.io/gorm"

type Pedido struct {
	gorm.Model
	Valor           float64         `json:"valorTotal"`
	MetodoPagamento MetodoPagamento `json:"metodoPagamento"`
	FeitoPor        Funcionario     `json:"feitoPor"`
}
