package todos

import (
	"time"
	"encoding/json"
	"net/http"
	"github.com/Caio-Nogueira/go-todo/database"
	"github.com/Caio-Nogueira/go-todo/models"
	"github.com/Caio-Nogueira/go-todo/cache"
)

func getUserFromContext(r *http.Request) models.User {
	// Get username from request context
	user_id := r.Context().Value("ID").(string)

	user := models.User{}
	database.DB.Where("ID = ?", user_id).First(&user)
	return user
}


func GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {

	user := getUserFromContext(r)
	if user.ID == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	todos := []models.Todo{}

	// Get todos from cache
	cacheKey := "todos:" + user.ID
	cached_todos, err := cache.Redis.Get(cache.Ctx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(cached_todos), &todos)
	
	} else {
		database.DB.Where("user_id = ?", user.ID).Find(&todos)
		serialized_todos, _ := json.Marshal(todos)
		// Add todos to cache
		cache.Redis.Set(cache.Ctx, cacheKey, serialized_todos, time.Minute)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(todos)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)
	
	if user.ID == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	var CreateTodo models.CreateTodo

	err := json.NewDecoder(r.Body).Decode(&CreateTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo := models.Todo{
		Title: CreateTodo.Title,
		Description: CreateTodo.Description,
		Status: 0,
		UserId: user.ID,
	}

	database.DB.Create(&todo)

	// Add todo to cache
	cacheKey := "todos:" + user.ID
	serialized_todo, _ := json.Marshal(todo)
	cache.Redis.Set(cache.Ctx, cacheKey, serialized_todo, time.Minute)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}