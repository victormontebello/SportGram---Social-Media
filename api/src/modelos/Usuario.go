package modelos

import "time"

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64   `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Email    string `json:"email,omitempty"`
	Senha    string `json:"senha"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
	Esporte  string `json:"esporte,omitempty"`
	AnosExperiencia int `json:"anosExperiencia,omitempty"`
	PossuiPatrocinio bool `json:"possuiPatrocinio,omitempty"`
}