package routes

import (
	"github.com/alepmedeiros/go-api/internal/handler"
	"github.com/alepmedeiros/go-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotasUsuarios(r *gin.Engine, usuarioHandler *handler.UsuarioHandler) {
	usuarios := r.Group("/usuarios")
	usuarios.POST("", usuarioHandler.CriarUsuario)
	usuarios.Use(middleware.AutenticarMiddleware())
	{
		usuarios.GET("", usuarioHandler.ListarUsuarios)
		usuarios.GET("/:id", usuarioHandler.BuscarUsuario)
		usuarios.PUT("/:id", usuarioHandler.AtualizarUsuario)
		usuarios.DELETE("/:id", usuarioHandler.DeletarUsuario)
	}
}
