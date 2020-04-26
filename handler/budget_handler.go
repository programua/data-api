package handler

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func BudgetHandler(w http.ResponseWriter, r *http.Request) {
	sr := mux.NewRouter()
	sr.HandleFunc("/list", budgetList)
}

func budgetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}