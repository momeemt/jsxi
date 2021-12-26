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
	gormDB, err := gorm.Open(
		mysql.Open("root:root@tcp(database:3600)/main?charset=utf8mb4&parseTime=True&loc=Local"),
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
