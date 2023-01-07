package database

import (
	"log"
	"os"

	blogModel "go-restapi/models/blog"
	userModel "go-restapi/models/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func RunDB() {
	var err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	database.AutoMigrate(&blogModel.Blog{}, &blogModel.Category{}, &blogModel.BlogCats{})
	database.AutoMigrate(&userModel.User{})
	DB = database
}
