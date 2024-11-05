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
		Handler: handlers.BuscarHerois,
	},
}
