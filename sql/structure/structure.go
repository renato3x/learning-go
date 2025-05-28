package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

  transaction, _ := db.Begin()
  statement, _ := db.Prepare("INSERT INTO users (id, name) VALUES (?, ?)")
  statement.Exec(200, "John Doe")
  statement.Exec(201, "Jason Momoa")
  _, e := statement.Exec(1, "Mike")

  if e != nil {
    transaction.Rollback()
    log.Fatal(e)
  }

  transaction.Commit()
}
