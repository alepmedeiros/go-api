package models

import "gorm.io/gorm"

type Usuario struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Nome  string `json:"nome"`
	Email string `gorm:"unique" json:"email"`
	Senha string `json:"senha"`
}

func MigrarUsuarios(db *gorm.DB) {
	db.AutoMigrate(&Usuario{})
}
