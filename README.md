# REST-API-go

## Оглавление
- [О проекте]($о-проекте)
- [Установка](#установка)
- [Настройка](#настройка)
- [Запуск проекта](#запуск-проекта)
- [Примеры запросов](#примеры-запросов)

## О проекте
Тестовое задание для VERBA-group.

**Разработка REST API для управления задачами (To-Do List).**

Цель:
Разработать REST API для системы управления задачами, которая позволяет пользователям создавать, просматривать, обновлять и удалять задачи.

## Установка
1. Клонируем репозиторий:
```shell
git clone https://github.com/PANDA-1703/REST-API-go.git
cd REST-API-go
```
2. Устанавливаем зависимости:
```shell
go mod tidy
```
## Настройка
1. Устанавливаем [PostgreSQL](https://www.postgresql.org/download/linux/) и создаём бд для проекта.

2. Установим `golint`, `gofmt`:
```shell
go install golang.org/x/lint/golint@latest
go install golang.org/x/tools/cmd/gofmt@latest
```

3. В файле config.json пропишем свои параметры подключения к бд или установим переменную окружения `DATABASE_URL`:
```shell
DATABASE_URL=postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable
```

## Запуск проекта
```shell
go run main.go
```
Сервер будет запущен на порту 8080 по умолчанию. Мы можем изменить порт в `config.json`.

## Примеры запросов
1. **Создать задачу (task):**
```shell
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title": "Task 1", "description": "Task 1 description", "due_date": "2024-08-27T16:32:00Z"}'
```

2. **Получение списка задач**:
```shell
curl -X GET http://localhost:8080/tasks
```

3. **Получение задачи по id**:
```shell
curl -X GET http://localhost:8080/tasks/1
```

4. **Обновление задачи по id**:
```shell
curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title": "Updated Task 1", "description": "Updated task 1 description", "due_date": "2024-09-01T16:33:00Z"}'
```

5. **Удаление задачи по id**:
```shell
curl -X DELETE http://localhost:8080/tasks/1
```


---

Это `README.md` должен помочь без проблем запустить проект. 


