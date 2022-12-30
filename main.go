package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"PUT", "GET", "DELETE", "POST", "OPTION"}
	config.AllowHeaders = []string{"X-Auth-Key", "X-Auth-Secret", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"}
	config.AllowCredentials = true

	route.Use(cors.New(config))
}
