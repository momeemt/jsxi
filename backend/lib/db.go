package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenMySQLConnection() *gorm.DB {
	gormDB, err := gorm.Open(
		mysql.Open("root:root@tcp(database:3600)/main?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database.")
	}
	return gormDB
}
