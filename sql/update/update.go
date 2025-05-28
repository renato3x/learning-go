package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "root:root@/gocourse")
  
  if err != nil {
    panic(err)
  }
  
  defer db.Close()

  stmt, _ := db.Prepare("UPDATE users SET name = ? WHERE id = ?")

  stmt.Exec("Uóxiton Istive", 1)
  stmt.Exec("Sheristone Uasleska", 2)
}
