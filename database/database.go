package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDBConnection() {
	var err error
	dsn := "root:@tcp(localhost)/backend-url-shortener?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database")
}
