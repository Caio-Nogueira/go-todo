package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	routers.TodoRouter(router)

	http.ListenAndServe(":8080", router)

}