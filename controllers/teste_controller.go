package controllers

import (
	"net/http"

	"github.com/alepmedeiros/go-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Definindo uma rota simples "/"
func RotaSimples(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "API em Go funcionando!"})
}

// Rota comparametro "/saudacao/:nome"
func RotaComParametro(ctx *gin.Context) {
	nome := ctx.Param("nome")
	ctx.JSON(http.StatusOK, gin.H{"message": "Olá, " + nome + "!"})
}

// Recebendo dados via json "/usuarios"
func RotaJSON(ctx *gin.Context, db *gorm.DB) {
	var usuario models.Usuario

	//Bind JSON para strunct
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&usuario)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuário cadastrado com sucesso!",
		"usuario": usuario,
	})
}

func RotaBuscaUsuario(ctx *gin.Context, db *gorm.DB) {
	var usuarios []models.Usuario
	db.Find(&usuarios)
	ctx.JSON(http.StatusOK, usuarios)
}
