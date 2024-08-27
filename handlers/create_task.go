package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"rest-api/db"
	"rest-api/models"
)

// Создание задачи
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Неправильный формат данных.", http.StatusBadRequest)
		return
	}

	// Устанавливаем значения created, updated и закидываем в базу
	task.CreatedAt = time.Now().UTC()
	task.UpdatedAt = task.CreatedAt

	err = db.Pool.QueryRow(
		r.Context(),
		"INSERT INTO tasks (title, description, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		task.Title, task.Description, task.Duedate, task.CreatedAt, task.UpdatedAt).Scan(&task.ID) // Если всё норм, возвращаем задачу с текущим ID

	if err != nil {
		http.Error(w, "Проблема на сервере.", http.StatusInternalServerError)
		log.Println("Ошибка выполнения запроса INSERT:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
