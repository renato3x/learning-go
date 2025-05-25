package main

import "time"

func isPrime(value int) bool {
	if value < 2 {
		return false
	}

	for i := 2; i < value; i++ {
		if value%i == 0 {
			return false
		}
	}

	return true
}

func primes(quantity int, c chan int) {
	start := 2

	for i := 0; i < quantity; i++ {
		for prime := start; ; prime++ {
			if isPrime(prime) {
				c <- prime
				start = prime + 1
        time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}

  close(c)
}

func main() {
  c := make(chan int, 30)

  go primes(cap(c), c)

  for prime := range c {
    println("Prime:", prime)
  }
}
