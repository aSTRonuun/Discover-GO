package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin123@/goCourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into users(name) values(?)")
	stmt.Exec("Mary")
	stmt.Exec("Joao")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	lines, _ := res.RowsAffected()
	fmt.Println(lines)

}
