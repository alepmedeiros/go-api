package routes

import (
	"github.com/alepmedeiros/go-api/internal/handler"
	"github.com/alepmedeiros/go-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotasAuth(r *gin.Engine, authHandler *handler.AuthHandler) {
	auth := r.Group("/auth")
	auth.POST("/login", authHandler.Login)
	auth.Use(middleware.AutenticarMiddleware())
	{
		auth.GET("/perfil", authHandler.Perfil)
	}
}
