package handlers

import (
	"iribeirodev/juros-compostos-web/internal/models"
	"iribeirodev/juros-compostos-web/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Simula Parcelas
// @Description Simula o pagamento de parcelas com base no input fornecido
// @Tags Simulação
// @Accept json
// @Produce json
// @Param input body models.InputSimulacaoParcela true "Dados da simulação"
// @Success 200 {object} models.SimulacaoParcela
// @Failure 400 {object} models.Error
// @Router /simulacao/parcelas [post]
func GetSimulacaoParcelas(c *gin.Context) {
	var input models.InputSimulacaoParcela
	if err := c.BindJSON(&input); err != nil {
		log.Printf("GetSimulacaoParcelas - Erro ao vincular JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Entrada inválida"})
		return
	}

	result := services.SimularParcelas(input)
	c.JSON(http.StatusOK, result)
}
