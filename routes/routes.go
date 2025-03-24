package routes

import (
	"iribeirodev/juros-compostos-web/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Inicializa as rotas da aplicação
func IniciarRotas() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)
	router.POST("/juros-compostos", handlers.GetCalculoJurosCompostos)
	router.POST("/simulacao-parcelas", handlers.GetSimulacaoParcelas)

	return router
}
