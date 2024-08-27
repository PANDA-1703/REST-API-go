package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/db"
	"rest-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

// Забираем конкретную задачу по id
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Задача не найдена.", http.StatusNotFound)
		return
	}

	var task models.Task
	err = db.Pool.QueryRow(r.Context(), "SELECT id, title, description, due_date, created_at, updated_at FROM tasks WHERE id = 1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Duedate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		http.Error(w, "Проблема на сервере.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
