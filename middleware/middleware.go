package middleware

import (
	"net/http"

	"github.com/alepmedeiros/go-api/auth"
	"github.com/gin-gonic/gin"
)

func AutenticarMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		usuarioID, err := auth.ValidarToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("usuario_id", usuarioID)
		c.Next()
	}
}
