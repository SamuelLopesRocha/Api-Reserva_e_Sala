package config

import (
    "fmt"      
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
    "reserva_salas_api/models" // Corrigido para o nome real do módulo
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("banco.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err)
        panic("Erro ao conectar ao banco")
    }

    // Realiza o AutoMigrate apenas aqui (ou no app.go, mas não nos dois)
    database.AutoMigrate(&models.Sala{})

    DB = database
}
