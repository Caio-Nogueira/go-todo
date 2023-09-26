package api


import (
	"github.com/gorilla/mux"
	"github.com/Caio-Nogueira/go-todo/api/todos"
)

func TodoRouter(router *mux.Router) {
	router.HandleFunc("/todos", todos.GetAllTodosHandler).Methods("GET")
}
