package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = map[string]string{}

func LoginHandler(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email and password are required",
		})
		return
	}

	storedPassword, exists := users[body.Email]

	if !exists || storedPassword != body.Password {
		ctx.JSON(401, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "login successfull",
	})
}

func SignupHandler(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email and pasword are required",
		})
		return
	}

	if _, exists := users[body.Email]; exists {
		ctx.JSON(400, gin.H{
			"error": "user already exists",
		})
		return
	}

	users[body.Email] = body.Password

	ctx.JSON(200, gin.H{
		"message": "user created successfully",
	})
}
