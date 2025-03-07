package main

import (
	"github.com/alepmedeiros/go-api/config"
	"github.com/alepmedeiros/go-api/internal/handler"
	"github.com/alepmedeiros/go-api/internal/repository"
	"github.com/alepmedeiros/go-api/internal/usecase"
	"github.com/alepmedeiros/go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConectarBanco()

	usuarioRepo := repository.NewUsuarioRepository(db)
	usuarioUseCase := usecase.NewUsuarioUseCase(usuarioRepo)
	usuarioHandler := handler.NewUsuarioHandler(usuarioUseCase)

	authUseCase := usecase.NewAuthUseCase(usuarioRepo)
	authHandler := handler.NewAuthHandler(authUseCase)

	r := gin.Default()
	routes.ConfigurarRotasUsuarios(r, usuarioHandler)
	routes.ConfigurarRotasAuth(r, authHandler)

	r.Run(":8080")
}
