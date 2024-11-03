package server

import (
	"backend/routes"
	"backend/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartServer() {

	r := mux.NewRouter()

	ConfigureRoutes(r)

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)

	c := ConfigurationCORS()

	handler := c.Handler(r)

	s := &http.Server{
		Addr:         "localhost:8089",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}

func ConfigureRoutes(r *mux.Router) {
	for _, route := range routes.Rotas {
		var handler http.Handler = route.Handler
		r.HandleFunc(route.Path, handler.ServeHTTP).Methods(route.Method)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	utils.WriteErrorInJson(w, 404, fmt.Sprintf("Path '%s' not found", path))
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteErrorInJson(w, 405, "Method not allowed for this path")
}
