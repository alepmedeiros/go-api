package routes

import (
	"github.com/alepmedeiros/go-api/internal/handler"
	"github.com/alepmedeiros/go-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigurarRotasWebhook(r *gin.Engine, webhookHandler *handler.WebhookHandler) {
	webhooks := r.Group("/webhook")
	webhooks.Use(middleware.AutenticarMiddleware())
	{
		webhooks.GET("", webhookHandler.ListarWebhooks)
		webhooks.POST("", webhookHandler.CriarWebhook)
		webhooks.GET("/:id", webhookHandler.BuscarWebhook)
		webhooks.PUT("/:id", webhookHandler.AtualizarWebhook)
		webhooks.DELETE("/:id", webhookHandler.ExcluirWebhook)
	}
}
