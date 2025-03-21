package handler

import (
    "go-api/internal/domain"
    "go-api/internal/usecase"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type DispositivoHandler struct {
    useCase usecase.DispositivoUseCase
}

func NewDispositivoHandler(useCase usecase.DispositivoUseCase) *DispositivoHandler {
    return &DispositivoHandler{useCase}
}

func (h *DispositivoHandler) RegistrarDispositivo(c *gin.Context) {
    var dispositivo domain.Dispositivo
    if err := c.ShouldBindJSON(&dispositivo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }
    if err := h.useCase.RegistrarDispositivo(&dispositivo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar dispositivo"})
        return
    }
    c.JSON(http.StatusCreated, dispositivo)
}

func (h *DispositivoHandler) ListarDispositivosPorUsuario(c *gin.Context) {
    usuarioID, err := strconv.ParseUint(c.Param("usuario_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
        return
    }
    dispositivos, err := h.useCase.ListarDispositivosPorUsuario(uint(usuarioID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dispositivos"})
        return
    }
    c.JSON(http.StatusOK, dispositivos)
}

func (h *DispositivoHandler) AtualizarDispositivo(c *gin.Context) {
    var dispositivo domain.Dispositivo
    if err := c.ShouldBindJSON(&dispositivo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }
    if err := h.useCase.AtualizarDispositivo(&dispositivo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar dispositivo"})
        return
    }
    c.JSON(http.StatusOK, dispositivo)
}

func (h *DispositivoHandler) RemoverDispositivo(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de dispositivo inválido"})
        return
    }
   
