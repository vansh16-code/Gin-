package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(Logger()) //we use "Use" attches global middleware to router

	// Basic routes
	r.GET("/welcome", WelcomeHandler)
	r.GET("/books/:id", BooksHandler)
	r.GET("/add", AddHandler)

	// User routes
	api := r.Group("/user")
	api.GET("/profile", CheckHeader(), ProfileHandler)

	r.GET("/dashboard", CheckHeader(), DashboardHandler)

	// Auth routes
	r.POST("/login", LoginHandler)
	r.POST("/signup", SignupHandler)

	r.Run(":8080")
}
