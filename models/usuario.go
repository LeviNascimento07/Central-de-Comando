package models

import "time"

type Usuario struct {
	UsuarioID     int       `json:"usuario_id"`
	Nome          string    `json:"nome"`
	Login         string    `json:"login"`
	Senha         string    `json:"senha,omitempty"`
	AtualizadoEm  time.Time `json:"atualizado_em"`
	AtualizadoPor *int      `json:"atualizado_por,omitempty"`
}
