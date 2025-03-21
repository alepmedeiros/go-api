package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/alepmedeiros/go-api/internal/domain"
	"github.com/alepmedeiros/go-api/internal/usecase"
	"github.com/alepmedeiros/go-api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	webhookCase usecase.WebhookUseCase
}

func NewWebhookHandler(webhookCase usecase.WebhookUseCase) *WebhookHandler {
	return &WebhookHandler{webhookCase}
}

func (h *WebhookHandler) CriarWebhook(c *gin.Context) {
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

	var webhook domain.Webhook
	if err := c.ShouldBindJSON(&webhook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	webhook.UserID = usuarioID

	if err := h.webhookCase.CriarWebhook(&webhook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar webhook"})
		return
	}

	c.JSON(http.StatusOK, webhook)
}

func (h *WebhookHandler) ListarWebhooks(c *gin.Context) {
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

	// Busca os Webhooks do usuário autenticado
	webhooks, err := h.webhookCase.ListarWebhookPorUsuario(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}

	// Retorna os Webhooks
	c.JSON(http.StatusOK, webhooks)
}

func (h *WebhookHandler) BuscarWebhook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	webhook, err := h.webhookCase.BuscarWebhookPorID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Webhook não encontrado"})
		return
	}
	c.JSON(http.StatusOK, webhook)
}

func (h *WebhookHandler) AtualizarWebhook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var webhook domain.Webhook
	if err := c.ShouldBindJSON(&webhook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	webhook.ID = uint(id)

	if err := h.webhookCase.AtualizarWebhook(&webhook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar webhook"})
		return
	}

	c.JSON(http.StatusOK, webhook)
}

func (h *WebhookHandler) ExcluirWebhook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.webhookCase.ExcluirWebhook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir webhook"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook deletado"})
}
