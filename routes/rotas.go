package routes

import (
	"backend/handlers"
	"backend/models"
)

var Rotas = map[string]models.Route{

	// exemplo de como colocar nas rotas

	"inserir heroi": {
		Path:    "/heroi",
		Method:  "POST",
		Handler: handlers.InserirHeroi,
	},
	"listar herois": {
		Path:    "/heroi",
		Method:  "GET",
		Handler: handlers.ListarHerois,
	},
	"listar heroi por id": {
		Path:    "/heroiid",
		Method:  "GET",
		Handler: handlers.ListarHeroiPorID,
	},
	"deletar heroi": {
		Path:    "/heroi/{id}",
		Method:  "DELETE",
		Handler: handlers.DeletarHeroi,
	},
	"inserir missão": {
		Path:    "/missão",
		Method:  "POST",
		Handler: handlers.InserirMissao,
	},
	"batalhar": {
		Path:    "/batalhar",
		Method:  "POST",
		Handler: handlers.ChamarBatalha,
	},
	"listademissão": {
		Path:    "/missao",
		Method:  "GET",
		Handler: handlers.ListadeMissões,
	},
	"deletarmissão": {
		Path:    "/missao/{id}",
		Method:  "DELETE",
		Handler: handlers.DeletarMissão,
	},
	"modificarmissão": {
		Path:    "/missao/{id}",
		Method:  "PUT",
		Handler: handlers.ModificarMissao,
	},
	"inserir crime": {
		Path:    "/crime",
		Method:  "POST",
		Handler: handlers.InserirCrime,
	},
	"listar crimes": {
		Path:    "/crimes",
		Method:  "GET",
		Handler: handlers.ListarCrimes,
	},
	"ocultar crime": {
		Path:    "/crimes/{id}/ocultar",
		Method:  "PATCH",
		Handler: handlers.OcultarCrime,
	},
}
