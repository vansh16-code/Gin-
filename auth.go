package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("supersecretkey")




func LoginHandler(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "email and password are required"})
		return
	}

	var user User
	if err := DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{"error": "invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(401, gin.H{"error": "invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body.Email,
		"role":  user.Role,
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
			"error": "Email and password are required",
		})
		return
	}

	var existingUser User
	if err := DB.Where("email = ?", body.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(400, gin.H{
			"error": "user already exists",
		})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Could not create User",
		})
		return
	}

	user := User{
		Email:    body.Email,
		Password: string(hashed),
		Role:     "user",
	}

	if err := DB.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "Could not create User",
		})
		return
	}

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

func AdminHandler(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"msg": "Welcome to Admin",
	})

}