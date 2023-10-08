package config

import (
	"account-selling/internal/entity"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
	mysqlDb := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", 
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"))
connection, err := gorm.Open(mysql.Open(mysqlDb), &gorm.Config{})

if err != nil {
panic("Cannot connect to the database")
}

DB = connection
autoMigration(connection)
}

func autoMigration(con *gorm.DB){

	con.AutoMigrate(&entity.User{})
	con.AutoMigrate(&entity.UserData{})
}