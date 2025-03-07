package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alepmedeiros/go-api/internal/domain"
	"github.com/alepmedeiros/go-api/internal/handler"
	"github.com/alepmedeiros/go-api/internal/repository"
	"github.com/alepmedeiros/go-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewUsuarioRepository(db)
	useCase := usecase.NewAuthUseCase(repo)
	handler := handler.NewAuthHandler(useCase)

	r := gin.Default()
	r.POST("/auth/login", handler.Login)

	usuario := domain.Usuario{
		Nome:  "Teste",
		Email: "teste@email.com",
		Senha: "123456",
	}
	db.Create(&usuario)

	credenciais := domain.Credenciais{
		Email: "teste@email.com",
		Senha: "123456",
	}
	jsonValue, _ := json.Marshal(credenciais)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
