package auth

import (
	"fmt"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type MyClaims struct {
	*jwt.StandardClaims
	ID    uint   `json:"ID"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

var envError = godotenv.Load(".env")
var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if envError != nil {
			fmt.Printf("error: %v", envError)
		}

		var header struct {
			Token *string `header:"Authorization" binding:"required"`
		}

		if err := ctx.ShouldBindHeader(&header); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "token is required",
			})
			ctx.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(*header.Token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Unauthorized token",
			})
			ctx.Abort()
			return
		}

		theClaims := token.Claims.(*MyClaims)
		ctx.Set("user", theClaims)
		ctx.Next()
	}
}
