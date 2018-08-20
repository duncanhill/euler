package main

import (
	"fmt"
	"strconv"
)

func main() {
	largest := 0

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			test := i * j
			str := strconv.Itoa(test)

			isPalin := true
			start, end := 0, len(str)-1

			for start < end {
				isPalin = str[start] == str[end]
				if !isPalin {
					break
				}

				start++
				end--
			}

			if isPalin && test > largest {
				largest = test
			}
		}
	}

	fmt.Println(largest)
}
