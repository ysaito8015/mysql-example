package main

import (
	//"context"
	"database/sql"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func prepareDBConfigure() mysql.Config {
	return mysql.Config{
		DBName:               "example",
		User:                 "mysql",
		Passwd:               "password",
		Addr:                 "127.0.0.1:13306",
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}
}
func TestGetRegisteredDriver(t *testing.T) {
	assert.Equal(t, []string{"mysql"}, sql.Drivers())
}

func TestPingMySql(t *testing.T) {
	c := prepareDBConfigure()
	db, err := sql.Open("mysql", c.FormatDSN())

	assert.Nil(t, err)
	assert.NotNil(t, db)

	defer db.Close()

	err = db.Ping()
	assert.Nil(t, err)
}

func TestExecuteDDL(t *testing.T) {
	c := prepareDBConfigure()
	db, err := sql.Open("mysql", c.FormatDSN())

	assert.Nil(t, err)

	defer db.Close()

	// create table
	_, err = db.Exec(`create table if not exists book(isbn varchar(14), title varchar(200), price int, primary key(isbn))`)

	assert.Nil(t, err)

	// insert
	result, err := db.Exec(`insert into book(isbn, title, price) values(?, ?, ?)`, "978-4798161488", "MySQL徹底入門 第4版", 4180)

	assert.Nil(t, err)

	rowsAffected, err := result.RowsAffected()

	assert.Nil(t, err)
	assert.Equal(t, int64(1), rowsAffected)

	result, err = db.Exec(`insert into book(isbn, title, price) values(?, ?, ?)`, "978-4873116389", "実践ハイパフォーマンスMySQL 第3版", 5280)

	assert.Nil(t, err)

	rowsAffected, err = result.RowsAffected()

	assert.Nil(t, err)
	assert.Equal(t, int64(1), rowsAffected)

	result, err = db.Exec(`insert into book(isbn, title, price) values(?, ?, ?)`, "978-4798147406", "詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド", 3960)

	assert.Nil(t, err)

	// see affected rows
	// get latest id from LastInsertId
	rowsAffected, err = result.RowsAffected()

	assert.Nil(t, err)
	assert.Equal(t, int64(1), rowsAffected)

	// query row
	row := db.QueryRow(`select count(*) from book`)

	assert.Nil(t, row.Err())

	var count int
	row.Scan(&count)
	assert.Equal(t, 3, count)

	row = db.QueryRow(`select * from book where isbn = ?`, "978-4798161488")

	assert.Nil(t, row.Err())

	var isbn, name string
	var price int
	// get values from row
	row.Scan(&isbn, &name, &price)

	assert.Equal(t, "978-4798161488", isbn)
	assert.Equal(t, "MySQL徹底入門 第4版", name)
	assert.Equal(t, 4180, price)

	// query rows
	rows, err := db.Query(`select title from book where price > ? order by price desc`, 4000)

	assert.Nil(t, err)

	names := []string{}

	for rows.Next() {
		var name string

		err := rows.Scan(&name)

		assert.Nil(t, err)

		names = append(names, name)
	}

	assert.Equal(t, []string{"実践ハイパフォーマンスMySQL 第3版", "MySQL徹底入門 第4版"}, names)

	// drop table
	_, err = db.Exec(`drop table if exists book`)

	assert.Nil(t, err)
}
