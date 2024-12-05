package routes

import (
	"backend/handlers"
	"backend/models"
)

var Rotas = map[string]models.Route{
	"Adicionar heroi": {
		Path:    "/heroi",
		Method:  "POST",
		Handler: handlers.InserirHeroi,
	},
	"Listar herois": {
		Path:    "/heroi",
		Method:  "GET",
		Handler: handlers.ListarHerois,
	},
	"Listar heroi por id": {
		Path:    "/heroiid",
		Method:  "GET",
		Handler: handlers.ListarHeroiPorID,
	},
	"Listar heroi por nome": {
		Path:    "/heroinome",
		Method:  "GET",
		Handler: handlers.ListarHeroisPorNome,
	},
	"Listar heroi por status": {
		Path:    "/heroistatus",
		Method:  "GET",
		Handler: handlers.ListarHeroisPorStatus,
	},
	"Listar heroi por popularidade": {
		Path:    "/heroipopularidade",
		Method:  "GET",
		Handler: handlers.ListarHeroisPorNome,
	},
	"Deletar heroi": {
		Path:    "/heroi/{id}",
		Method:  "DELETE",
		Handler: handlers.DeletarHeroi,
	},
	"Modificar heroi": {
		Path:    "/heroi/{id}",
		Method:  "PUT",
		Handler: handlers.ModificarHeroi,
	},

	"Inserir missão": {
		Path:    "/missao",
		Method:  "POST",
		Handler: handlers.InserirMissao,
	},
	"Listar missão": {
		Path:    "/missao",
		Method:  "GET",
		Handler: handlers.ListadeMissões,
	},
	"Deletar missão": {
		Path:    "/missao/{id}",
		Method:  "DELETE",
		Handler: handlers.DeletarMissão,
	},
	"Modificar missão": {
		Path:    "/missao/{id}",
		Method:  "PUT",
		Handler: handlers.ModificarMissao,
	},
	"Resultado de missão": {
		Path:    "/missao/resultadomissao",
		Method:  "POST",
		Handler: handlers.Resultadomissão,
	},

	"Batalhar": {
		Path:    "/batalhar",
		Method:  "POST",
		Handler: handlers.ChamarBatalha,
	},

	"Inserir crime": {
		Path:    "/crimes",
		Method:  "POST",
		Handler: handlers.InserirCrime,
	},
	"Listar crimes": {
		Path:    "/crimes",
		Method:  "GET",
		Handler: handlers.ListarCrimes,
	},
	"Ocultar crime": {
		Path:    "/crimes/{id}",
		Method:  "PATCH",
		Handler: handlers.OcultarCrime,
	},
	"deletar crime": {
		Path:    "/crimes/{id}",
		Method:  "DELETE",
		Handler: handlers.DeletarCrime,
	},
	"editar crime": {
		Path:    "/crimes/{id}",
		Method:  "PUT",
		Handler: handlers.EditarCrime,
	},
}
