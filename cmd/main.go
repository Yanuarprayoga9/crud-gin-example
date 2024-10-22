package main

import (
	"day1/db"
	"day1/handler"
	"day1/repository"
	"day1/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	db := db.ConnectDb()


    // Initialize repository, service, and handler
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    // Setup Gin router
    r := gin.Default()

    // Define routes
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

    // Start the server
    r.Run(":8000")
}
