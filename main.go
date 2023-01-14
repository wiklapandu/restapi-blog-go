package main

import (
	"os"

	AuthController "go-restapi/controllers/auth"
	BlogController "go-restapi/controllers/blog"
	CategoryController "go-restapi/controllers/category"
	dbHelp "go-restapi/helper/database"
	authMiddleware "go-restapi/middleware/auth"

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

	route := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"PUT", "GET", "DELETE", "POST", "OPTION"}
	config.AllowHeaders = []string{"X-Auth-Key", "X-Auth-Secret", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"}
	config.AllowCredentials = true

	route.Use(cors.New(config))

	route.POST("/register", AuthController.Create)

	route.POST("/login", AuthController.Login)

	blog := route.Group("/blog")
	{
		blog.POST("/", authMiddleware.AuthCheck(), BlogController.Postblog)
		blog.PUT("/:id", authMiddleware.AuthCheck(), BlogController.Putblog)
		blog.DELETE("/:id", authMiddleware.AuthCheck(), BlogController.Deleteblog)
		blog.GET("/:slug", BlogController.GetblogBySlug)
		blog.GET("/", BlogController.Getblog)

		category := blog.Group("/category")
		{
			category.POST("/", authMiddleware.AuthCheck(), CategoryController.Postcategory)
		}
	}
	route.Run()
}
