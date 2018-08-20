package main

import "fmt"

func main() {
	n, found := 0, false

	for !found {
		n++
		found = true

		for x := 1; x <= 20 && found; x++ {
			found = n%x == 0
		}
	}

	fmt.Println(n)
}
