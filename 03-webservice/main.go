package main

import (
	"log"
	"net/http"
	"webservice/internal/database"
	"webservice/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	db := database.NewConnDB()

	userUsecase := user.NewUserUsecase(db)
	r.POST("/register", userUsecase.Register)

	r.Run(":4567")
}
