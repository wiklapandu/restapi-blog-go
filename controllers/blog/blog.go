package blog

import (
	dbHelp "go-restapi/helper/database"
	blogModel "go-restapi/models/blog"

	"github.com/gin-gonic/gin"
)

func Getblog(ctx *gin.Context) {
	DB := dbHelp.DB
	var blogs blogModel.Blog
	DB.Find(&blogs)

	ctx.JSON(200, gin.H{
		"status": "success",
		"blogs":  blogs,
	})
}
