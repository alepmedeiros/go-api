package handler

import (
	"net/http"

	"github.com/alepmedeiros/go-api/internal/domain"
	"github.com/alepmedeiros/go-api/internal/usecase"

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
