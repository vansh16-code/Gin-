package main

import (
	"net/http"
	"strings"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]string{}
var jwtSecret = []byte("supersecretkey")

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
	if !ok {
		ctx.JSON(401, gin.H{"error": "invalid token payload"})
		ctx.Abort()
		return
	}

		ctx.Set("email",email)
		ctx.Next()
	}
}



func LoginHandler(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	storedPassword, exists := users[body.Email]
	if !exists {
		ctx.JSON(401, gin.H{"error": "invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(body.Password)); err != nil {
		ctx.JSON(401, gin.H{"error": "invalid email or password"})
		return
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "could not create token"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "login successful",
		"token":   tokenString,
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


	hashed , err := bcrypt.GenerateFromPassword([]byte(body.Password),10)
	if err != nil {
		ctx.JSON(500,gin.H{
			"error": "Could not create User",
		})
		return
	}

	users[body.Email] = string(hashed)



	ctx.JSON(200, gin.H{
		"message": "user created successfully",
	})

}


func protectedRoute(ctx *gin.Context){
	
	email,_ := ctx.Get("email")

	ctx.JSON(200,gin.H{
		"user" : email,
	})
}