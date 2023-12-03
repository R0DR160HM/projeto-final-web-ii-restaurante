package main

import (
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
