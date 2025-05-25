package main

import (
	"time"

	"github.com/renato3x/generator"
)

func fastest(url1, url2, url3 string) string {
  c1 := generator.Title(url1)
  c2 := generator.Title(url2)
  c3 := generator.Title(url3)

  select {
  case title := <-c1:
    return title
  case title := <-c2:
    return title
  case title := <-c3:
    return title
  case <-time.After(2 * time.Second):
    return "Timeout"
  }
}

func main() {
  winner := fastest(
    "https://www.github.com",
    "https://www.google.com",
    "https://www.reddit.com",
  )

  println(winner)
}
