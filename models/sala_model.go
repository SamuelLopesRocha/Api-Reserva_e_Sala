package models      // vai ser importado como "models.Sala"

type Sala struct {
    SalaID   uint   `json:"sala_id" gorm:"primaryKey;autoIncrement"`   // ID primário gerado automatico
    Recursos string `json:"recursos" gorm:"type:varchar(100);not null"`// Campo obrigatório com até 100 caracteres
    Ativo    bool   `json:"ativo" gorm:"default:false"`                 // Booleano com valor padrão false
}