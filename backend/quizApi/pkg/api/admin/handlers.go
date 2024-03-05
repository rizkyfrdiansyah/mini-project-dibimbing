package admin

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

func (h *AdminHandler) GetQuizzes(c *gin.Context) {
	var quizzes []models.Quiz
	if err := h.app.DB.Find(&quizzes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quizzes)
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

func (h *AdminHandler) UpdateQuiz(c *gin.Context) {
	var quiz models.Quiz
	id := c.Param("id")
	if err := h.app.DB.First(&quiz, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.app.DB.Save(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quiz)
}

func (h *AdminHandler) DeleteQuiz(c *gin.Context) {
	var quiz models.Quiz
	id := c.Param("id")
	if err := h.app.DB.First(&quiz, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	if err := h.app.DB.Delete(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quiz deleted successfully"})
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
