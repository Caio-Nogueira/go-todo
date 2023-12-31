package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Caio-Nogueira/go-todo/api"
	"github.com/Caio-Nogueira/go-todo/database"
	"github.com/Caio-Nogueira/go-todo/cache"
)

func main() {
	cache.ConnectRedis()
	database.Connect()
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	api.TodoRouter(apiRouter)
	http.ListenAndServe(":8080", router)
}