package router

import (
	"devnotes/middleware" // Добавляем middleware
	"devnotes/model"
	"devnotes/service"
	"devnotes/utils" // Добавляем утилиту для JWT
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(svc *service.Service) *gin.Engine {
	r := gin.Default()

	// Роут для регистрации
	r.POST("/register", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка запроса"})
			return
		}

		createdUser, err := svc.Register(user.Username, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка регистрации пользователя"})
			return
		}

		c.JSON(http.StatusOK, createdUser)
	})

	// Роут для логина
	r.POST("/login", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ошибка запроса"})
			return
		}

		// Генерация JWT токена
		token, err := utils.GenerateJWT(user.Username) // Генерируем токен с email (или username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Ошибка генерации токена"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Вход успешен", "token": token})
	})

	// Роут для создания заметки
	r.POST("/notes", func(c *gin.Context) {
		var req struct {
			UserID  string `json:"user_id"`
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ошибка запроса"})
			return
		}

		// Создание заметки
		createdNote, err := svc.CreateNote(req.UserID, req.Title, req.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Ошибка при создании заметки"})
			return
		}

		c.JSON(http.StatusOK, createdNote)
	})

	// Роут для получения заметок
	// Здесь мы добавляем middleware для проверки JWT
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware()) // Применяем middleware

	auth.GET("/notes", func(c *gin.Context) {
		userID := c.GetString("email") // Берём email, который мы установили в контекст
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Отсутствует user ID"})
			return
		}

		notes := []model.Note{}
		for _, note := range svc.GetNotesByUserID(userID) {
			notes = append(notes, note)
		}

		c.JSON(http.StatusOK, notes)
	})

	auth.PUT("/notes/:id", func(c *gin.Context) {
		noteID := c.Param("id")
		var req struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
			return
		}

		updatedNote, err := svc.UpdateNote(noteID, req.Title, req.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления заметки"})
			return
		}

		c.JSON(http.StatusOK, updatedNote)
	})

	auth.DELETE("/notes/:id", func(c *gin.Context) {
		noteID := c.Param("id")
		if err := svc.DeleteNote(noteID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления заметки"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Заметка удалена"})
	})

	return r
}
