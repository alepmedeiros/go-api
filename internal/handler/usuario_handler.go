package handler

import (
	"net/http"
	"strconv"

	"github.com/alepmedeiros/go-api/internal/domain"
	"github.com/alepmedeiros/go-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	useCase usecase.UsuarioUseCase
}

func NewUsuarioHandler(useCase usecase.UsuarioUseCase) *UsuarioHandler {
	return &UsuarioHandler{useCase}
}

func (h *UsuarioHandler) CriarUsuario(c *gin.Context) {
	var usuario domain.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := h.useCase.CriarUsuario(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (h *UsuarioHandler) ListarUsuarios(c *gin.Context) {
	usuarios, err := h.useCase.ListarUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

func (h *UsuarioHandler) BuscarUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	usuario, err := h.useCase.BuscarUsuarioPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func (h *UsuarioHandler) AtualizarUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var usuario domain.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	usuario.ID = uint(id)

	if err := h.useCase.AtualizarUsuario(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (h *UsuarioHandler) DeletarUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.useCase.DeletarUsuario(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado"})
}
