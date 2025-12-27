package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func WelcomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Welcome to Gin",
	})
}

func SendEmailHandler(ctx *gin.Context) {

	to := ctx.Query("to")

	emailJobs <- EmailJob{
		To:      to,
		Subject: "Welcome!",
		Body:    " Testing background workers ",
	}

	ctx.JSON(200, gin.H{
		"status": "queued",
		"to":     to,
	})
}

