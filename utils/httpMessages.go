package utils

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Code   int
	Error  string
	Reason string
}

var HttpMessages = map[int]Message{
	404: {
		Code:   404,
		Error:  "Path não encontrado",
		Reason: "",
	},
	500: {
		Code:   500,
		Error:  "Internal Server Error",
		Reason: "",
	},
}

func WriteErrorInJson(res http.ResponseWriter, code int, reason string) {
	message := HttpMessages[code]
	message.Reason = reason
	jon, _ := json.Marshal(message)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write(jon)
}
