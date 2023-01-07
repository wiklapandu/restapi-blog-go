package main

import (
	"net/http"
	"os"

	"go-restapi/helper"
	dbHelp "go-restapi/helper/database"
	userModel "go-restapi/models/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	*jwt.StandardClaims
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func main() {
	dbHelp.RunDB()
	DB := dbHelp.DB

	route := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"PUT", "GET", "DELETE", "POST", "OPTION"}
	config.AllowHeaders = []string{"X-Auth-Key", "X-Auth-Secret", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"}
	config.AllowCredentials = true

	route.Use(cors.New(config))

	route.POST("/register", func(ctx *gin.Context) {
		var inputPost struct {
			Name           string `json:"name" form:"name" binding:"required"`
			Email          string `json:"email" form:"email" binding:"required,email"`
			Password       string `json:"password" form:"password" binding:"required"`
			RepeatPassword string `json:"repeat_password" form:"repeat_password" binding:"required"`
		}

		if err := ctx.Bind(&inputPost); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "errors",
				"errors": err.Error(),
			})

			ctx.Abort()
			return
		}

		if inputPost.Password != inputPost.RepeatPassword {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "password not matches with repeat password",
			})

			ctx.Abort()
			return
		}
		password, _ := helper.HashPassword(inputPost.Password)
		user := userModel.User{Name: inputPost.Name, Email: inputPost.Email, Password: password}

		result := DB.Create(&user)

		if result.RowsAffected <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "user is fail to create",
			})

			ctx.Abort()
			return
		}

		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), MyClaims{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})

		token, err := sign.SignedString(SecretKey)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "fail create token",
			})

			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"token":  token,
		})
	})

	route.POST("/login", func(ctx *gin.Context) {
		var inputPost struct {
			Email    string `json:"email" form:"email" binding:"required,email"`
			Password string `json:"password" form:"password" binding:"required"`
		}

		if err := ctx.BindJSON(&inputPost); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "errors",
				"errors": err.Error(),
			})
			ctx.Abort()
			return
		}

		var user = userModel.User{Email: inputPost.Email}

		result := DB.First(&user)

		if result.RowsAffected <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "user not register",
			})
			ctx.Abort()
			return
		}

		if helper.CheckPassword(inputPost.Password, user.Password) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  "fail",
				"message": "incorrect email and password",
			})
			ctx.Abort()
			return
		}

		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), MyClaims{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})

		token, err := sign.SignedString(SecretKey)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "fail",
				"message": "failed create token",
			})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"status": "success",
			"token":  token,
		})
	})

	route.Run()
}
