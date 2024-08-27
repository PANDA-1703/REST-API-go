package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/db"
	"rest-api/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Обновление задачи
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	var task models.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Неправильный формат данных.", http.StatusBadRequest)
		return
	}

	task.UpdatedAt = time.Now().UTC()

	_, err = db.Pool.Exec(r.Context(), "UPDATE tasks SET title = $1, description = $2, due_date = $3, updated_at = $4 WHERE id = $5", task.Title, task.Description, task.Duedate, task.UpdatedAt, id)

	if err != nil {
		http.Error(w, "Проблема на сервере", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
