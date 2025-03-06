package routes

import (
	"github.com/alepmedeiros/go-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigurarRotasUsuarios(r *gin.Engine, db *gorm.DB) {
	usuarios := r.Group("/usuarios")
	{
		usuarios.POST("", func(c *gin.Context) { controllers.CriarUsuario(c, db) })
		usuarios.GET("", func(c *gin.Context) { controllers.ListarUsuarios(c, db) })
		usuarios.GET("/:id", func(c *gin.Context) { controllers.BuscarUsuario(c, db) })
		usuarios.PUT("/:id", func(c *gin.Context) { controllers.AtualizarUsuario(c, db) })
		usuarios.DELETE("/:id", func(c *gin.Context) { controllers.DeletarUsuario(c, db) })
	}
}
