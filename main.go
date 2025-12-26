package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


func Logger()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		println("Incoming Request :" , ctx.Request.Method,ctx.Request.URL.Path)
		ctx.Next()
	}
//middleware returns an function  that Gin can use , that's why return type is gin.HandlerFunc
}

func CheckHeader()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		user := ctx.GetHeader("X-USER")

		if user == ""{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error" : "Missing user header",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}

}


func main() {
	r := gin.Default()
	r.Use(Logger()) //we use "Use" attches global middleware to router

	r.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Welcome to Gin",
		})
	})

	r.GET("/books/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"book_id": id,
		})
	})

	r.GET("/add", func(ctx *gin.Context) {
		a := ctx.Query("a")
		b := ctx.Query("b")

		numA, _ := strconv.Atoi(a)
		numB, _ := strconv.Atoi(b)

		sum := numA + numB

		ctx.JSON(http.StatusOK, gin.H{
			"sum": sum,
		})
	})

	api := r.Group("/user")

	api.GET("/profile",CheckHeader(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "user profile",
		})
	})


	r.POST("/login",func(ctx *gin.Context) {
		var body struct {
			Email string `json:"email" binding:"required" `
			Password string `json:"password" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&body); err!= nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error": "email and password are required ",
			})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{
			"message" : "Login data received",
			"email": body.Email,
		})
	})

	r.Run(":8080")
}
