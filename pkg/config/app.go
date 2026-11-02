package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	_ = godotenv.Load()
	dsn := os.Getenv("BOOKSTORE_DSN")
	if dsn == "" {
		dsn = "adarsh:adarsh@/bookstore?charset=utf8&parseTime=True&loc=Local"
	}
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to MySQL: %v", err))
	}
	DB = db
	return DB
}
