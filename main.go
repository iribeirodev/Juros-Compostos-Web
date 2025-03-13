package main

import (
	"fmt"
	_ "iribeirodev/juros-compostos-web/docs"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Entrada struct {
	CapitalInicial float32 `json:"capital"`
	Parcelas       int32   `json:"parcelas"`
	Taxa           float32 `json:"taxa"`
}

type Montante struct {
	Mes int32  `json:"mes"`
	MA  string `json:"montante_acumulado"`
}

// @title Juros Compostos API
// @version 1.0
// @description Exemplo de API com Swagger no Gin
func main() {
	router := gin.Default()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", ping)
	router.POST("/compostos", getCalculoCompostos)

	router.Run(":8080")
}

// Retorna uma resposta "pong" para verificar se a API está funcionando corretamente
// @Summary Verifica o status da API
// @Description Retorna uma resposta "pong" indicando que a API está ativa
// @Tags HealthCheck
// @Produce json
// @Success 200 {object} map[string]string "Resposta com mensagem 'pong'"
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Realiza o cálculo de juros compostos com base nos valores de entrada
// @Summary Calcula juros compostos
// @Description Recebe valores de entrada e retorna o montante acumulado ao longo do tempo
// @Tags Calculo
// @Accept json
// @Produce json
// @Param input body Entrada true "Valores de entrada para o cálculo"
// @Success 200 {array} Montante "Lista de montantes calculados por mês"
// @Failure 400 {object} map[string]string "Mensagem de erro para entrada inválida"
// @Router /compostos [post]
func getCalculoCompostos(c *gin.Context) {
	var input Entrada
	if err := c.BindJSON(&input); err != nil {
		log.Printf("Erro ao vincular JSON -> %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Entrada inválida"})
		return
	}

	// Valores de entrada
	if input.CapitalInicial <= 0 || input.Parcelas <= 0 || input.Taxa <= 0 {
		log.Printf("Os valores de entrada devem ser fornecidos e positivos: %v, %v e %v", input.CapitalInicial, input.Parcelas, input.Taxa)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Os valores devem ser positivos"})
		return
	}

	var montantes []Montante

	parcelas := input.Parcelas
	valor := input.CapitalInicial
	taxa := input.Taxa

	// Taxa em fração
	taxa = taxa / 100

	// Faz o cálculo parcela a parcela e adiciona ao montante
	for i := 1; i <= int(parcelas); i++ {

		m := valor * float32(math.Pow(float64(1+taxa), float64(i)))
		montante := Montante{Mes: int32(i), MA: fmt.Sprintf("%.2f", m)}

		montantes = append(montantes, montante)
	}

	c.JSON(http.StatusOK, montantes)
}
