package repository

import (
	"go-api/internal/domain"

	"gorm.io/gorm"
)

type DispositivoRepository interface {
	Salvar(dispositivo *domain.Dispositivo) error
	BuscarPorID(id uint) (*domain.Dispositivo, error)
	BuscarPorUsuarioID(usuarioID uint) ([]domain.Dispositivo, error)
	Atualizar(dispositivo *domain.Dispositivo) error
	Deletar(id uint) error
}

type dispositivoRepo struct {
	db *gorm.DB
}

func NewDispositivoRepository(db *gorm.DB) DispositivoRepository {
	return &dispositivoRepo{db}
}

func (r *dispositivoRepo) Salvar(dispositivo *domain.Dispositivo) error {
	return r.db.Create(dispositivo).Error
}

func (r *dispositivoRepo) BuscarPorID(id uint) (*domain.Dispositivo, error) {
	var dispositivo domain.Dispositivo
	err := r.db.First(&dispositivo, id).Error
	return &dispositivo, err
}

func (r *dispositivoRepo) BuscarPorUsuarioID(usuarioID uint) ([]domain.Dispositivo, error) {
	var dispositivos []domain.Dispositivo
	err := r.db.Where("usuario_id = ?", usuarioID).Find(&dispositivos).Error
	return dispositivos, err
}

func (r *dispositivoRepo) Atualizar(dispositivo *domain.Dispositivo) error {
	return r.db.Save(dispositivo).Error
}

func (r *dispositivoRepo) Deletar(id uint) error {
	return r.db.Delete(&domain.Dispositivo{}, id).Error
}
