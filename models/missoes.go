package models

type Missoes struct {
	Nome           string `json:"nome da missão"`
	Descrição      string `json:"descrição da missão"`
	Classificação  string `json:"classificação"`
	Rank_Escolhida string `json:"rank_escolhido"`
	Rank_SS        string `json:"rank_SS"`
	Rank_S         string `json:"rank_S"`
	Rank_A         string `json:"rank_A"`
	Rank_B         string `json:"rank_B"`
	Rank_C         string `json:"rank_C"`
	Rank_D         string `json:"rank_D"`
	Rank_E         string `json:"rank_E"`
}
