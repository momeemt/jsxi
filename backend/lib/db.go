package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBClient() (*gorm.DB, *sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 3 * time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbname)
	fmt.Println(connection)
	gormDB, err := gorm.Open(
		mysql.Open(connection),
		&gorm.Config{
			CreateBatchSize: 1000,
			Logger:          gormLogger,
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("error occured when connecting database\n%s", err)
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, nil, err
	}
	return gormDB, sqlDB, nil
}
