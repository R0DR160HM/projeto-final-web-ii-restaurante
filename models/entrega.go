package models

import (
	"time"

	"gorm.io/gorm"
)

type Entrega struct {
	gorm.Model
	PrevisaoEntrega time.Time `json:"previsaoEntrega"`
	EntregueEm      time.Time `json:"entregueEm"`
	Entregue        bool      `json:"entregue"`
	PedidoID        int
	Pedido          Pedido `json:"pedido"`
	ClienteID       int
	Cliente         Cliente `json:"cliente"`
	FuncionarioID   int
	Funcionario     Funcionario `json:"entreguePor"`
}
