package router

import (
	"devnotes/model"
	"devnotes/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(svc *service.Service) *gin.Engine {
	r := gin.Default()

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

	r.POST("/login", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ошибка запроса"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Вход успешен", "user_id": user.ID})
	})

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

		// Теперь создаем заметку, используя данные из req
		createdNote, err := svc.CreateNote(req.UserID, req.Title, req.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Ошибка при создании заметки"})
			return
		}

		c.JSON(http.StatusOK, createdNote)
	})

	r.GET("/notes", func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")
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

	return r
}
