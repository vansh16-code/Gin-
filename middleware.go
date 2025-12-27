package main

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"fmt"
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


func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(401, gin.H{"error": "missing or invalid token"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(401, gin.H{
				"error": "Could not read claims",
			})
			ctx.Abort()
			return 
		}

		email, ok := claims["email"].(string)
		role, ok := claims["role"].(string)
	if !ok {
		ctx.JSON(401, gin.H{"error": "invalid token payload"})
		ctx.Abort()
		return
	}

		ctx.Set("email",email)
		ctx.Set("role",role)
		ctx.Next()
	}
}

func AdminOnly()gin.HandlerFunc{
	return  func(ctx *gin.Context) {

		role,_:= ctx.Get("role")

		if role != "admin" {
			ctx.JSON(403,gin.H{
				"error" : "admin access only ",
			})
			ctx.Abort()
			return 
		}
		ctx.Next()
	}
}
