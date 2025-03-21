package main

import (
	"github.com/alepmedeiros/go-api/config"
	"github.com/alepmedeiros/go-api/internal/handler"
	"github.com/alepmedeiros/go-api/internal/repository"
	"github.com/alepmedeiros/go-api/internal/usecase"
	"github.com/alepmedeiros/go-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConectarBanco()

	usuarioRepo := repository.NewUsuarioRepository(db)
	usuarioUseCase := usecase.NewUsuarioUseCase(usuarioRepo)
	usuarioHandler := handler.NewUsuarioHandler(usuarioUseCase)

	authUseCase := usecase.NewAuthUseCase(usuarioRepo)
	authHandler := handler.NewAuthHandler(authUseCase)

	webhookRepo := repository.NewWebhookRepository(db)
	webhookUseCase := usecase.NewWebhookUseCase(webhookRepo)
	webhookHandler := handler.NewWebhookHandler(webhookUseCase)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todas as origens (você pode restringir)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.ConfigurarRotasUsuarios(r, usuarioHandler)
	routes.ConfigurarRotasAuth(r, authHandler)
	routes.ConfigurarRotasWebhook(r, webhookHandler)

	r.Run(":8080")
}
