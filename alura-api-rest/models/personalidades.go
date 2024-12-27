package models

// Definição de um modelo de dados
// o json é usado para futuramente serializar
type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
