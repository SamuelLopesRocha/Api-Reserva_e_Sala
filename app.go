package main

import (
    "github.com/gin-gonic/gin"
    "reserva_salas_api/config"
    "reserva_salas_api/route"
)

func main() {
    config.ConnectDatabase()

    r := gin.Default()
    route.SalaRoutes(r)
    route.ReservaRoutes(r)


    r.Run(":6000") // Roda o servidor na porta 6000
}
