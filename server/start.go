package server

import (
	"backend/routes"
	"backend/utils"
	"log"
	"net/http"
	"time"
)

type myHandler struct{}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if !routes.ItsMethodPathValid(r) {
		utils.WriteResultError(w, 404)
		return
	}

	routes.Rotas[r.Method][r.URL.Path](w, r)
}

func StartServer() {
	s := &http.Server{
		Addr:         "localhost:8089",
		Handler:      myHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
