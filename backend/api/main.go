package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotEnvErr := godotenv.Load(".env")
	if godotEnvErr != nil {
		panic("Error loading .env file")
	}
	googleClientID := os.Getenv("GoogleClientID")
	googleClientSecret := os.Getenv("GoogleClientSecret")
	fmt.Println(googleClientID, googleClientSecret)

	e := echo.New()
	e.GET("/", index)
	e.Logger.Fatal(e.Start(":8000"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}
