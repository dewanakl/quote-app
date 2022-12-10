package view

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func Respond(w http.ResponseWriter, code int, data interface{}, err error) {
	respond := Response{Code: code, Data: data, Error: nil}
	if err != nil {
		respond.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(respond)
}
