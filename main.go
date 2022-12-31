package main

import (
	"log"
	"net/http"
	"os"

	blogModel "go-restapi/models/blog"
	userModel "go-restapi/models/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyClaims struct {
	*jwt.StandardClaims
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

var DB *gorm.DB

func main() {
	var err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	database.AutoMigrate(&blogModel.Blog{}, &blogModel.Category{}, &blogModel.BlogCats{})
	database.AutoMigrate(&userModel.User{})

	DB = database

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

		user := userModel.User{Name: inputPost.Name, Email: inputPost.Email}

		sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), MyClaims{
			ID:    123,
			Name:  inputPost.Name,
			Email: inputPost.Email,
		})

		token, err := sign.SignedString(SecretKey)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "unauthorized token",
			})

			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"token":  token,
		})
	})

	route.Run()
}
