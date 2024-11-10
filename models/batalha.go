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
		Nome:      "Brasília",
		HeroisVan: []string{"Batman", "Superman"},
		HeroisDes: []string{"Coringa", "Lex Luthor"},
	},
	2: {
		Nome:      "Gotham",
		HeroisVan: []string{"Batman"},
		HeroisDes: []string{"Pinguim"},
	},
}
