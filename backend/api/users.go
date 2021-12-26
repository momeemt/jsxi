package main

import (
	"momeemt/jsxi/lib"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		Id   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
		Name string `json:"name"`
	}

	UsersResponse struct {
		Users []User `json:"users"`
	}
)

func GetUsers(c echo.Context) error {
	gormDB, sqlDB, err := lib.GetDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	defer sqlDB.Close()
	users := []User{}
	query := gormDB.Model(&User{})
	query.Find(&users)
	response := new(UsersResponse)
	response.Users = users

	return c.JSON(http.StatusOK, response)
}

func PostUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	gormDB, sqlDB, err := lib.GetDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	defer sqlDB.Close()
	gormDB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
