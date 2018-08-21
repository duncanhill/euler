package main

import (
	"fmt"
	"math"
)

func main() {
	sum, next := 1, 2
	for numDivisors(sum) <= 500 {
		sum += next
		next++
	}

	// Found the one with over 500 divisors.
	fmt.Println(sum)
}

func numDivisors(n int) int {
	count := 0
	// Only need to go up to the square root, because we can count in both directions.
	// e.g. 100 - (1, 100), (2, 50), etc...
	for x := 1; float64(x) <= math.Sqrt(float64(n)); x++ {
		if n%x == 0 {
			// If they're equal only count once.
			// e.g. (10, 10).
			if n/x == x {
				count++
			} else {
				count += 2
			}
		}
	}

	fmt.Printf("(Number, Divisors): (%d, %d)\n", n, count)
	return count
}
