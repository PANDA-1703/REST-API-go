package main

import (
	"log"
	"net/http"
	"rest-api/db"
	"rest-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Подключение к бд
	db.Init()
	defer db.Pool.Close()

	// Создадим таблицу tasks (если не создана)
	db.CreateTable()

	// Создаём роутер
	r := mux.NewRouter()

	// Настройка маршрутов
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", handlers.DeleteTask).Methods("DELETE")

	// Запуск серва
	log.Println("Сервер запущен на порту 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
