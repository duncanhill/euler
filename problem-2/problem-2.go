package main

import "fmt"

func main() {
	limit := 4000000
	a, b := 0, 1
	sum := 0

	for b <= limit {
		a, b = b, a+b

		if b%2 == 0 {
			sum += b
		}
	}

	fmt.Println(sum)
}
