package todos

import (
	"encoding/json"
	"net/http"
	"github.com/Caio-Nogueira/go-todo/database"
	"github.com/Caio-Nogueira/go-todo/models"
)



func GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	db := database.DB
	w.Header().Set("Content-Type", "application/json")
	db.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}