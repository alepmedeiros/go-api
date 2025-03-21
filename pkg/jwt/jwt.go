package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GerarToken(usuarioID uint) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuario_id": usuarioID,
		"exp":        time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString(secretKey)
}

func ValidarToken(tokenString string) (uint, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

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
