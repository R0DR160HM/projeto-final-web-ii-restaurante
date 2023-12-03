package database

import (
	"log"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectaer com o banco de dados")
	}
	DB.AutoMigrate(
		&models.Endereco{},
		&models.Cargo{},
		&models.Cliente{},
		&models.Funcionario{},
		&models.Ingrediente{},
		&models.Prato{},
		&models.MetodoPagamento{},
		&models.Pedido{},
		&models.Entrega{},
	)
}
