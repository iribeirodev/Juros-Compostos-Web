package services

import (
	"fmt"
	"iribeirodev/juros-compostos-web/internal/models"
	"math"
)

// Calcula os juros compostos para um determinado período de tempo e retorna um slice com o montante mês a mês
func CalcularJurosCompostos(input models.InputJurosCompostos) []models.Montante {
	var montantes []models.Montante

	taxa := input.Taxa / 100

	for i := 1; i <= int(input.Parcelas); i++ {
		m := input.CapitalInicial * float32(math.Pow(float64(1+taxa), float64(i)))
		montantes = append(montantes, models.Montante{
			Mes: int32(i),
			MA:  fmt.Sprintf("%.2f", m),
		})
	}

	return montantes
}
