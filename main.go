package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(Logger()) //we use "Use" attches global middleware to router


	r.GET("/welcome", WelcomeHandler)
	r.GET("/books/:id", BooksHandler)
	r.GET("/add", AddHandler)
	api := r.Group("/user")
	api.GET("/profile", CheckHeader(), ProfileHandler)
	r.GET("/dashboard", CheckHeader(), DashboardHandler)
	r.POST("/login", LoginHandler)
	r.POST("/signup", SignupHandler)
	r.GET("/me", AuthMiddleware(),protectedRoute)

	r.Run(":8080")
}
