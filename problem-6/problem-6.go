package main

import "fmt"

func main() {
	sumOfSquares := 0
	for x := 1; x <= 100; x++ {
		sumOfSquares += x * x
	}

	squareOfSums := 0
	for x := 1; x <= 100; x++ {
		squareOfSums += x
	}
	squareOfSums *= squareOfSums

	diff := squareOfSums - sumOfSquares
	fmt.Println(diff)
}
