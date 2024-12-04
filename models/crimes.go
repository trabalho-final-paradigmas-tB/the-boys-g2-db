package models

type Crime struct {
	ID               int    `json:"id"`
	NomeCrime        string `json:"nome"`
	Descricao        string `json:"descricao"`
	DataCrime        string `json:"data"`
	HeroiResponsavel string `json:"heroi_responsavel"`
	Severidade       int    `json:"severidade"`
	Oculto           bool   `json:"oculto"`
}
