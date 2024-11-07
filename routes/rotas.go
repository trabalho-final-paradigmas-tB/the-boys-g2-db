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
		Path:    "/heroi/{id}",
		Method:  "GET",
		Handler: handlers.ListarHeroiPorID,
	},
	"deletar heroi": {
		Path:    "/heroi/{id}",
		Method:  "DELETE",
		Handler: handlers.DeletarHeroi,
	},
	/*"inserir missão": {
		Path:    "/missão",
		Method:  "POST",
		Handler: handlers.inserirMissao,
	},*/
}
