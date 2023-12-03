package controllers

import (
	"net/http"
	"strconv"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"github.com/gin-gonic/gin"
)

func BuscarTodosFuncionarios(c *gin.Context) {
	var funcionarios []models.Funcionario
	database.DB.
		Preload("Cargo").
		Preload("Endereco").
		Find(&funcionarios)
	c.JSON(http.StatusOK, funcionarios)
}

func BuscarFuncionarioPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var funcionario models.Funcionario
	database.DB.
		Preload("Cargo").
		Preload("Endereco").
		First(&funcionario, id)

	if funcionario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Funcionário não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, funcionario)
}

func CriaFuncionario(c *gin.Context) {
	var funcionario models.Funcionario

	if err := c.ShouldBindJSON(&funcionario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}
	database.DB.Create(&funcionario)
	c.Params = gin.Params{
		gin.Param{Key: "id", Value: strconv.FormatUint(uint64(funcionario.ID), 10)},
	}
	BuscarFuncionarioPorId(c)
}

func AtualizarFuncionario(c *gin.Context) {
	var funcionario models.Funcionario

	id := c.Params.ByName("id")
	database.DB.First(&funcionario, id)

	if funcionario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Funcionário não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&funcionario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}

	database.DB.Model(&funcionario).UpdateColumns(funcionario)
	c.Params = gin.Params{
		gin.Param{Key: "id", Value: strconv.FormatUint(uint64(funcionario.ID), 10)},
	}
	BuscarFuncionarioPorId(c)
}

func ExcluirFuncionario(c *gin.Context) {
	id := c.Params.ByName("id")
	var funcionario models.Funcionario
	database.DB.First(&funcionario, id)

	if funcionario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Funcionário não encontrado",
		})
		return
	}

	database.DB.Delete(&funcionario, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"mesagem": "OK",
	})
}
