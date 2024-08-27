package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"rest-api/db"
	"rest-api/handlers"

	"github.com/gorilla/mux"
)

// Config содержит параметры подключения к бд
type Config struct {
	DatabaseURL string `json:"database_url"`
	Port        string `json:"port"`
}

func main() {
	// Чтение из файла конфигурации
	var configFile string
	flag.StringVar(&configFile, "config", "config.json", "Путь до файла конфига")
	flag.Parse()

	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Ошибка чтения файла конфига: %v\n", err)
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Ошибка парсинга файла конфига: %v\n", err)
	}

	os.Setenv("DATABASE_URL", config.DatabaseURL)

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
