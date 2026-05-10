package models

import "time"

type Incidente struct {
	IncidenteID   int       `json:"incidente_id"`
	Data          time.Time `json:"data"`
	Descricao     string    `json:"descricao"`
	EquipamentoID int       `json:"equipamento_id"`
	PessoaID      int       `json:"pessoa_id"`
	StatusID      int       `json:"status_id"`
}
