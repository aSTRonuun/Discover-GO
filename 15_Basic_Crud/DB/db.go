package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	connectionString := "root:admin123@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", connectionString)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}

func CreateSchema() error {
	db, erro := sql.Open("mysql", "root:admin123@/")
	if erro != nil {
		return erro
	}
	defer db.Close()

	db.Exec("create database if not exists devbook")
	db.Exec("use devbook")
	db.Exec(`create table if not exists users (
		id integer auto_increment,
		name varchar(80),
		email varchar(80),
		PRIMARY KEY (id)
	)`)

	return nil
}
