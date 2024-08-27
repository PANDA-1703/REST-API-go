package handlers

import (
	"encoding/json"
	"net/http"

	"rest-api/db"
	"rest-api/models"
)

// Получение списка задач в формате JSON
func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Pool.Query(r.Context(), "SELECT id, title, description, due_date, created_at, updated_at FROM tasks")
	if err != nil {
		http.Error(w, "Проблема на сервере.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Duedate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			http.Error(w, "Проблема на сервере.", http.StatusInternalServerError)
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
