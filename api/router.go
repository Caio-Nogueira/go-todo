package api


import (
	"github.com/gorilla/mux"
	"github.com/Caio-Nogueira/go-todo/api/todos"
	"github.com/Caio-Nogueira/go-todo/api/auth"
	"github.com/Caio-Nogueira/go-todo/middlewares"
)

func TodoRouter(router *mux.Router) {
	router.HandleFunc("/todos", middlewares.AuthMiddleware(todos.GetAllTodosHandler)).Methods("GET")
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")
}


