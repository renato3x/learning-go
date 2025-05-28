package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT * FROM users WHERE id > ?", 5)

  defer rows.Close()

  for rows.Next() {
    var user User
    rows.Scan(&user.id, &user.name)

    fmt.Println(user)
  }
}
