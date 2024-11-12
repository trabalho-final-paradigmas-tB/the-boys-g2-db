package handlers

import (
	"backend/models"
	"math/rand"
)

func DiminuirPopularidade(turnos []models.Turno) models.Evento {
	indice := rand.Intn(len(turnos))
	turnos[indice].PopularidadeAtual -= 20
	return models.Evento{
		Nome:          "Diminuir Popularidade",
		Consequencias: "A popularidade de " + turnos[indice].Nome + " diminuiu.",
	}
}

func ChuvaDeRaios(turnos []models.Turno) models.Evento {
	consequencias := ""
	for i := range turnos {
		if turnos[i].Nome == "Raio Negro" {
			turnos[i].Vida += 30
			consequencias += turnos[i].Nome + " foi favorecido pela chuva de raios. "
		} else {
			turnos[i].Vida -= 30
			consequencias += turnos[i].Nome + " foi desfavorecido pela chuva de raios. "
		}
	}
	return models.Evento{
		Nome:          "Chuva de Raios",
		Consequencias: consequencias,
	}
}

func AparicaoTempesta(turnos []models.Turno) models.Evento {
	for i := range turnos {
		if turnos[i].Nome != "Tempesta" {
			turnos[i].PopularidadeAtual -= 15
		}
	}
	return models.Evento{
		Nome:          "Aparição de Tempesta",
		Consequencias: "A popularidade de todos os heróis exceto Tempesta diminuiu.",
	}
}

func ChegadaCompoundV(turnos []models.Turno) models.Evento {
	indice := rand.Intn(len(turnos))
	turnos[indice].Vida += 50
	return models.Evento{
		Nome:          "Chegada de Compound V",
		Consequencias: turnos[indice].Nome + " teve seu nível de força aumentado.",
	}
}

func IntervencaoCapitaoPatria(turnos []models.Turno) models.Evento {
	indice := rand.Intn(len(turnos))
	turnos[indice].Vida -= 50
	return models.Evento{
		Nome:          "Intervenção do Capitão Pátria",
		Consequencias: turnos[indice].Nome + " teve sua vida reduzida drasticamente.",
	}
}

func ManipulacaoMidiaVought(turnos []models.Turno) models.Evento {
	consequencias := ""
	for i := range turnos {
		if turnos[i].Nome == "Capitão Pátria" || turnos[i].Nome == "Maeve" || turnos[i].Nome == "Black Noir" {
			turnos[i].PopularidadeAtual += 20
			consequencias += turnos[i].Nome + " teve sua popularidade aumentada. "
		} else {
			turnos[i].PopularidadeAtual -= 20
			consequencias += turnos[i].Nome + " teve sua popularidade reduzida. "
		}
	}
	return models.Evento{
		Nome:          "Manipulação de Mídia pela Vought",
		Consequencias: consequencias,
	}
}

func ConfusaoBlackNoir(turnos []models.Turno) models.Evento {
	indice := rand.Intn(len(turnos))
	turnos[indice].Vida -= 40
	return models.Evento{
		Nome:          "Confusão Causada por Black Noir",
		Consequencias: turnos[indice].Nome + " teve sua vida reduzida devido à confusão.",
	}
}

func AtaqueTerrorista(turnos []models.Turno) models.Evento {
	consequencias := ""
	for i := range turnos {
		if turnos[i].Nome == "Capitão Pátria" || turnos[i].Nome == "Maeve" || turnos[i].Nome == "Black Noir" {
			turnos[i].Vida += 30
			consequencias += turnos[i].Nome + " foi favorecido pelo ataque terrorista. "
		} else {
			turnos[i].Vida -= 30
			consequencias += turnos[i].Nome + " foi desfavorecido pelo ataque terrorista. "
		}
	}
	return models.Evento{
		Nome:          "Ataque Terrorista",
		Consequencias: consequencias,
	}
}

func TraicaoProfundo(turnos []models.Turno) models.Evento {
	indice := rand.Intn(len(turnos))
	turnos[indice].Vida -= 25
	return models.Evento{
		Nome:          "Traição de Profundo",
		Consequencias: turnos[indice].Nome + " foi traído por Profundo e teve sua vida reduzida.",
	}
}

func EventosAleatorios(turnos []models.Turno) models.Evento {
	eventos := []func([]models.Turno) models.Evento{
		DiminuirPopularidade,
		ChuvaDeRaios,
		AparicaoTempesta,
		ChegadaCompoundV,
		IntervencaoCapitaoPatria,
		ManipulacaoMidiaVought,
		ConfusaoBlackNoir,
		AtaqueTerrorista,
		TraicaoProfundo,
	}

	evento := eventos[rand.Intn(len(eventos))]
	return evento(turnos)
}
