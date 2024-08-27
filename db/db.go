package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool - лобальный пул коннекта к бд
var Pool *pgxpool.Pool

// CreateTable создаёт таблицу tasks, если её не существует
func CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS tasks(
		id SERIAL PRIMARY KEY, 
		title VARCHAR(255) NOT NULL,
		description TEXT,
		due_date TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := Pool.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v\n", err)
	}
	log.Println("Таблица tasks создана/существует.")
}

// Init - нициализация подключения к бд
func Init() {
	var err error

	// Переменная окружения DATABASE_URL для подключения к бд (postgres)
	db := os.Getenv("DATABASE_URL")
	if db == "" {
		log.Println("Переменная окружения DATABASE_URL не установлена!")
		db = "postgres://panda:0000@localhost:5432/restapi?sslmode=disable"
	}

	Pool, err = pgxpool.New(context.Background(), db)
	if err != nil {
		log.Fatalf("Ошибка подключения к бд: %v\n", err)
	}

	log.Println("Подключение к бд успешно.")
}
