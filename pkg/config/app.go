package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB is the globally shared GORM connection for the app.
var DB *gorm.DB

// Connect opens the MySQL connection and stores it in DB.
// DSN is read from the BOOKSTORE_DSN env var, falling back to the local
// bookstore database on the default address.
func Connect() *gorm.DB {
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
