package database

import (
	"os"

	_ "github.com/go-sql-driver/mysql" ///import from db
	"github.com/jinzhu/gorm"
)

//MYSQLConnection ...
func MYSQLConnection() *gorm.DB {
	host := os.Getenv("DB_URI")
	db, err := gorm.Open("mysql", host)
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	return db
}
