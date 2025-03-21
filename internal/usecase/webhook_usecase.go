package usecase

import (
	"github.com/alepmedeiros/go-api/internal/domain"
	"github.com/alepmedeiros/go-api/internal/repository"
)

type WebhookUseCase interface {
	CriarWebhook(webhook *domain.Webhook) error
	ListarWebhookPorUsuario(userID uint) ([]domain.Webhook, error)
	BuscarWebhookPorID(id uint) (*domain.Webhook, error)
	AtualizarWebhook(webhook *domain.Webhook) error
	ExcluirWebhook(id uint) error
}

type webhookUseCase struct {
	repo repository.WebhookRepository
}

func NewWebhookUseCase(repo repository.WebhookRepository) WebhookUseCase {
	return &webhookUseCase{repo}
}

func (u *webhookUseCase) CriarWebhook(webhook *domain.Webhook) error {
	return u.repo.Save(webhook)
}

func (u *webhookUseCase) ListarWebhookPorUsuario(userID uint) ([]domain.Webhook, error) {
	return u.repo.FindByUserID(userID)
}

func (u *webhookUseCase) BuscarWebhookPorID(id uint) (*domain.Webhook, error) {
	return u.repo.FindByID(id)
}

func (u *webhookUseCase) AtualizarWebhook(webhook *domain.Webhook) error {
	return u.repo.Update(webhook)
}

func (u *webhookUseCase) ExcluirWebhook(id uint) error {
	return u.repo.Delete(id)
}
