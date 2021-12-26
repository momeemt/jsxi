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

	gormDB, sqlDB, _ := lib.GetDBClient()
	defer sqlDB.Close()
	gormDB.AutoMigrate(&Post{})

	e := echo.New()
	initRouting(e)
	e.Logger.Fatal(e.Start(":8000"))
}

func initRouting(e *echo.Echo) {
	e.GET("/posts", GetPosts)
	e.POST("/posts", PostPost)
	e.GET("/users", GetUsers)
	e.POST("/users", PostUser)
	e.GET("/", index)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!2")
}
