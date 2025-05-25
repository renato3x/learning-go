package main

import (
	"fmt"

	"github.com/renato3x/generator"
)

func forward(origin <-chan string, destination chan string) {
  for {
    destination <- <-origin
  }
}

/* 
  Multiplexer is a pattern that allows you to combine multiple channels into one.
  This is useful when you want to listen to multiple channels and process the data
  as it comes in.
  In this example, we have two channels that are generating titles from different URLs.
  We want to combine them into one channel so we can process the titles as they come in.
*/

func merge(inputs ...<-chan string) <-chan string {
  c := make(chan string)

  for _, input := range inputs {
    go forward(input, c)
  }

  return c
}

func main() {
  c := merge(
    generator.Title("https://www.google.com", "https://www.github.com"),
    generator.Title("https://www.reddit.com", "https://www.stackoverflow.com"),
    generator.Title("https://www.renato3x.dev", "https://www.golang.org"),
  )

  fmt.Println(<-c, <-c)
  fmt.Println(<-c, <-c)
  fmt.Println(<-c, <-c)
}
