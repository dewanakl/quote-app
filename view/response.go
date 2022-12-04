package view

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

func ErrorRespond(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response{Error: err.Error()})
}
