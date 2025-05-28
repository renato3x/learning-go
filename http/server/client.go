package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))

	switch {
	case r.Method == "GET" && id > 0:
    getUserById(w, id)
  case r.Method == "GET":
    getUsers(w)
  default:
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Sorry :(")
	}
}

func getUserById(w http.ResponseWriter, id int) {
  db, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

  var user User

  db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name)

  json, _ := json.Marshal(user)

  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, "%s", string(json))
}

func getUsers(w http.ResponseWriter) {
  db, err := sql.Open("mysql", "root:root@/gocourse")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

  result, err := db.Query("SELECT * FROM users")

  if err != nil {
    log.Fatal(err)
  }

  var users []User

  for result.Next() {
    var user User
    result.Scan(&user.Id, &user.Name)
    users = append(users, user)
  }

  json, _ := json.Marshal(users)

  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, "%s", string(json))
}
