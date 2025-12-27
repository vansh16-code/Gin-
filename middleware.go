package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		println("Incoming Request :", ctx.Request.Method, ctx.Request.URL.Path)
		ctx.Next()
	}
	//middleware returns an function that Gin can use, that's why return type is gin.HandlerFunc
}

func CheckHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.GetHeader("X-USER")

		if user == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing user header",
			})
			ctx.Abort()
			return
		}

		ctx.Set("userName", user)

		ctx.Next()
	}
}
