package sala

import (
    "github.com/gin-gonic/gin"
    "reserva_salas_api/controller" 
)


func SalaRoutes(router *gin.Engine) {
    salaGroup := router.Group("/salas")
    {
        salaGroup.GET("/", controller.GetSalas)       
        salaGroup.POST("/", controller.CreateSala)
        salaGroup.GET("/:id", controller.GetSalaByID)
        salaGroup.PUT("/:id", controller.UpdateSala)
        salaGroup.DELETE("/:id", controller.DeleteSala)
    }
}