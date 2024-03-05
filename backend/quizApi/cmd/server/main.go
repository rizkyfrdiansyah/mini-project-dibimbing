package main

import (
	"log"
	"os"
	"quizApi/internal/app"
	"quizApi/internal/config"
	"quizApi/internal/handlers"
	"quizApi/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := config.NewConfig()

	db, err := repository.NewDatabase(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	application := app.NewApp(db, config)

	router := gin.Default()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			authHandler := handlers.NewAuthHandler(application)
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		admin := api.Group("/admin")
		{
			adminHandler := handlers.NewAdminHandler(application)
			admin.GET("/quizzes", adminHandler.GetQuizzes)
			admin.POST("/quiz", adminHandler.CreateQuiz)
			admin.PUT("/quiz/:id", adminHandler.UpdateQuiz)
			admin.DELETE("/quiz/:id", adminHandler.DeleteQuiz)
			admin.GET("/answers/:quiz_id", adminHandler.GetAnswers)
		}

		participant := api.Group("/participant")
		{
			participantHandler := handlers.NewParticipantHandler(application)
			participant.GET("/quizzes", participantHandler.GetActiveQuizzes)
			participant.POST("/quiz/:id/start", participantHandler.StartQuiz)
			participant.POST("/quiz/:id/submit", participantHandler.SubmitQuiz)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}