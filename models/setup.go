package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabse()  {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_jwt_go"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Product{})

	DB = database
}