package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:admin123@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "create database if not exists goCourse")
	exec(db, "use goCourse")
	exec(db, "drop table if exists usurs")
	exec(db, `create table users (
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY (id)
	)`)
}
