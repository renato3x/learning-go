package main

import "fmt"

func showApproved(approved ...string) {
	for index, appr := range approved {
		fmt.Printf("%d) %s\n", index+1, appr)
	}
}

func main() {
  approved  := []string{"John", "Renato", "Leo", "Kezia"}
  showApproved(approved...)
}
