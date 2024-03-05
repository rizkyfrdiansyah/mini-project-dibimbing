package admin

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, handler *AdminHandler) {
	admin := router.Group("/admin")
	{
		admin.GET("/quizzes", handler.GetQuizzes)
		admin.POST("/quiz", handler.CreateQuiz)
		admin.PUT("/quiz/:id", handler.UpdateQuiz)
		admin.DELETE("/quiz/:id", handler.DeleteQuiz)
		admin.GET("/answers/:quiz_id", handler.GetAnswers)
	}
}
