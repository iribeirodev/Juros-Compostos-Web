package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Verificação de saúde (Ping)
// @Description Rota para verificar se o servidor está ativo (ping)
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{"message": "pong"} "Retorno indicando que o servidor está funcionando"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
