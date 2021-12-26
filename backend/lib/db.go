import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func db() {
	gormDB, err := gorm.Open(
			mysql.Open("root:root@tcp(database:3600)/main?charset=utf8mb4&parseTime=True&loc=Local"), 
			&gorm.Config{}}
	)
}
