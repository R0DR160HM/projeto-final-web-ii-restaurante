package controllers

import (
	"net/http"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"github.com/gin-gonic/gin"
)

func BuscarTodasEntregas(c *gin.Context) {
	var entregas []models.Entrega
	database.DB.
		Preload("Pedido").
		Preload("Cliente").
		Preload("Funcionario").
		Find(&entregas)
	c.JSON(http.StatusOK, entregas)
}

func BuscarEntregaPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var entrega models.Entrega
	database.DB.
		Preload("Pedido").
		Preload("Cliente").
		Preload("Funcionario").
		First(&entrega, id)

	if entrega.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Entrega não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, entrega)
}

func CriaEntrega(c *gin.Context) {
	var entrega models.Entrega

	if err := c.ShouldBindJSON(&entrega); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}
	database.DB.Create(&entrega)
	BuscarEntregaPorId(c)
}

func AtualizarEntrega(c *gin.Context) {
	var entrega models.Entrega
	id := c.Params.ByName("id")
	database.DB.First(&entrega, id)

	if entrega.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Entrega não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&entrega); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}

	database.DB.Model(&entrega).UpdateColumns(entrega)
	BuscarEntregaPorId(c)
}

func ExcluirEntrega(c *gin.Context) {
	id := c.Params.ByName("id")
	var entrega models.Entrega
	database.DB.First(&entrega, id)

	if entrega.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Entrega não encontrado",
		})
		return
	}

	database.DB.Delete(&entrega, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"mesagem": "OK",
	})
}
