package domain

type Usuario struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `gorm:"unique" json:"email"`
	Senha string `json:"senha"`
}
