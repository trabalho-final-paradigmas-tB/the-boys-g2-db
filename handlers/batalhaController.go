package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"
)

func batalhar(w http.ResponseWriter, r *http.Request) {
	var herois models.Batalha
	if err := json.NewDecoder(r.Body).Decode(&herois); err != nil {
		http.Error(w, "Erro ao receber herois"+err.Error(), http.StatusBadRequest)
		return
	}
}
