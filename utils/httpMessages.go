package utils

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Code  int
	Error string
}

var HttpMessages = map[int]Message{
	404: {
		Code:  404,
		Error: "Path não encontrado",
	},
}

func WriteResultError(w http.ResponseWriter, code int) {
	jon, _ := json.Marshal(HttpMessages[code])
	w.Write(jon)
}
