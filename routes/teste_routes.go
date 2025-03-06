package routes

import (
	"github.com/alepmedeiros/go-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigurarRotasTeste(r *gin.Engine, db *gorm.DB) {
	teste := r.Group("/")
	{
		teste.GET("", func(c *gin.Context) { controllers.RotaSimples(c) })
		teste.GET("/saudacao/:nome", func(c *gin.Context) { controllers.RotaComParametro(c) })
		teste.POST("/usuarios", func(c *gin.Context) { controllers.RotaJSON(c, db) })
		teste.GET("/usuarios", func(c *gin.Context) { controllers.RotaBuscaUsuario(c, db) })
	}
}
