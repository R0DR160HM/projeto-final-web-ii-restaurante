package models

import (
	"gorm.io/gorm"
)

type Entrega struct {
	gorm.Model
	PrevisaoEntrega string      `json:"previsaoEntrega"`
	Entregue        bool        `json:"entregue" gorm:"default:false"`
	PedidoID        int         `json:"pedidoId"`
	Pedido          Pedido      `json:"pedido"`
	ClienteID       int         `json:"clienteId"`
	Cliente         Cliente     `json:"cliente"`
	FuncionarioID   int         `json:"funcionarioId"`
	Funcionario     Funcionario `json:"funcionario"`
}
