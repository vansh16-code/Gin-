package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Welcome to Gin",
	})
}

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"book_id": id,
	})
}

func AddHandler(ctx *gin.Context) {
	a := ctx.Query("a")
	b := ctx.Query("b")

	numA, _ := strconv.Atoi(a)
	numB, _ := strconv.Atoi(b)

	sum := numA + numB

	ctx.JSON(http.StatusOK, gin.H{
		"sum": sum,
	})
}

func ProfileHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "user profile",
	})
}

func DashboardHandler(ctx *gin.Context) {
	user, _ := ctx.Get("userName")

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "welcome",
		"user": user,
	})
}
