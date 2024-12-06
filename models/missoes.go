package models

type Missoes struct {
	ID              int      `json:"id"`
	Nome            string   `json:"nome"`
	Descrição       string   `json:"descricao"`
	Classificação   string   `json:"classificacao"`
	Dificuldade     int      `json:"dificuldade"`
	Herois          []string `json:"herois"`
	RecompensaTipo  string   `json:"recompensa_tipo"`
	RecompensaValor int      `json:"recompensa_valor"`
}
