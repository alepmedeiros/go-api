package repository

import (
	"github.com/alepmedeiros/go-api/internal/domain"

	"gorm.io/gorm"
)

type UsuarioRepository interface {
	Salvar(usuario *domain.Usuario) error
	BuscarPorEmail(email string) (*domain.Usuario, error)
	BuscarPorID(id uint) (*domain.Usuario, error)
	Atualizar(usuario *domain.Usuario) error
	Deletar(id uint) error
	BuscarTodos() ([]domain.Usuario, error)
}

type usuarioRepo struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepo{db}
}

func (r *usuarioRepo) Salvar(usuario *domain.Usuario) error {
	return r.db.Create(usuario).Error
}

func (r *usuarioRepo) BuscarPorEmail(email string) (*domain.Usuario, error) {
	var usuario domain.Usuario
	err := r.db.Where("email = ?", email).First(&usuario).Error
	return &usuario, err
}

func (r *usuarioRepo) BuscarPorID(id uint) (*domain.Usuario, error) {
	var usuario domain.Usuario
	err := r.db.First(&usuario, id).Error
	return &usuario, err
}

func (r *usuarioRepo) Atualizar(usuario *domain.Usuario) error {
	return r.db.Save(usuario).Error
}

func (r *usuarioRepo) Deletar(id uint) error {
	return r.db.Delete(&domain.Usuario{}, id).Error
}

func (r *usuarioRepo) BuscarTodos() ([]domain.Usuario, error) {
	var usuarios []domain.Usuario
	err := r.db.Find(&usuarios).Error
	return usuarios, err
}
