package repository

import (
	"reserva_salas_api/config"
	"reserva_salas_api/models"
)

// Buscar todas as reservas
func GetAllReservas() ([]models.Reserva, error) {
	var reservas []models.Reserva
	result := config.DB.Preload("Sala").Find(&reservas) //Preload("Sala"), Carrega a sala vinculada à reserva automaticamente.
	return reservas, result.Error
}

// Criar nova reseva
func CreateReserva(reserva *models.Reserva) error {
	result := config.DB.Create(reserva)
	return result.Error
}

// Buscar reserva por ID
func GetReservaByID(id uint) (*models.Reserva, error) {
	var reserva models.Reserva
	result := config.DB.Preload("Sala").First(&reserva, id) //Preload("Sala"), Carrega a sala vinculada à reserva automaticamente.
	if result.Error != nil {
		return nil, result.Error
	}
	return &reserva, nil
}

// Atualizar reserva
func UpdateReserva(reserva *models.Reserva) error {
	result := config.DB.Save(reserva)
	return result.Error
}

// Deletar reserva
func DeleteReserva(id uint) error {
	result := config.DB.Delete(&models.Reserva{}, id)
	return result.Error
}
