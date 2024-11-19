package models

type Heroi struct {
	CodigoHeroi       string  `json:"codigo_heroi"`
	NomeReal          string  `json:"nome_real,omitempty"`
	NomeHeroi         string  `json:"nome_heroi"`
	Sexo              string  `json:"sexo"`
	AlturaHeroi       float64 `json:"altura_heroi"`
	PesoHeroi         float64 `json:"peso_heroi"`
	DataNascimento    string  `json:"data_nascimento"`
	LocalNascimento   string  `json:"local_nascimento,omitempty"`
	Poderes           string  `json:"poderes"`
	NivelForca        int     `json:"nivel_forca"`
	Popularidade      int     `json:"popularidade"`
	Status            string  `json:"status"`
	HistoricoBatalhas []int   `json:"historico_batalhas,omitempty"`
}
