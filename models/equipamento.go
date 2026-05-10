package models

type Equipamento struct {
	EquipamentoID int     `json:"equipamento_id"`
	Descricao     string  `json:"descricao"`
	ValorDiaria   float64 `json:"valor_diaria"`
}
