package routes

import (
	"backend/handlers"
	"net/http"
)

var Rotas = map[string]map[string]http.HandlerFunc{

	// exemplo de como colocar nas rotas

	"GET": {
		"/heroi": handlers.InserirHeroi,
	},
}

func ItsMethodPathValid(r *http.Request) bool {
	if methodRoutes, exists := Rotas[r.Method]; exists {
		if _, pathExists := methodRoutes[r.URL.Path]; pathExists {
			return true
		}
	}
	return false
}
