package blog

import "gorm.io/gorm"

type Blog struct {
	*gorm.Model
	Thumbnail string `json:"img" gorm:"column:img"`
	Title     string `json:"title" gorm:"column:title"`
	Slug      string `json:"slug" gorm:"column:slug"`
	Desc      string `json:"desc" gorm:"column:desc"`
	Author    uint   `json:"user_id" gorm:"column:user_id"`
}

type CatMeta struct {
	Desc string `json:"desc"`
}

type Category struct {
	*gorm.Model
	Name     string  `json:"name" gorm:"column:name"`
	MetaData CatMeta `json:"metadata" gorm:"embedded;column:metadata;type:longtext"`
}

type BlogCats struct {
	*gorm.Model
	Blog     uint `json:"blog_id" gorm:"column:blog_id"`
	Category uint `json:"cat_id" gorm:"column:cat_id"`
}
