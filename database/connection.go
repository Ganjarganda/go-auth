package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:@/yt_go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
}
