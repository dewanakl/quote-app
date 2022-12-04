package controller

import (
	"encoding/json"
	"net/http"
	"quoteapp/model"
	"quoteapp/repository"
)

type Quote struct {
	quote *repository.Quote
}

func NewQuoteController(q *repository.Quote) *Quote {
	return &Quote{quote: q}
}

func (q Quote) Fetch(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Fetch Quote"))
}

func (q Quote) Show(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{Data: "Show Quote"})
}

func (q Quote) Count(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Count Quote"))
}

func (q Quote) Store(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Store Quote"))
}
