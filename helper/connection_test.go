package helper

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestOpenConnection(t *testing.T) {
	mysqlDb:= fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", 
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"))
	_, err := gorm.Open(mysql.Open(mysqlDb), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}