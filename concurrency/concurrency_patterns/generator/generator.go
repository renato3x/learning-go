package generator

import (
	"io"
	"net/http"
	"regexp"
)

/* 
  Generator pattern
  - A generator is a function that returns a channel and can be used to generate values.
  - It is a way to produce values in a lazy manner, meaning that the values are generated on demand.
  - This is useful when you want to generate a large number of values without having to store them all in memory.
  - The generator function can be called multiple times, and each time it will return a new channel that can be used to receive the generated values.
*/

func Title(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			response, _ := http.Get(url)
			html, _ := io.ReadAll(response.Body)
			regex, _ := regexp.Compile("<title[^>]*>([^<]+)</title")
			c <- regex.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}
