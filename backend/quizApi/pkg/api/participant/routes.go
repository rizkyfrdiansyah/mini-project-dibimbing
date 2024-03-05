package participant

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, handler *ParticipantHandler) {
	participant := router.Group("/participant")
	{
		participant.GET("/quizzes", handler.GetActiveQuizzes)
		participant.POST("/quiz/:id/start", handler.StartQuiz)
		participant.POST("/quiz/:id/submit", handler.SubmitQuiz)
	}
}
