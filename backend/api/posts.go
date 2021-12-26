package main

import (
	"momeemt/jsxi/lib"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Post struct {
		Id          int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
		Name        string `json:"name"`
		Code        string `json:"code"`
		Description string `json:"description"`
		UserId      int    `json:"user_id"`
	}

	Response struct {
		Posts []Post `json:"posts"`
	}
)

func GetPosts(c echo.Context) error {
	gormDB, sqlDB, err := lib.GetDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	defer sqlDB.Close()
	posts := []Post{}
	query := gormDB.Model(&Post{})
	query.Find(&posts)
	response := new(Response)
	response.Posts = posts

	return c.JSON(http.StatusOK, response)
}

func PostPosts(c echo.Context) error {
	post := Post{}
	if err := c.Bind(&post); err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	gormDB, sqlDB, err := lib.GetDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	defer sqlDB.Close()
	gormDB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}
