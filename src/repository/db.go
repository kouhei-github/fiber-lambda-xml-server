package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB
var err error

func init() {
	// connect DB
	user := os.Getenv("USERNAME")
	password := os.Getenv("USERPASS")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	name := os.Getenv("DATABASE")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(
		&User{},
	)
}
