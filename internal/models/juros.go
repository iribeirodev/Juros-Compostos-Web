package models

type InputJurosCompostos struct {
	CapitalInicial float32 `json:"capital"`
	Parcelas       int32   `json:"parcelas"`
	Taxa           float32 `json:"taxa"`
}

type Montante struct {
	Mes int32  `json:"mes"`
	MA  string `json:"montante_acumulado"`
}
