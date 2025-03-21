package handler

import (
	"net/http"
	"strings"

	"github.com/alepmedeiros/go-api/internal/domain"
	"github.com/alepmedeiros/go-api/internal/usecase"
	"github.com/alepmedeiros/go-api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	useCase usecase.AuthUseCase
}

func NewAuthHandler(useCase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{useCase}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var credenciais domain.Credenciais
	if err := c.ShouldBindJSON(&credenciais); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	token, err := h.useCase.Login(credenciais.Email, credenciais.Senha)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Perfil do usuário autenticado
func (h *AuthHandler) Perfil(c *gin.Context) {
	// Obtém o token do Header
	tokenString := c.GetHeader("Authorization")

	// Remove "Bearer " do token
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Valida o token e extrai o ID do usuário
	usuarioID, err := jwt.ValidarToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	usuario, err := h.useCase.Perfil(uint(usuarioID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      usuario.ID,
		"nome":    usuario.Nome,
		"email":   usuario.Email,
		"webhook": usuario.Webhooks,
	})
}
