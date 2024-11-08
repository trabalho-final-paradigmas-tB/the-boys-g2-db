package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"
)

func Batalhar(w http.ResponseWriter, r *http.Request) {
	var luts models.Lutadores
	if err := json.NewDecoder(r.Body).Decode(&luts); err != nil {
		http.Error(w, "Erro ao receber herois"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Herói atualizado com sucesso!"})
}
