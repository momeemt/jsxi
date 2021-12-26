module momeemt/jsxi/api

go 1.16

require (
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo/v4 v4.6.1
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	gorm.io/gorm v1.22.4 // indirect
	momeemt/jsxi/lib v0.0.1
)

replace momeemt/jsxi/lib v0.0.1 => ../lib
