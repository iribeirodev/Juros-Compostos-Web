package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
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

func main() {
	router := gin.Default()
	router.GET("/", ola)
	router.POST("/calculo", getCalculo)

	router.Run("localhost:8080")
}

func ola(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "ola")
}

func getCalculo(c *gin.Context) {
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
