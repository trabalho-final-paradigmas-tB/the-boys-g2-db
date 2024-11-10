package handlers

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func ChamarBatalha(w http.ResponseWriter, r *http.Request) {
	var luts models.Lutadores
	if err := json.NewDecoder(r.Body).Decode(&luts); err != nil {
		http.Error(w, "Erro ao receber herois"+err.Error(), http.StatusBadRequest)
		return
	}

	Amb3 := ChamarAmbiente()

	for _, heroiNome := range Amb3.HeroisVan {
		for _, lutador := range luts.Lutadores {
			if strings.EqualFold(heroiNome, lutador.NomeHeroi) {
				fmt.Printf("Herroi encontrado: %s no ambiente: %s\n", lutador.NomeHeroi, Amb3.Nome)
			}
		}
	}

	for _, heroiNome := range Amb3.HeroisDes {
		for _, lutador := range luts.Lutadores {
			if strings.EqualFold(heroiNome, lutador.NomeHeroi) {
				fmt.Printf("Herroi encontrado: %s no ambiente: %s\n", lutador.NomeHeroi, Amb3.Nome)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Herói atualizado com sucesso!"})

}

func ChamarAmbiente() models.Local {
	rand.Seed(time.Now().UnixNano())
	N := 2
	numeroAleatorio := rand.Intn(N) + 1

	Amb := models.Ambientes

	Amb2 := Amb[numeroAleatorio]

	return Amb2
}

func ChamarVariavel() {

}
