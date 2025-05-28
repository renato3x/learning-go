package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func currentDatetime(w http.ResponseWriter, r *http.Request) {
  now := time.Now().Format("02/01/2006 03:04:05")
  fmt.Fprintf(w, "<h1>Now: %s</h1>", now)
}

func main() {
  http.HandleFunc("/", currentDatetime)
  log.Println("Running...")
  log.Fatal(http.ListenAndServe(":3000", nil))
}
