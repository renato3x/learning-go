package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		panic(err)
	}

	defer db.Close()
  
  statement, _ := db.Prepare("INSERT INTO users (name) VALUES (?)")
  statement.Exec("Mary")
  statement.Exec("John")

  result, _ := statement.Exec("Mike")

  id, _ := result.LastInsertId()
  fmt.Println(id)

  affected_rows, _ := result.RowsAffected()
  fmt.Println(affected_rows)
}
