package domain

import "time"

type Usuario struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nome      string    `json:"nome"`
	Email     string    `gorm:"unique" json:"email"`
	Senha     string    `json:"senha"`
	Webhooks  []Webhook `gorm:"foreignKey:UserID"` // Relacionamento com Webhooks
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
