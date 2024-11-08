package models

type Lutadores struct {
	Lutadores []Heroi `json:"lutadores"`
}

type Local struct {
	HeroisVan []string `json:"herois_van,omitempty"`
	HeroisDes []string `json:"herois_des,omitempty"`
}

var Ambientes = map[string]Local{
	"Brasília": {
		HeroisVan: []string{},
		HeroisDes: []string{},
	},
}
