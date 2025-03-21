package config

import (
	"fmt"
	"log"
	"os"

	"github.com/alepmedeiros/go-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func excuteMigrate(db *gorm.DB) {
	// Rodar a migração automática
	err := db.AutoMigrate(
		&domain.Usuario{},
		&domain.Webhook{},
	)
	if err != nil {
		log.Fatal("Erro ao rodar migração:", err)
	}

	fmt.Println("📌 Banco de dados conectado e migrado com sucesso!")
}

func ConectarBanco() *gorm.DB {
	// Criando a conexão com o banco
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco de dados")
	}

	excuteMigrate(db)

	return db
}
