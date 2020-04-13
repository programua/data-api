package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeRouter() *mux.Router {
	r := mux.NewRouter()
	r.Methods("GET", "POST")
	r.Host("http://localhost:8080").Subrouter()

	r.HandleFunc("/top", topHandler)
	return r
}

func topHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}