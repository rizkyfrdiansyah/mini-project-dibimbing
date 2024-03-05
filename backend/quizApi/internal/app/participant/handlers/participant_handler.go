package handlers

import (
	"net/http"
	"quizApi/internal/app/admin/models"

	"github.com/gin-gonic/gin"
)

type ParticipantHandler struct {
	app *app.App
}

func NewParticipantHandler(app *app.App) *ParticipantHandler {
	return &ParticipantHandler{app: app}
}

func (h *ParticipantHandler) GetActiveQuizzes(c *gin.Context) {
	
	quizzes, err := h.app.QuizRepository.GetActiveQuizzes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quizzes)
}

func (h *ParticipantHandler) StartQuiz(c *gin.Context) {
	
	quizID := c.Param("quiz_id")
	err := h.app.QuizService.StartQuiz(quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Quiz started successfully"})
}

func (h *ParticipantHandler) SubmitQuiz(c *gin.Context) {
	
	var answers []models.Answer
	if err := c.ShouldBindJSON(&answers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.app.QuizService.SubmitQuiz(answers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Quiz submitted successfully"})
}
