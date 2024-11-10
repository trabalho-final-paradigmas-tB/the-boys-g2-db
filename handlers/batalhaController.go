package handlers

import (
	"backend/models"
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func inicializarTurnos(herois []models.Heroi) []models.Turno {
	var turnos []models.Turno
	for _, heroi := range herois {
		turno := models.Turno{
			Nome:              heroi.NomeHeroi,
			Vida:              500,
			PoderUsado:        "",
			PopularidadeAtual: heroi.Popularidade,
		}
		turnos = append(turnos, turno)
	}
	return turnos
}

func ChamarBatalha(w http.ResponseWriter, r *http.Request) {
	var luts models.Lutadores
	if err := json.NewDecoder(r.Body).Decode(&luts); err != nil {
		http.Error(w, "Erro ao receber heróis: "+err.Error(), http.StatusBadRequest)
		return
	}

	Amb3 := ChamarAmbiente()

	var heroiVantegem, heroiDesvantagem, heroiNeutro []models.Heroi

	for _, lutador := range luts.Lutadores {
		encontrado := false
		for _, heroiNome := range Amb3.HeroisVan {
			if strings.EqualFold(heroiNome, lutador.NomeHeroi) {
				lutador.NivelForca += int(float64(lutador.NivelForca) * 0.30)
				heroiVantegem = append(heroiVantegem, lutador)
				encontrado = true
				break
			}
		}
		for _, heroiNome := range Amb3.HeroisDes {
			if strings.EqualFold(heroiNome, lutador.NomeHeroi) {
				lutador.NivelForca -= int(float64(lutador.NivelForca) * 0.30)
				heroiDesvantagem = append(heroiDesvantagem, lutador)
				encontrado = true
				break
			}
		}
		if !encontrado {
			heroiNeutro = append(heroiNeutro, lutador)
		}
	}

	turnos := inicializarTurnos(luts.Lutadores)

	resultados := batalhar(turnos, heroiVantegem, heroiDesvantagem)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultados)
}

func ChamarAmbiente() models.Local {
	rand.Seed(time.Now().UnixNano())
	N := 2
	numeroAleatorio := rand.Intn(N) + 1

	Amb := models.Ambientes

	Amb2 := Amb[numeroAleatorio]

	return Amb2
}

func processarTurno(turno *models.Turno, poderUsado string) {
	turno.PoderUsado = poderUsado
	dano := rand.Intn(50) + 1
	turno.Vida -= dano
	/*turno.PopularidadeAtual -= 5 AQUI NO CHAMAR A FUNÇÃO DO EVENTO*/
}

func batalhar(turnos []models.Turno, heroisVantagem []models.Heroi, heroisDesvantagem []models.Heroi) [][]models.Turno {
	var resultadosPorTurno [][]models.Turno

	for turnoNum := 1; turnoNum <= 4; turnoNum++ {
		for i := range turnos {
			poderUsado := "Ataque"
			processarTurno(&turnos[i], poderUsado)
		}

		copiaTurnos := make([]models.Turno, len(turnos))
		copy(copiaTurnos, turnos)
		resultadosPorTurno = append(resultadosPorTurno, copiaTurnos)
	}

	return resultadosPorTurno
}
