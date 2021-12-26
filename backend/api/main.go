package main

import (
	"fmt"
	"momeemt/jsxi/lib"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type (
	User struct {
		Id   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
		Name string `json:"name"`
	}
)

func GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
	}
	return config
}

func main() {
	godotEnvErr := godotenv.Load(".env")
	if godotEnvErr != nil {
		panic("Error loading .env file")
	}

	config := GetConnect()
	fmt.Println(config)
	// url := config.AuthCodeURL(os.Getenv("OAUTH_STATE"), oauth2.AccessTypeOffline)

	e := echo.New()
	initRouting(e)
	e.Logger.Fatal(e.Start(":8000"))
}

func initRouting(e *echo.Echo) {
	e.GET("/posts", GetPosts)
	e.GET("/", index)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}

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
