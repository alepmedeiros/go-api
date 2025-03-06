package routes

import (
	"github.com/alepmedeiros/go-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigurarRotasAuth(r *gin.Engine, db *gorm.DB) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", func(c *gin.Context) { controllers.Login(c, db) })
		auth.GET("/perfil", func(c *gin.Context) { controllers.ListarUsuarios(c, db) })
	}
}
