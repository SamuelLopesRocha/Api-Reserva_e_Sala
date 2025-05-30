package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "reserva_salas_api/models"
    "reserva_salas_api/repository"
)

// Buscar todas as salas
func GetSalas(c *gin.Context) {
    salas, err := repository.GetAllSalas()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar salas"})
        return
    }
    c.JSON(http.StatusOK, salas)
}

// Criar nova sala
func CreateSala(c *gin.Context) {
    var sala models.Sala
    if err := c.ShouldBindJSON(&sala); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := repository.CreateSala(&sala); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar sala"})
        return
    }

    c.JSON(http.StatusCreated, sala)
}

// Buscar sala por ID
func GetSalaByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    sala, err := repository.GetSalaByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Sala não encontrada"})
        return
    }

    c.JSON(http.StatusOK, sala)
}

// Atualizar sala
func UpdateSala(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    sala, err := repository.GetSalaByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Sala não encontrada"})
        return
    }

    if err := c.ShouldBindJSON(&sala); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    sala.SalaID = uint(id) // Garante que o ID não seja sobrescrito
    if err := repository.UpdateSala(sala); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar sala"})
        return
    }

    c.JSON(http.StatusOK, sala)
}

// Deletar sala
func DeleteSala(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    if err := repository.DeleteSala(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar sala"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Sala deletada com sucesso"})
}
