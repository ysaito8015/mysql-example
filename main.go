package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	c := mysql.Config{
		DBName:               "example",
		User:                 "mysql",
		Passwd:               "password",
		Addr:                 "127.0.0.1:13306",
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		AllowNativePasswords: true,
	}
	if _, err := sql.Open("mysql", c.FormatDSN()); err != nil {
		log.Fatal(err)
	}
}
