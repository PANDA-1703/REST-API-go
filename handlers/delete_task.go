package handlers

import (
	"net/http"
	"rest-api/db"
	"strconv"

	"github.com/gorilla/mux"
)

// DeleteTask - удаление задачи по id
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Задача не найдена.", http.StatusNotFound)
		return
	}

	_, err = db.Pool.Exec(r.Context(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Проблема на сервере.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
