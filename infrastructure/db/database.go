package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBConnection() *gorm.DB {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")
	port_number := os.Getenv("MYSQL_PORT_NUMBER")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/dbname?charset=utf8mb4&parseTime=True&loc=Local", user, pw, db_name, port_number)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
