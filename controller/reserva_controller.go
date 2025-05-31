package controller

import (
    "net/http"
    "strconv"
	"time"

    "github.com/gin-gonic/gin"
    "reserva_salas_api/models"
    "reserva_salas_api/repository"
	"reserva_salas_api/config"
)

// Buscar todas as reservas
func GetReservas(c *gin.Context) {
    reservas, err := repository.GetAllReservas()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar Reservas"})
        return
    }
    c.JSON(http.StatusOK, reservas)
}

// Criar nova Reserva
func CreateReserva(c *gin.Context) {
    var reserva models.Reserva

    // Tenta converter o JSON da requisição para a estrutura da Reserva. Retorna erro 400 se o JSON estiver inválido.
    if err := c.ShouldBindJSON(&reserva); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }

    // Valida o formato da data
    layout := "2006-01-02" // formato esperado: YYYY-MM-DD
    dataReservaParsed, err := time.Parse(layout, reserva.DataReserva)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido. Use YYYY-MM-DD"})
        return
    }

    // Verifica se a data está no passado
    if dataReservaParsed.Before(time.Now().Truncate(24 * time.Hour)) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "A data da reserva não pode ser no passado"})
        return
    }

    // Verifica se a sala existe no banco de dados
    var sala models.Sala
    if err := config.DB.First(&sala, reserva.SalaID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Sala não encontrada"})
        return
    }

    // Verifica se já existe uma reserva para a mesma sala no mesmo dia
    var reservaExistente models.Reserva
    if err := config.DB.Where("sala_id = ? AND data_reserva = ?", reserva.SalaID, reserva.DataReserva).First(&reservaExistente).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Já existe uma reserva para esta sala nesta data"})
        return
    }

    // Cria a reserva no banco
    if err := config.DB.Create(&reserva).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar a reserva"})
        return
    }

    // Retorna a reserva criada
    c.JSON(http.StatusCreated, gin.H{"message": "Reserva criada"})

}

// Buscar sala por ID
func GetReservaByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    reserva, err := repository.GetReservaByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reserva não encontrada"})
        return
    }

    c.JSON(http.StatusOK, reserva)
}

// Atualizar sala
func UpdateReserva(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    // Busca a reserva atual pelo ID
    reservaAtual, err := repository.GetReservaByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reserva não encontrada"})
        return
    }

    // Cria uma variável temporária para bind do JSON recebido
    var novaReserva models.Reserva
    if err := c.ShouldBindJSON(&novaReserva); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }

    // Valida o formato da data e se não está no passado
    layout := "2006-01-02"
    dataReserva, err := time.Parse(layout, novaReserva.DataReserva)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato da data inválido. Use yyyy-mm-dd"})
        return
    }
    if dataReserva.Before(time.Now().Truncate(24 * time.Hour)) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "A data da reserva não pode ser no passado"})
        return
    }

    // Verifica se a sala existe
    var sala models.Sala
    if err := config.DB.First(&sala, novaReserva.SalaID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Sala não encontrada"})
        return
    }

    // Verifica se já existe outra reserva para a mesma sala e data (excluindo a atual)
    var reservaExistente models.Reserva
    if err := config.DB.
        Where("sala_id = ? AND data_reserva = ? AND reserva_id <> ?", novaReserva.SalaID, novaReserva.DataReserva, id).
        First(&reservaExistente).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Já existe uma reserva para esta sala nesta data"})
        return
    }

    // Atualiza os campos da reserva atual com os novos dados
    reservaAtual.SalaID = novaReserva.SalaID
    reservaAtual.DataReserva = novaReserva.DataReserva
    reservaAtual.Descricao = novaReserva.Descricao // se quiser atualizar descrição também

    // Salva a atualização no banco
    if err := repository.UpdateReserva(reservaAtual); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar reserva"})
        return
    }

    c.JSON(http.StatusOK, reservaAtual)
}

// Deletar reserva
func DeleteReserva	(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    if err := repository.DeleteReserva(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar reserva"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Reserva deletada com sucesso"})
}