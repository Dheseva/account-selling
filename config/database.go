package config

import (
	modelitems "account-selling/models/items"
	modeltransact "account-selling/models/transaction"
	modelsuser "account-selling/models/user"
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
AutoMigration(connection)
}

func AutoMigration(con *gorm.DB){

	con.AutoMigrate(&modelsuser.User{})
	con.AutoMigrate(&modelsuser.UserData{})

	con.AutoMigrate(&modelitems.Items{})
	
	con.AutoMigrate(&modeltransact.Transaction{})
}