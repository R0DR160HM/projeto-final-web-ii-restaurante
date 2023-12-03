package controllers

import (
	"net/http"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"github.com/gin-gonic/gin"
)

func BuscarTodosCargos(c *gin.Context) {
	var cargos []models.Cargo
	database.DB.Find(&cargos)
	c.JSON(http.StatusOK, cargos)
}

func BuscarCargoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var cargo models.Cargo
	database.DB.First(&cargo, id)

	if cargo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Cargo não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, cargo)
}

func CriaCargo(c *gin.Context) {
	var cargo models.Cargo

	if err := c.ShouldBindJSON(&cargo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}
	database.DB.Create(&cargo)
	c.JSON(http.StatusOK, cargo)
}

func AtualizarCargo(c *gin.Context) {
	var cargo models.Cargo
	id := c.Params.ByName("id")
	database.DB.First(&cargo, id)

	if cargo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Cargo não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&cargo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}

	database.DB.Model(&cargo).UpdateColumns(cargo)
	c.JSON(http.StatusOK, cargo)
}

func ExcluirCargo(c *gin.Context) {
	id := c.Params.ByName("id")
	var cargo models.Cargo
	database.DB.First(&cargo, id)

	if cargo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Cargo não encontrado",
		})
		return
	}

	database.DB.Delete(&cargo, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"mesagem": "OK",
	})
}
