package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("- - - DB Connection: Success - - -")
	return DB
}
