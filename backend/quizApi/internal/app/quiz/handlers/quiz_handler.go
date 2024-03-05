package handlers

import (
	"net/http"
	"quizApi/internal/app/admin/models"

	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	app *app.App
}

func NewQuizHandler(app *app.App) *QuizHandler {
	return &QuizHandler{app: app}
}

func (h *QuizHandler) GetActiveQuizzes(c *gin.Context) {
	var quizzes []models.Quiz
	if err := h.app.DB.Where("active = ?", true).Find(&quizzes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quizzes)
}

func (h *QuizHandler) StartQuiz(c *gin.Context) {
	
	quizID := c.Param("quiz_id")

	
	var quiz models.Quiz
	if err := h.app.DB.Where("id = ?", quizID).Preload("Questions").First(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quiz"})
		return
	}

	
	if len(quiz.Questions) > 0 {
		firstQuestion := quiz.Questions[0]
		c.JSON(http.StatusOK, firstQuestion)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No questions found for the quiz"})
	}
}


func (h *QuizHandler) SubmitQuiz(c *gin.Context) {
	
	var answers []models.Answer
	if err := c.ShouldBindJSON(&answers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	for _, answer := range answers {
		
		if answer.QuestionID == 0 || answer.OptionID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answer"})
			return
		}
	}

	
	if err := h.app.DB.Create(&answers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit quiz"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quiz submitted successfully"})
}

