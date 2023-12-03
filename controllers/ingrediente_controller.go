package controllers

import (
	"net/http"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"github.com/gin-gonic/gin"
)

func BuscarTodosIngredientes(c *gin.Context) {
	var ingredientes []models.Ingrediente
	database.DB.Find(&ingredientes)
	c.JSON(http.StatusOK, ingredientes)
}

func BuscarIngredientePorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var ingrediente models.Ingrediente
	database.DB.First(&ingrediente, id)

	if ingrediente.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Ingrediente não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, ingrediente)
}

func CriaIngrediente(c *gin.Context) {
	var ingrediente models.Ingrediente

	if err := c.ShouldBindJSON(&ingrediente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}
	database.DB.Create(&ingrediente)
	c.JSON(http.StatusOK, ingrediente)
}

func AtualizarIngrediente(c *gin.Context) {
	var ingrediente models.Ingrediente
	id := c.Params.ByName("id")
	database.DB.First(&ingrediente, id)

	if ingrediente.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Ingrediente não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&ingrediente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}

	database.DB.Model(&ingrediente).UpdateColumns(ingrediente)
	c.JSON(http.StatusOK, ingrediente)
}

func ExcluirIngrediente(c *gin.Context) {
	id := c.Params.ByName("id")
	var ingrediente models.Ingrediente
	database.DB.First(&ingrediente, id)

	if ingrediente.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Ingrediente não encontrado",
		})
		return
	}

	database.DB.Delete(&ingrediente, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"mesagem": "OK",
	})
}
