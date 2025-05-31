package route

import (
	"reserva_salas_api/controller" // Importa o controller correto

	"github.com/gin-gonic/gin"
)

func ReservaRoutes(router *gin.Engine) {
	reservaGroup := router.Group("/reservas")
	{
		reservaGroup.GET("/", controller.GetReservas)
		reservaGroup.POST("/", controller.CreateReserva)
		reservaGroup.GET("/:id", controller.GetReservaByID)
		reservaGroup.PUT("/:id", controller.UpdateReserva)
		reservaGroup.DELETE("/:id", controller.DeleteReserva)
	}
}
