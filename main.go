package main

import (
	"net/http"
	"F:\assignments\go-todo\routers\todoRouter.go"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	routers.TodoRouter(router)

	http.ListenAndServe(":8080", router)

}