package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func Init(user string, pass string, schema string) (*sql.DB, error) {
	dbLink, err := sql.Open("mysql", user+":"+pass+"@"+schema)
	if err != nil {
		return nil, err
	}

	err = dbLink.Ping()
	if err != nil {
		return nil, err
	}

	database = dbLink

	return database, nil
}
