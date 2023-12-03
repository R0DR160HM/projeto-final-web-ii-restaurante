package controllers

import (
	"net/http"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"github.com/gin-gonic/gin"
)

func BuscarTodosMetodoPagamentos(c *gin.Context) {
	var metodoPagamentos []models.MetodoPagamento
	database.DB.Find(&metodoPagamentos)
	c.JSON(http.StatusOK, metodoPagamentos)
}

func BuscarMetodoPagamentoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var metodoPagamento models.MetodoPagamento
	database.DB.First(&metodoPagamento, id)

	if metodoPagamento.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Método de Pagamento não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, metodoPagamento)
}

func CriaMetodoPagamento(c *gin.Context) {
	var metodoPagamento models.MetodoPagamento

	if err := c.ShouldBindJSON(&metodoPagamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}
	database.DB.Create(&metodoPagamento)
	BuscarMetodoPagamentoPorId(c)
}

func AtualizarMetodoPagamento(c *gin.Context) {
	var metodoPagamento models.MetodoPagamento
	id := c.Params.ByName("id")
	database.DB.First(&metodoPagamento, id)

	if metodoPagamento.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Método de Pagamento não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&metodoPagamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}

	database.DB.Model(&metodoPagamento).UpdateColumns(metodoPagamento)
	BuscarMetodoPagamentoPorId(c)
}

func ExcluirMetodoPagamento(c *gin.Context) {
	id := c.Params.ByName("id")
	var metodoPagamento models.MetodoPagamento
	database.DB.First(&metodoPagamento, id)

	if metodoPagamento.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Método de Pagamento não encontrado",
		})
		return
	}

	database.DB.Delete(&metodoPagamento, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"mesagem": "OK",
	})
}
