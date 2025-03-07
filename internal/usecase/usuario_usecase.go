package usecase

import (
	"github.com/alepmedeiros/go-api/internal/domain"
	"github.com/alepmedeiros/go-api/internal/repository"
)

type UsuarioUseCase interface {
	CriarUsuario(usuario *domain.Usuario) error
	ListarUsuarios() ([]domain.Usuario, error)
	BuscarUsuarioPorID(id uint) (*domain.Usuario, error)
	AtualizarUsuario(usuario *domain.Usuario) error
	DeletarUsuario(id uint) error
}

type usuarioUseCase struct {
	repo repository.UsuarioRepository
}

func NewUsuarioUseCase(repo repository.UsuarioRepository) UsuarioUseCase {
	return &usuarioUseCase{repo}
}

func (u *usuarioUseCase) CriarUsuario(usuario *domain.Usuario) error {
	return u.repo.Salvar(usuario)
}

func (u *usuarioUseCase) ListarUsuarios() ([]domain.Usuario, error) {
	return u.repo.BuscarTodos()
}

func (u *usuarioUseCase) BuscarUsuarioPorID(id uint) (*domain.Usuario, error) {
	return u.repo.BuscarPorID(id)
}

func (u *usuarioUseCase) AtualizarUsuario(usuario *domain.Usuario) error {
	return u.repo.Atualizar(usuario)
}

func (u *usuarioUseCase) DeletarUsuario(id uint) error {
	return u.repo.Deletar(id)
}
