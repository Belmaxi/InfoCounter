package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/counseling_statistics?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
}
