package handlers

import (
	"iribeirodev/juros-compostos-web/internal/models"
	"iribeirodev/juros-compostos-web/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary GetCalculoJurosCompostos
// @Description Retorna o cálculo de juros compostos baseado nos parâmetros de entrada.
// @Tags Juros
// @Accept json
// @Produce json
// @Param input body models.InputJurosCompostos true "Parâmetros para cálculo de juros compostos"
// @Success 200 {object} models.Juros "Resultado do cálculo"
// @Failure 400 {object} models.Error "Erro de entrada"
// @Router /juros-compostos [post]
func GetCalculoJurosCompostos(c *gin.Context) {
	var input models.InputJurosCompostos
	if err := c.BindJSON(&input); err != nil {
		log.Printf("GetCalculoJurosCompostos - erro ao vincular JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Entrada inválida"})
		return
	}

	result := services.CalcularJurosCompostos(input)
	c.JSON(http.StatusOK, result)
}
