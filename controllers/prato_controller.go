package controllers

import (
	"net/http"
	"strconv"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"github.com/gin-gonic/gin"
)

func BuscarTodosPratos(c *gin.Context) {
	var pratos []models.Prato
	database.DB.Preload("Ingredientes").Find(&pratos)
	c.JSON(http.StatusOK, pratos)
}

func BuscarPratoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var prato models.Prato
	database.DB.Preload("Ingredientes").First(&prato, id)

	if prato.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Prato não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, prato)
}

func CriaPrato(c *gin.Context) {
	var prato models.Prato

	if err := c.ShouldBindJSON(&prato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}
	database.DB.Create(&prato)
	c.Params = gin.Params{
		gin.Param{Key: "id", Value: strconv.FormatUint(uint64(prato.ID), 10)},
	}
	BuscarPratoPorId(c)
}

func AtualizarPrato(c *gin.Context) {
	var prato models.Prato
	id := c.Params.ByName("id")
	database.DB.First(&prato, id)

	if prato.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Prato não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&prato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}

	database.DB.Model(&prato).UpdateColumns(prato)
	c.Params = gin.Params{
		gin.Param{Key: "id", Value: strconv.FormatUint(uint64(prato.ID), 10)},
	}
	BuscarPratoPorId(c)
}

func ExcluirPrato(c *gin.Context) {
	id := c.Params.ByName("id")
	var prato models.Prato
	database.DB.First(&prato, id)

	if prato.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Prato não encontrado",
		})
		return
	}

	database.DB.Delete(&prato, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"mesagem": "OK",
	})
}
