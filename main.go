package main

import (
	"devnotes/db"
	"devnotes/router"
	"devnotes/service"
	"devnotes/storage"
)

func main() {
	// Инициализация хранилища и сервиса
	store := storage.NewMemoryStore()
	svc := service.NewService(store)

	db.InitDB()

	// Инициализация роутера
	r := router.SetupRouter(svc)

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
