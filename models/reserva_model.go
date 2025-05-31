package models      // vai ser importado como "models.reserva_model"

type Reserva struct {
    ReservaID   uint   `json:"reserva_id" gorm:"primaryKey;autoIncrement"` // ID primário
    DataReserva string `json:"data_reserva" gorm:"type:date;not null"`     // Dia da data da reserva
    Descricao   string `json:"descricao" gorm:"type:varchar(100);not null"`// Descrição da reserva

	SalaID uint `json:"sala_id" gorm:"not null"` // Chave estrangeira obrigatória
    Sala   Sala `json:"sala" gorm:"foreignKey:SalaID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"` // Relacionamento

}

