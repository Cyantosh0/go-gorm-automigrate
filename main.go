package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/Cyantosh0/go-gorm-automigrate/config"
	"github.com/Cyantosh0/go-gorm-automigrate/migration"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("- - - Connecting to Databse - - -")
	db := config.SetupDatabase()

	fmt.Println("- - - Running Migration - - -")
	migration.Migrate(db)
}
