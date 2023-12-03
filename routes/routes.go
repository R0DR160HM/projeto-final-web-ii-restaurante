package routes

import (
	"github.com/R0DR160HM/projeto-final-web-ii-restaurante/controllers"
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

func HandleRequest() {
	r = gin.Default()

	handleEndereco()
	handleCargo()
	handleCliente()
	handleEntrega()
	handleFuncionario()
	handleIngrediente()
	handleMetodoPagamento()
	handlePedido()
	handlePrato()

	r.Run()
}

func handleEndereco() {
	r.GET("/enderecos/:id", controllers.BuscarEnderecoPorId)
	r.PATCH("/enderecos/:id", controllers.AtualizarEndereco)
	r.DELETE("/enderecos/:id", controllers.ExcluirEndereco)
	r.POST("/enderecos", controllers.CriaEndereco)
	r.GET("/enderecos", controllers.BuscarTodosEnderecos)
}

func handleCargo() {
	r.GET("/cargos/:id", controllers.BuscarCargoPorId)
	r.PATCH("/cargos/:id", controllers.AtualizarCargo)
	r.DELETE("/cargos/:id", controllers.ExcluirCargo)
	r.POST("/cargos", controllers.CriaCargo)
	r.GET("/cargos", controllers.BuscarTodosCargos)
}

func handleCliente() {
	r.GET("/clientes/:id", controllers.BuscarClientePorId)
	r.PATCH("/clientes/:id", controllers.AtualizarCliente)
	r.DELETE("/clientes/:id", controllers.ExcluirCliente)
	r.POST("/clientes", controllers.CriaCliente)
	r.GET("/clientes", controllers.BuscarTodosClientes)
}

func handleEntrega() {
	r.GET("/entregas/:id", controllers.BuscarEntregaPorId)
	r.PATCH("/entregas/:id", controllers.AtualizarEntrega)
	r.DELETE("/entregas/:id", controllers.ExcluirEntrega)
	r.POST("/entregas", controllers.CriaEntrega)
	r.GET("/entregas", controllers.BuscarTodasEntregas)
}

func handleFuncionario() {
	r.GET("/funcionarios/:id", controllers.BuscarFuncionarioPorId)
	r.PATCH("/funcionarios/:id", controllers.AtualizarFuncionario)
	r.DELETE("/funcionarios/:id", controllers.ExcluirFuncionario)
	r.POST("/funcionarios", controllers.CriaFuncionario)
	r.GET("/funcionarios", controllers.BuscarTodosFuncionarios)
}

func handleIngrediente() {
	r.GET("/ingredientes/:id", controllers.BuscarIngredientePorId)
	r.PATCH("/ingredientes/:id", controllers.AtualizarIngrediente)
	r.DELETE("/ingredientes/:id", controllers.ExcluirIngrediente)
	r.POST("/ingredientes", controllers.CriaIngrediente)
	r.GET("/ingredientes", controllers.BuscarTodosIngredientes)
}

func handleMetodoPagamento() {
	r.GET("/metodos_pagamento/:id", controllers.BuscarMetodoPagamentoPorId)
	r.PATCH("/metodos_pagamento/:id", controllers.AtualizarMetodoPagamento)
	r.DELETE("/metodos_pagamento/:id", controllers.ExcluirMetodoPagamento)
	r.POST("/metodos_pagamento", controllers.CriaMetodoPagamento)
	r.GET("/metodos_pagamento", controllers.BuscarTodosMetodoPagamentos)
}

func handlePedido() {
	r.GET("/pedidos/:id", controllers.BuscarPedidoPorId)
	r.PATCH("/pedidos/:id", controllers.AtualizarPedido)
	r.DELETE("/pedidos/:id", controllers.ExcluirPedido)
	r.POST("/pedidos", controllers.CriaPedido)
	r.GET("/pedidos", controllers.BuscarTodosPedidos)
}

func handlePrato() {
	r.GET("/pratos/:id", controllers.BuscarPratoPorId)
	r.PATCH("/pratos/:id", controllers.AtualizarPrato)
	r.DELETE("/pratos/:id", controllers.ExcluirPrato)
	r.POST("/pratos", controllers.CriaPrato)
	r.GET("/pratos", controllers.BuscarTodosPratos)
}
