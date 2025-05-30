package repository

import (
    "reserva_salas_api/config"
    "reserva_salas_api/models"
)

// Buscar todas as salas
func GetAllSalas() ([]models.Sala, error) {
    var salas []models.Sala
    result := config.DB.Find(&salas)
    return salas, result.Error
}

// Criar nova sala
func CreateSala(sala *models.Sala) error {
    result := config.DB.Create(sala)
    return result.Error
}

// Buscar sala por ID
func GetSalaByID(id uint) (*models.Sala, error) {
    var sala models.Sala
    result := config.DB.First(&sala, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &sala, nil
}

// Atualizar sala
func UpdateSala(sala *models.Sala) error {
    result := config.DB.Save(sala)
    return result.Error
}

// Deletar sala
func DeleteSala(id uint) error {
    result := config.DB.Delete(&models.Sala{}, id)
    return result.Error
}
