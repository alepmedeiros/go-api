package controllers

import (
	"net/http"

	"github.com/alepmedeiros/go-api/auth"
	"github.com/alepmedeiros/go-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context, db *gorm.DB) {
	var req struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	var usuario models.Usuario
	if err := db.Where("email = ?", req.Email).First(&usuario).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if usuario.Senha != req.Senha {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha incorreta"})
		return
	}

	token, err := auth.GerarToken(usuario.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Perfil(c *gin.Context, db *gorm.DB) {
	usuarioID, _ := c.Get("usuario_id")

	var usuario models.Usuario
	if err := db.First(&usuario, usuarioID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": usuario.ID, "nome": usuario.Nome, "email": usuario.Email})
}
