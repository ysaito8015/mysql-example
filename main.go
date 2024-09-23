package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/ysaito8015/mysql-example/config"
)

func main() {
	c := config.New()
	if _, err := sql.Open("mysql", c.FormatDSN()); err != nil {
		log.Fatal(err)
	}
}
