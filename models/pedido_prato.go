package models

type PedidoPrato struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Pedido Pedido `json:"pedido"`
	Prato  Prato  `json:"prato"`
}
