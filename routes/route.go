package routes

import (
	"encoding/json"
	"net/http"
	"quoteapp/controller"
	"quoteapp/model"
)

type Route struct {
	mux   *http.ServeMux
	quote *controller.Quote
}

func (r *Route) get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.Response{Error: "Method is not allowed!"})
		}
	})
}

func (r *Route) post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.Response{Error: "Method is not allowed!"})
		}
	})
}

func (r *Route) Run() *http.ServeMux {
	return r.mux
}

func NewRoute(q *controller.Quote) *Route {
	route := &Route{
		mux:   http.NewServeMux(),
		quote: q,
	}

	route.routes()
	return route
}

func (r *Route) routes() {
	r.mux.Handle("/fetch", r.get(http.HandlerFunc(r.quote.Fetch)))
	r.mux.Handle("/get", r.get(http.HandlerFunc(r.quote.Show)))
	r.mux.Handle("/count", r.get(http.HandlerFunc(r.quote.Count)))
	r.mux.Handle("/add", r.post(http.HandlerFunc(r.quote.Store)))
}
