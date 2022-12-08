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

func (q *Quote) Fetch(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://animechan.vercel.app/api/quotes")
	if err != nil {
		view.ErrorRespond(w, http.StatusInternalServerError, err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		view.ErrorRespond(w, http.StatusInternalServerError, err)
		return
	}

	quote := []db.Quotes{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		view.ErrorRespond(w, http.StatusInternalServerError, err)
		return
	}

	err = q.quote.Create(&quote)
	if err != nil {
		view.ErrorRespond(w, http.StatusInternalServerError, err)
		return
	}

	data := map[string]bool{
		"fetch": true,
	}

	view.SuccessRespond(w, data)
}

func (q *Quote) Show(w http.ResponseWriter, r *http.Request) {
	quote, err := q.quote.Get()
	if err != nil {
		view.ErrorRespond(w, http.StatusInternalServerError, err)
		return
	}

	data := make(map[string]string)
	data["anime"] = quote.Anime
	data["character"] = quote.Character
	data["quote"] = quote.Quote

	view.SuccessRespond(w, data)
}

func (q *Quote) Count(w http.ResponseWriter, r *http.Request) {
	count, err := q.quote.Count()
	if err != nil {
		view.ErrorRespond(w, http.StatusInternalServerError, err)
		return
	}

	data := map[string]int64{
		"count": count,
	}

	view.SuccessRespond(w, data)
}

func (q *Quote) Store(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		view.ErrorRespond(w, http.StatusInternalServerError, err)
		return
	}

	quote := db.Quotes{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		view.ErrorRespond(w, http.StatusInternalServerError, err)
		return
	}

	err = q.quote.Create(&quote)
	if err != nil {
		view.ErrorRespond(w, http.StatusInternalServerError, err)
		return
	}

	data := map[string]bool{
		"add": true,
	}

	view.SuccessRespond(w, data)
}
