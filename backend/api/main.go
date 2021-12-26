package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func getConnect() *oauth2.Config {
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
	googleClientID := os.Getenv("GoogleClientID")
	googleClientSecret := os.Getenv("GoogleClientSecret")
	fmt.Println(googleClientID, googleClientSecret)

	config := getConnect()
	fmt.Println(config)
	// url := config.AuthCodeURL(os.Getenv("OAUTH_STATE"), oauth2.AccessTypeOffline)

	e := echo.New()
	e.GET("/", index)
	e.Logger.Fatal(e.Start(":8000"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}
