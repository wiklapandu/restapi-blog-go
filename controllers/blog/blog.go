package blog

import (
	"encoding/json"
	dbHelp "go-restapi/helper/database"
	blogModel "go-restapi/models/blog"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type MyClaims struct {
	*jwt.StandardClaims
	ID    uint   `json:"ID"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func Getblog(ctx *gin.Context) {
	var blogs blogModel.Blog
	var DB = dbHelp.DB
	DB.Find(&blogs)

	ctx.JSON(200, gin.H{
		"status": "success",
		"blogs":  blogs,
	})
}

func Postblog(ctx *gin.Context) {
	data, _ := ctx.Get("user")
	var user MyClaims
	theJson, _ := json.Marshal(data)
	json.Unmarshal(theJson, &user)

	var inputPost struct {
		Title string `json:"title" form:"title" binding:"required"`
		Desc  string `json:"desc" form:"desc" binding:"required"`
		Cat   []uint `json:"cat" form:"cat" binding:"required"`
	}

	var DB = dbHelp.DB

	if err := ctx.Bind(&inputPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "errors",
			"errors": err.Error(),
		})
		ctx.Abort()
		return
	}

	newBlog := blogModel.Blog{
		Title:  inputPost.Title,
		Slug:   slug.Make(inputPost.Title),
		Desc:   inputPost.Desc,
		Author: user.ID,
	}

	if result := DB.Create(&newBlog); result.RowsAffected <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed create blog",
		})
		ctx.Abort()
		return
	}

	for _, cat := range inputPost.Cat {
		DB.Create(&blogModel.BlogCats{
			Blog:     newBlog.ID,
			Category: cat,
		})
	}

	blog := struct {
		blogModel.Blog
		Cat []uint `json:"categories"`
	}{newBlog, inputPost.Cat}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"blog":   blog,
	})
}
