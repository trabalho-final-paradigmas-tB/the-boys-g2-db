package models

type Lutadores struct {
	Lutadores []Heroi `json:"lutadores"`
}

type Local struct {
	Nome      string   `json: "nome"`
	HeroisVan []string `json:"herois_van,omitempty"`
	HeroisDes []string `json:"herois_des,omitempty"`
}

var Ambientes = map[int]Local{
	1: {
		Nome:      "Nova Iorque",
		HeroisVan: []string{"Billy Butcher", "Hughie Campbell", "Luz Estrela", "Kimiko", "Soldier Boy"},
		HeroisDes: []string{"Capitão Patria", "Maeve", "Black Noir", "Trem Bala", "Profundo"},
	},
	2: {
		Nome:      "Vought Tower",
		HeroisVan: []string{"Luz Estrela"},
		HeroisDes: []string{"Capitão Patria", "Black Noir", "Trem Bala", "Cindy"},
	},
	3: {
		Nome:      "Maine",
		HeroisVan: []string{"Billy Butcher", "Leitinho", "Francês"},
		HeroisDes: []string{"Profundo", "Tempesta"},
	},
}

type Turno struct {
	Nome              string `json: "nome"`
	Vida              int    `json: "vida"`
	PoderUsado        string `json: "poder_usado"`
	PopularidadeAtual int    `json: "popularidade_atual"`
}
