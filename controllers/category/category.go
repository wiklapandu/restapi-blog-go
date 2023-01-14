package category

import (
	dbHelp "go-restapi/helper/database"
	CategoryModel "go-restapi/models/blog"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Postcategory(ctx *gin.Context) {
	var input struct {
		Name string `json:"name" form:"name" binding:"required"`
		Desc string `json:"desc" form:"desc"`
	}
	DB := dbHelp.DB

	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "errors",
			"errors": err.Error(),
		})
		ctx.Abort()
		return
	}

	category := CategoryModel.Category{
		Name: input.Name,
		MetaData: CategoryModel.CatMeta{
			Desc: input.Desc,
		},
	}

	if result := DB.Create(&category); result.RowsAffected <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "failed create category",
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":   "success",
		"category": category,
	})
}
