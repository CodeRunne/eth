package database

import (
	_ "database/sql"
	"fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var Db *gorm.DB

func Connect(data map[string]string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", data["user"], data["pass"], data["host"], data["port"], data["name"])
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("Could'nt establish connection with database")
	}
	Db = db	
}

func GetDB() *gorm.DB {
	return Db
}