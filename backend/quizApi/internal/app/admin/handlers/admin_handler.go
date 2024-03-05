package handlers

import (
	"net/http"

	"quizApi/internal/app/admin/models"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	app *app.App
}

func NewAdminHandler(app *app.App) *AdminHandler {
	return &AdminHandler{app: app}
}

func (h *AdminHandler) CreateQuiz(c *gin.Context) {
	var quiz models.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.app.DB.Create(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, quiz)
}

func (h *AdminHandler) GetAnswers(c *gin.Context) {
	quizID := c.Param("quiz_id")
	var answers []models.Answer
	if err := h.app.DB.Where("quiz_id = ?", quizID).Find(&answers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answers)
}
