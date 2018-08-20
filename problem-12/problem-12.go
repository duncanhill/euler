package main

import "fmt"

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
	for x := 1; x <= n; x++ {
		if n%x == 0 {
			count++
		}
	}

	fmt.Printf("(Number, Divisors): (%d, %d)\n", n, count)
	return count
}
