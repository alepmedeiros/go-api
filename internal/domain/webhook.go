package domain

import "time"

type Webhook struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"` // Chave estrangeira para o usuário
	URL       string    `gorm:"not null" json:"url"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
