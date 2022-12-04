package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"quoteapp/db"
	"quoteapp/model"
	"quoteapp/view"
)

type Quote struct {
	quote *model.Quote
}

func NewQuoteController(q *model.Quote) *Quote {
	return &Quote{quote: q}
}

func (q *Quote) ErrorRespond(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(view.Response{Error: err.Error()})
}

func (q *Quote) Fetch(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://animechan.vercel.app/api/quotes")
	if err != nil {
		q.ErrorRespond(w, err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		q.ErrorRespond(w, err)
		return
	}

	quote := []db.Quotes{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		q.ErrorRespond(w, err)
		return
	}

	err = q.quote.Insert(&quote)
	if err != nil {
		q.ErrorRespond(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: "Fetch Quote Ok"})
}

func (q *Quote) Show(w http.ResponseWriter, r *http.Request) {
	quote, err := q.quote.Get()
	if err != nil {
		q.ErrorRespond(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: quote})
}

func (q *Quote) Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: "Count Quote"})
}

func (q *Quote) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: "Store Quote"})
}
