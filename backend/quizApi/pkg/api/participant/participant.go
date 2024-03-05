package participant

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ParticipantHandler struct {
	app *app.App
}

func NewParticipantHandler(app *app.App) *ParticipantHandler {
	return &ParticipantHandler{app: app}
}

func (h *ParticipantHandler) GetActiveQuizzes(c *gin.Context) {
	
	activeQuizzes, err := h.app.QuizService.GetActiveQuizzes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activeQuizzes)
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
	
	quizID := c.Param("quiz_id")
	var answers []Answer
	if err := c.ShouldBindJSON(&answers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.app.QuizService.SubmitQuiz(quizID, answers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Quiz submitted successfully"})
}
