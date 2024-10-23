package main

import (
	"day1/db"
	"day1/handler"
	"day1/repository"
	"day1/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db := db.ConnectDb()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"}, // Allow only from these origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                   // Allow only specific HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},        // Allow specific headers
		ExposeHeaders:    []string{"Content-Length"},                                 // Expose these headers to the frontend
		AllowCredentials: true,                                                       // Allow cookies and other credentials
		MaxAge:           12 * time.Hour,                                             // Cache preflight requests
	}))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping",
			"status":  "success",
		})
	})

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.Run(":8000")
}
