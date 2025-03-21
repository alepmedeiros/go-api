package main

import (
	"fmt"
	"log"

	"github.com/alepmedeiros/go-api/config"
	"github.com/alepmedeiros/go-api/internal/domain"
)

func main() {
	fmt.Println("🚀 Iniciando migrações do banco de dados...")

	db := config.ConectarBanco()

	// Executa as migrações
	err := db.AutoMigrate(
		&domain.Usuario{},
	)
	if err != nil {
		log.Fatalf("❌ Erro ao rodar migrações: %v", err)
	}

	fmt.Println("✅ Migrações concluídas!")
}
