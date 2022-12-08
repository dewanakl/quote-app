package view

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func respond(w http.ResponseWriter, code int, data Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func SuccessRespond(w http.ResponseWriter, data interface{}) {
	respond(w, http.StatusOK, Response{Data: data})
}

func ErrorRespond(w http.ResponseWriter, code int, err error) {
	respond(w, code, Response{Error: err.Error()})
}
