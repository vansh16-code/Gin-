package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found (using system env)")
	}

	ConnectDatabase()

	r := gin.Default()
	r.Use(Logger())
	r.Use(cors.Default())


	r.GET("/welcome", WelcomeHandler)
	r.POST("/login", LoginHandler)
	r.POST("/signup", SignupHandler)

	r.GET("/me", AuthMiddleware(), protectedRoute)
	r.GET("/admin", AuthMiddleware(), AdminOnly(), AdminHandler)

	StartEmailWorker()
	r.GET("/send-email", SendEmailHandler)
	r.Run(":8080")
}
