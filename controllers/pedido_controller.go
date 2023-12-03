package controllers

import (
	"net/http"
	"strconv"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"github.com/gin-gonic/gin"
)

func BuscarTodosPedidos(c *gin.Context) {
	var pedidos []models.Pedido
	database.DB.
		Preload("MetodoPagamento").
		Preload("Funcionario").
		Preload("Pratos").
		Find(&pedidos)
	c.JSON(http.StatusOK, pedidos)
}

func BuscarPedidoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var pedido models.Pedido
	database.DB.
		Preload("MetodoPagamento").
		Preload("Funcionario").
		Preload("Pratos").
		First(&pedido, id)

	if pedido.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Pedido não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, pedido)
}

func CriaPedido(c *gin.Context) {
	var pedido models.Pedido

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}
	database.DB.Create(&pedido)
	retornarPedido(c, pedido)
}

func AtualizarPedido(c *gin.Context) {
	var pedido models.Pedido
	id := c.Params.ByName("id")
	database.DB.First(&pedido, id)

	if pedido.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Pedido não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}

	database.DB.Model(&pedido).UpdateColumns(pedido)
	retornarPedido(c, pedido)
}

func ExcluirPedido(c *gin.Context) {
	id := c.Params.ByName("id")
	var pedido models.Pedido
	database.DB.First(&pedido, id)

	if pedido.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Pedido não encontrado",
		})
		return
	}

	database.DB.Delete(&pedido, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"mesagem": "OK",
	})
}

func retornarPedido(c *gin.Context, m models.Pedido) {
	c.Params = gin.Params{
		gin.Param{Key: "id", Value: strconv.FormatUint(uint64(m.ID), 10)},
	}
	BuscarPedidoPorId(c)
}
