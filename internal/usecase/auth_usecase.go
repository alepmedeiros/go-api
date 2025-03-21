package usecase

import (
	"errors"

	"github.com/alepmedeiros/go-api/internal/domain"
	"github.com/alepmedeiros/go-api/internal/repository"
	"github.com/alepmedeiros/go-api/pkg/jwt"
)

type AuthUseCase interface {
	Login(email, senha string) (string, error)
	Perfil(userID uint) (*domain.Usuario, error)
}

type authUseCase struct {
	usuarioRepo repository.UsuarioRepository
}

func NewAuthUseCase(usuarioRepo repository.UsuarioRepository) AuthUseCase {
	return &authUseCase{usuarioRepo}
}

func (a *authUseCase) Login(email, senha string) (string, error) {
	usuario, err := a.usuarioRepo.BuscarPorEmail(email)
	if err != nil {
		return "", errors.New("usuário não encontrado")
	}

	if usuario.Senha != senha {
		return "", errors.New("senha incorreta")
	}

	token, err := jwt.GerarToken(usuario.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *authUseCase) Perfil(userID uint) (*domain.Usuario, error) {
	return a.usuarioRepo.BuscarPorID(userID)
}
