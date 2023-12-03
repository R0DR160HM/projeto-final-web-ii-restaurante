package controllers

import (
	"net/http"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"github.com/gin-gonic/gin"
)

func BuscarTodosClientes(c *gin.Context) {
	var clientes []models.Cliente
	database.DB.Find(&clientes)
	c.JSON(http.StatusOK, clientes)
}

func BuscarClientePorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var cliente models.Cliente
	database.DB.First(&cliente, id)

	if cliente.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Cliente não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, cliente)
}

func CriaCliente(c *gin.Context) {
	var cliente models.Cliente

	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}
	database.DB.Create(&cliente)
	c.JSON(http.StatusOK, cliente)
}

func AtualizarCliente(c *gin.Context) {
	var cliente models.Cliente
	id := c.Params.ByName("id")
	database.DB.First(&cliente, id)

	if cliente.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Cliente não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}

	database.DB.Model(&cliente).UpdateColumns(cliente)
	c.JSON(http.StatusOK, cliente)
}

func ExcluirCliente(c *gin.Context) {
	id := c.Params.ByName("id")
	var cliente models.Cliente
	database.DB.Delete(&cliente, id)

	if cliente.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Cliente não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"mesagem": "OK",
	})
}
