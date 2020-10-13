package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func PrintUsers() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gotest?charset=utf8")
	defer db.Close()
	if err != nil {
		fmt.Println("sql.Open() ERR!")
		return
	}

	stmt, err := db.Prepare("select * from user limit 10")
	defer stmt.Close()
	if err != nil {
		return
	}

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		var name string
		var sex int8
		rows.Scan(&id, &name, &sex)
		fmt.Println(id, name, sex)
	}

}
