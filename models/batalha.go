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
		HeroisDes: []string{"Capitão Pátria", "Maeve", "Black Noir", "Trem Bala", "Profundo"},
	},
	2: {
		Nome:      "Vought Tower",
		HeroisVan: []string{"Luz Estrela"},
		HeroisDes: []string{"Capitão Pátria", "Black Noir", "Trem Bala", "Cindy"},
	},
	3: {
		Nome:      "Maine",
		HeroisVan: []string{"Billy Butcher", "Leitinho", "Francês"},
		HeroisDes: []string{"Profundo", "Tempesta"},
	},
	4: {
		Nome:      "Farmácia do Leite",
		HeroisVan: []string{"Billy Butcher", "Leitinho", "Francês", "Hughie Campbell"},
		HeroisDes: []string{"Capitão Pátria", "Black Noir"},
	},
	5: {
		Nome:      "Casa Segura da CIA",
		HeroisVan: []string{"Grace Mallory", "Billy Butcher", "Hughie Campbell"},
		HeroisDes: []string{"Stormfront", "Maeve", "Trem Bala"},
	},
	6: {
		Nome:      "Parque Temático “Herogasm”",
		HeroisVan: []string{"Soldier Boy", "Luz Estrela"},
		HeroisDes: []string{"Capitão Pátria", "Maeve", "Black Noir"},
	},
	7: {
		Nome:      "Casa de Maeve",
		HeroisVan: []string{"Maeve"},
		HeroisDes: []string{"Billy Butcher", "Francês", "Hughie Campbell"},
	},
	8: {
		Nome:      "O Litoral",
		HeroisVan: []string{"Profundo"},
		HeroisDes: []string{"Billy Butcher", "Luz Estrela", "Leitinho"},
	},
}

type Turno struct {
	Nome              string `json: "nome"`
	Vida              int    `json: "vida"`
	PoderUsado        string `json: "poder_usado"`
	PopularidadeAtual int    `json: "popularidade_atual"`
}

type Evento struct {
	Nome          string
	Consequencias string
}

type ResultadoTurno struct {
	Turnos        []Turno `json:"turnos"`
	Evento        string  `json:"evento"`
	Consequencias string  `json:"consequencias"`
}
