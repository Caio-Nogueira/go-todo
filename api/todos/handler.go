package todos

import (
	// "encoding/json"
	"encoding/json"
	"net/http"
	"github.com/Caio-Nogueira/go-todo/database"
	"github.com/Caio-Nogueira/go-todo/models"
)



func GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	// Get username from request context
	username := r.Context().Value("username").(string)

	user := models.User{}
	database.DB.Where("username = ?", username).First(&user)

	if user.Username == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	todos := []models.Todo{}
	database.DB.Where("user_id = ?", user.ID).Find(&todos)
	
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(todos)

}