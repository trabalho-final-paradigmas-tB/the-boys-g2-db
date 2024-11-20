package models

type Missoes struct {
	ID            int      `json:"id"`
	Nome          string   `json:"nome"`
	Descrição     string   `json:"descrição"`
	Classificação string   `json:"classificação"`
	Dificuldade   int      `json:"dificuldade"`
	Herois        []string `json:"herois"`
}
