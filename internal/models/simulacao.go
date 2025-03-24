package models

type InputSimulacaoParcela struct {
	ValorProduto float64 `json:"valor_produto"`
	Parcelas     int32   `json:"parcelas"`
	Taxa         float64 `json:"taxa_anual"`
}

type Simulacao struct {
	Mes          int32  `json:"mes"`
	ValorParcela string `json:"valor_parcela"`
	JurosPago    string `json:"juros_pago"`
	Amortizacao  string `json:"amortizacao"`
	SaldoDevedor string `json:"saldo_devedor"`
}
