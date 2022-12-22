package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"quoteapp/db"
	"quoteapp/model"
	"quoteapp/view"

	"github.com/gin-gonic/gin"
)

type Quote struct {
	quote *model.Quote
}

func NewQuoteController(q *model.Quote) *Quote {
	return &Quote{quote: q}
}

func (q *Quote) Fetch(g *gin.Context) {
	resp, err := http.Get("https://animechan.vercel.app/api/quotes")
	if err != nil {
		view.Respond(g, http.StatusInternalServerError, nil, err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		view.Respond(g, http.StatusInternalServerError, nil, err)
		return
	}

	quote := []db.Quotes{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		view.Respond(g, http.StatusInternalServerError, nil, err)
		return
	}

	err = q.quote.Create(&quote)
	if err != nil {
		view.Respond(g, http.StatusInternalServerError, nil, err)
		return
	}

	data := map[string]bool{
		"fetch": true,
	}

	view.Respond(g, http.StatusCreated, data, nil)
}

func (q *Quote) Show(g *gin.Context) {
	quote, err := q.quote.Get()
	if err != nil {
		view.Respond(g, http.StatusInternalServerError, nil, err)
		return
	}

	data := make(map[string]string)
	data["anime"] = quote.Anime
	data["character"] = quote.Character
	data["quote"] = quote.Quote

	view.Respond(g, http.StatusOK, data, nil)
}

func (q *Quote) Count(g *gin.Context) {
	count, err := q.quote.Count()
	if err != nil {
		view.Respond(g, http.StatusInternalServerError, nil, err)
		return
	}

	data := map[string]int64{
		"count": count,
	}

	view.Respond(g, http.StatusOK, data, nil)
}

func (q *Quote) Store(g *gin.Context) {
	quote := db.Quotes{}
	err := g.BindJSON(&quote)
	if err != nil {
		view.Respond(g, http.StatusInternalServerError, nil, err)
		return
	}

	err = q.quote.Create(&quote)
	if err != nil {
		view.Respond(g, http.StatusInternalServerError, nil, err)
		return
	}

	data := map[string]bool{
		"add": true,
	}

	view.Respond(g, http.StatusCreated, data, nil)
}
