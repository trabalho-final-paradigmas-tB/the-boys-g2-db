package models

type Crime struct {
	ID               int    `json:"id"`
	NomeCrime        string `json:"nome_crime"`
	Descricao        string `json:"descricao"`
	DataCrime        string `json:"data_crime"`
	HeroiResponsavel int    `json:"heroi_responsavel"`
	Severidade       string `json:"severidade"`
}
