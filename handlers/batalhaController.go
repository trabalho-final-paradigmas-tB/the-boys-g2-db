package handlers

import (
	"backend/models"
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	VantagemMultiplicador    = 0.30
	DesvantagemMultiplicador = 0.30
	DanoMinimo               = 1
	DanoMaximo               = 50
	VidaInicial              = 150
)

var poderes = []string{"Soco", "Chute", "Raio Laser", "Força", "Batarangue", "Voar"}

func inicializarTurnos(herois []models.Heroi) []models.Turno {
	var turnos []models.Turno
	for _, heroi := range herois {
		turno := models.Turno{
			Nome:              heroi.NomeHeroi,
			Vida:              VidaInicial,
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

	var heroiVantagem, heroiDesvantagem, heroiNeutro []models.Heroi

	for _, lutador := range luts.Lutadores {
		encontrado := false
		for _, heroiNome := range Amb3.HeroisVan {
			if strings.EqualFold(heroiNome, lutador.NomeHeroi) {
				lutador.NivelForca += int(float64(lutador.NivelForca) * VantagemMultiplicador)
				heroiVantagem = append(heroiVantagem, lutador)
				encontrado = true
				break
			}
		}
		for _, heroiNome := range Amb3.HeroisDes {
			if strings.EqualFold(heroiNome, lutador.NomeHeroi) {
				lutador.NivelForca -= int(float64(lutador.NivelForca) * DesvantagemMultiplicador)
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

	resultados := batalhar(turnos, heroiVantagem, heroiDesvantagem)

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
	dano := rand.Intn(DanoMaximo-DanoMinimo+1) + DanoMinimo
	turno.Vida -= dano
	// Chamar função de evento aqui se necessário
}

func batalhar(turnos []models.Turno, heroisVantagem []models.Heroi, heroisDesvantagem []models.Heroi) [][]models.Turno {
	var resultadosPorTurno [][]models.Turno

	for turnoNum := 1; turnoNum <= 4; turnoNum++ {
		for i := range turnos {
			poderUsado := poderes[rand.Intn(len(poderes))]
			processarTurno(&turnos[i], poderUsado)
		}

		copiaTurnos := make([]models.Turno, len(turnos))
		copy(copiaTurnos, turnos)
		resultadosPorTurno = append(resultadosPorTurno, copiaTurnos)
	}

	return resultadosPorTurno
}
