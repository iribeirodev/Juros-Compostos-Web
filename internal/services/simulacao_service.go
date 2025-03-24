package services

import (
	"fmt"
	"iribeirodev/juros-compostos-web/internal/models"
	"math"
)

// Simula as parcelas de um financiamento com base na taxa de juros, valor do produto e parcelas, retorna um slice com os detalhes parcela a parcela.
// Fórmula para calcular o valor da parcela (parcela fixa):
// Parcela = P * (i * (1 + i)^n) / ((1 + i)^n - 1)
//
// Onde:
// P = Valor do produto,
// i = Taxa de juros mensal (Taxa / 100),
// n = Número de parcelas
func SimularParcelas(input models.InputSimulacaoParcela) []models.Simulacao {
	var simulacoes []models.Simulacao

	taxaJuros := input.Taxa / 100

	valorParcela := input.ValorProduto * (taxaJuros * math.Pow(1+taxaJuros, float64(input.Parcelas))) /
		(math.Pow(1+taxaJuros, float64(input.Parcelas)) - 1)

	saldoDevedor := input.ValorProduto

	for i := 1; i <= int(input.Parcelas); i++ {
		jurosPago := saldoDevedor * taxaJuros
		amortizacao := valorParcela - jurosPago

		// Atualiza o saldo devedor
		saldoDevedor -= amortizacao

		simulacoes = append(simulacoes, models.Simulacao{
			Mes:          int32(i),
			ValorParcela: fmt.Sprintf("%.2f", valorParcela),
			JurosPago:    fmt.Sprintf("%.2f", jurosPago),
			Amortizacao:  fmt.Sprintf("%.2f", amortizacao),
			SaldoDevedor: fmt.Sprintf("%.2f", saldoDevedor),
		})
	}

	return simulacoes
}
