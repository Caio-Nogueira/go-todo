package routers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []todo{
	{ID: "1", Title: "pay phone bill", Status: "active"},
	{ID: "2", Title: "pay water bill", Status: "active"},
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("%v", todos)))
}

func todoRouter(router *mux.Router) {
	router.HandleFunc("/todos", getTodosHandler).Methods("GET")
}
