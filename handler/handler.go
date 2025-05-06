package handler

import (
	"devnotes/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.service.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(200, user)
}

func (h *Handler) CreateNote(c *gin.Context) {
	var req struct {
		UserID  string `json:"user_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	note, err := h.service.CreateNote(req.UserID, req.Title, req.Content)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(200, note)
}
