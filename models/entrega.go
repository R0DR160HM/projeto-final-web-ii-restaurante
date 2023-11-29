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
	Pedido          Pedido      `json:"pedido"`
	Cliente         Cliente     `json:"cliente"`
	EntreguePor     Funcionario `json:"entreguePor"`
}
