package models

type Crime struct {
	NomeCrime        string `json:"nome_crime"`
	Descricao        string `json:"descricao"`
	DataCrime        string `json:"data_crime"`
	HeroiResponsavel int    `json:"heroi_responsavel"`
	Severidade       int    `json:"severidade"`
}
