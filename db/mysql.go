package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(username, password, address, dbName string) error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, address, dbName)
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	DB = db

	return nil
}

func GetDB() *sql.DB {
	return DB
}
