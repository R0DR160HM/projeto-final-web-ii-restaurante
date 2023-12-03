package models

import (
	"time"

	"gorm.io/gorm"
)

type Entrega struct {
	gorm.Model
	PrevisaoEntrega time.Time   `json:"previsaoEntrega"`
	EntregueEm      time.Time   `json:"entregueEm"`
	Entregue        bool        `json:"entregue"`
	PedidoID        int         `json:"pedidoId"`
	Pedido          Pedido      `json:"pedido"`
	ClienteID       int         `json:"clienteId"`
	Cliente         Cliente     `json:"cliente"`
	FuncionarioID   int         `json:"funcionarioId"`
	Funcionario     Funcionario `json:"entreguePor"`
}
