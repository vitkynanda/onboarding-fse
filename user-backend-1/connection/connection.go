package connection

import (
	"go-api/models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func Connect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/backend1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&entity.User{})

	return DB
}