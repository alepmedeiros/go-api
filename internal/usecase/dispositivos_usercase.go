package usecase

import (
	"go-api/internal/domain"
	"go-api/internal/repository"
)

type DispositivoUseCase interface {
	RegistrarDispositivo(dispositivo *domain.Dispositivo) error
	ListarDispositivosPorUsuario(usuarioID uint) ([]domain.Dispositivo, error)
	AtualizarDispositivo(dispositivo *domain.Dispositivo) error
	RemoverDispositivo(id uint) error
}

type dispositivoUseCase struct {
	dispositivoRepo repository.DispositivoRepository
}

func NewDispositivoUseCase(dr repository.DispositivoRepository) DispositivoUseCase {
	return &dispositivoUseCase{dr}
}

func (uc *dispositivoUseCase) RegistrarDispositivo(dispositivo *domain.Dispositivo) error {
	return uc.dispositivoRepo.Salvar(dispositivo)
}

func (uc *dispositivoUseCase) ListarDispositivosPorUsuario(usuarioID uint) ([]domain.Dispositivo, error) {
	return uc.dispositivoRepo.BuscarPorUsuarioID(usuarioID)
}

func (uc *dispositivoUseCase) AtualizarDispositivo(dispositivo *domain.Dispositivo) error {
	return uc.dispositivoRepo.Atualizar(dispositivo)
}

func (uc *dispositivoUseCase) RemoverDispositivo(id uint) error {
	return uc.dispositivoRepo.Deletar(id)
}
