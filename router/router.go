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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
			return
		}

		createdUser, err := svc.Register(user.Username, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusOK, createdUser)
	})

	r.POST("/login", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user_id": user.ID})
	})

	r.POST("/notes", func(c *gin.Context) {
		var req struct {
			UserID  string `json:"user_id"`
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request"})
			return
		}

		// Теперь создаем заметку, используя данные из req
		createdNote, err := svc.CreateNote(req.UserID, req.Title, req.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to create note"})
			return
		}

		c.JSON(http.StatusOK, createdNote)
	})

	r.GET("/notes", func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Missing user ID"})
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
