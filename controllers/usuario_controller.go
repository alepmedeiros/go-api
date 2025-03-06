package controllers

import (
	"net/http"

	"github.com/alepmedeiros/go-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Criar um novo usuário
func CriarUsuario(c *gin.Context, db *gorm.DB) {
	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := db.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// Listar todos os usuários
func ListarUsuarios(c *gin.Context, db *gorm.DB) {
	var usuarios []models.Usuario
	db.Find(&usuarios)
	c.JSON(http.StatusOK, usuarios)
}

// Buscar um usuário específico
func BuscarUsuario(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var usuario models.Usuario

	if err := db.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// Atualizar um usuário
func AtualizarUsuario(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var usuario models.Usuario

	if err := db.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	db.Save(&usuario)
	c.JSON(http.StatusOK, usuario)
}

// Deletar um usuário
func DeletarUsuario(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var usuario models.Usuario

	if err := db.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	db.Delete(&usuario)
	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado"})
}
