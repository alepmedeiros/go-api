package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Chave secreta usada para assinar os tokens
var secretKey = []byte("alice")

// GerarToken cria um JWT válido por 1 hora
func GerarToken(usuarioID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // 🟢 Alterado para HS256
		"usuario_id": usuarioID,
		"exp":        time.Now().Add(time.Hour).Unix(), // Token expira em 1 hora
	})

	return token.SignedString(secretKey)
}

// ValidarToken verifica se o token JWT é válido
func ValidarToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID := uint(claims["usuario_id"].(float64))
		return usuarioID, nil
	}

	return 0, errors.New("token inválido")
}
