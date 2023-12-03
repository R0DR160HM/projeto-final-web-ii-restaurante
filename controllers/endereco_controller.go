package controllers

import (
	"net/http"
	"strconv"

	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/database"
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/models"
	"github.com/gin-gonic/gin"
)

func BuscarTodosEnderecos(c *gin.Context) {
	var enderecos []models.Endereco
	database.DB.Find(&enderecos)
	c.JSON(http.StatusOK, enderecos)
}

func BuscarEnderecoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var endereco models.Endereco
	database.DB.First(&endereco, id)

	if endereco.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Endereço não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, endereco)
}

func CriaEndereco(c *gin.Context) {
	var endereco models.Endereco

	if err := c.ShouldBindJSON(&endereco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}
	database.DB.Create(&endereco)
	c.Params = gin.Params{
		gin.Param{Key: "id", Value: strconv.FormatUint(uint64(endereco.ID), 10)},
	}
	BuscarEnderecoPorId(c)
}

func AtualizarEndereco(c *gin.Context) {
	var endereco models.Endereco
	id := c.Params.ByName("id")
	database.DB.First(&endereco, id)

	if endereco.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Endereço não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&endereco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "ERRO",
			"mensagem": err.Error(),
		})
		return
	}

	database.DB.Model(&endereco).UpdateColumns(endereco)
	c.Params = gin.Params{
		gin.Param{Key: "id", Value: strconv.FormatUint(uint64(endereco.ID), 10)},
	}
	BuscarEnderecoPorId(c)
}

func ExcluirEndereco(c *gin.Context) {
	id := c.Params.ByName("id")
	var endereco models.Endereco
	database.DB.First(&endereco, id)

	if endereco.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "ERRO",
			"mensagem": "Endereço não encontrado",
		})
		return
	}

	database.DB.Delete(&endereco, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"mesagem": "OK",
	})
}
