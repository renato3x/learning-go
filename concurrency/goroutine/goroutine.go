package main

import (
	"fmt"
	"time"
)

func talk(person, text string, quantity int) {
	for i := range quantity {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
  // talk("Alice", "Why don't you talk to me?", 3)
  // talk("John", "Cause I just can talk after you!", 1)

  /* 
    The goroutine is a lightweight thread managed by the Go runtime.
    The "go" statement starts a new goroutine. The function call will be run in a separate goroutine.
    The main function will not wait for the goroutine to finish. It will continue executing the next line of code.
    The goroutine will run concurrently with the main function.
  */
  go talk("Mary", "Hello!", 500)
  go talk("John", "Hi!", 500)
  talk("Michael", "Hey!", 10)
}
