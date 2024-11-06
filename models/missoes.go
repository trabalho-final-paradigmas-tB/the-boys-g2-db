package models

type Missoes struct {
	Nome          string `json:"Nome da missão"`
	Descrição     string `json:"Descrição da Missão"`
	Classificação string `json:"Classificação"`
	Rank_SS       string `json:"Rank_SS"`
	Rank_S        string `json:"Rank_S"`
	Rank_A        string `json:"Rank_A"`
	Rank_B        string `json:"Rank_B"`
	Rank_C        string `json:"Rank_C"`
	Rank_D        string `json:"Rank_D"`
	Rank_E        string `json:"Rank_E"`
}
