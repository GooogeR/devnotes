package main

import (
	"devnotes/router"
	"devnotes/service"
	"devnotes/storage"
)

func main() {
	// Инициализация хранилища и сервиса
	store := storage.NewMemoryStore()
	svc := service.NewService(store)

	// Инициализация роутера
	r := router.SetupRouter(svc)

	// Убираем дублирующую регистрацию маршрута /notes
	// r.POST("/notes", h.CreateNote)  // Удалить эту строку

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
