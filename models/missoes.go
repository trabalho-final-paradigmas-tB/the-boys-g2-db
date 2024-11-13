package models

type Missoes struct {
	Nome          string `json:"nome_da_missão"`
	Descrição     string `json:"descrição_da_missão"`
	Classificação string `json:"classificação"`
	Dificuldade   int    `json:"dificuldade"`
}
