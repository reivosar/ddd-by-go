package db

import (
	"database/sql"
	"fmt"
	"os"
)

func GetDBConnection() *sql.DB {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")
	var path string = fmt.Sprintf("%s:%s@tcp(db-container:3306)/%s?charset=utf8", user, pw, db_name)
	db, err := sql.Open("mysql", path)
	if err != nil {
		panic(err.Error())
	}
	return db
}
