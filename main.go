package main

import (
	"github.com/alepmedeiros/go-api/routes"

	"github.com/alepmedeiros/go-api/resources"
	"github.com/gin-gonic/gin"
)

func main() {
	resources.ConectarBanco()

	//Criando o router
	r := gin.Default()

	// Configurando as rotas de usuário
	routes.ConfigurarRotasUsuarios(r, resources.DB)
	routes.ConfigurarRotasAuth(r, resources.DB)

	r.Run(":8080")
}
