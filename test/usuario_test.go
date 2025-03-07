package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alepemedeiros/go-api/internal/domain"
	"github.com/alepemedeiros/go-api/internal/handler"
	"github.com/alepemedeiros/go-api/internal/repository"
	"github.com/alepemedeiros/go-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&domain.Usuario{})
	return db
}

func TestCriarUsuario(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewUsuarioRepository(db)
	useCase := usecase.NewUsuarioUseCase(repo)
	handler := handler.NewUsuarioHandler(useCase)

	r := gin.Default()
	r.POST("/usuarios", handler.CriarUsuario)

	usuario := domain.Usuario{
		Nome:  "Teste",
		Email: "teste@email.com",
		Senha: "123456",
	}

	jsonValue, _ := json.Marshal(usuario)
	req, _ := http.NewRequest("POST", "/usuarios", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListarUsuarios(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewUsuarioRepository(db)
	useCase := usecase.NewUsuarioUseCase(repo)
	handler := handler.NewUsuarioHandler(useCase)

	r := gin.Default()
	r.GET("/usuarios", handler.ListarUsuarios)

	req, _ := http.NewRequest("GET", "/usuarios", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
