package repository

import (
	"github.com/alepmedeiros/go-api/internal/domain"

	"gorm.io/gorm"
)

type WebhookRepository interface {
	Save(webhook *domain.Webhook) error
	FindByID(id uint) (*domain.Webhook, error)
	FindByUserID(userID uint) ([]domain.Webhook, error)
	Update(webhook *domain.Webhook) error
	Delete(id uint) error
}

type webhookRepo struct {
	db *gorm.DB
}

func NewWebhookRepository(db *gorm.DB) WebhookRepository {
	return &webhookRepo{db}
}

func (r *webhookRepo) Save(webhook *domain.Webhook) error {
	return r.db.Create(webhook).Error
}

func (r *webhookRepo) FindByID(id uint) (*domain.Webhook, error) {
	var webhook domain.Webhook
	err := r.db.First(&webhook, id).Error
	if err != nil {
		return nil, err
	}
	return &webhook, nil
}

func (r *webhookRepo) FindByUserID(userID uint) ([]domain.Webhook, error) {
	var webhooks []domain.Webhook
	err := r.db.Where("user_id = ?", userID).Find(&webhooks).Error
	if err != nil {
		return nil, err
	}
	return webhooks, nil
}

func (r *webhookRepo) Update(webhook *domain.Webhook) error {
	return r.db.Model(&domain.Webhook{}).
		Where("id = ?", webhook.ID).
		Updates(webhook).Error
}

func (r *webhookRepo) Delete(id uint) error {
	return r.db.Delete(&domain.Webhook{}, id).Error
}
