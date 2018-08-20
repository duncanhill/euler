package main

import "fmt"

func main() {
	p := NewPrimes()

	fmt.Println(p.At(10000)) // 10001st prime.
}

type Primes struct {
	primes []int
	facts  map[int][]int
	next   int
}

func NewPrimes() *Primes {
	p := new(Primes)
	p.primes = make([]int, 1)
	p.facts = make(map[int][]int)

	// Initialize with knowledge of the first prime: 2.
	p.primes[0] = 2
	p.facts[4] = []int{2}
	p.next = 3

	return p
}

func (p *Primes) At(index int) int {
	if index >= len(p.primes) { // Not found.
		// The following is an incremental Sieve of Eratosthenes.
		for len(p.primes) <= index {
			considering := p.next
			facts := p.facts[considering]

			if facts != nil { // Not a prime; redistribute facts.
				for _, fact := range facts {
					new := considering + fact

					if p.facts[new] != nil {
						p.facts[new] = append(p.facts[new], fact)
					} else {
						p.facts[new] = []int{fact}
					}
				}
			} else { // Prime; append the prime.
				p.primes = append(p.primes, considering)
				p.facts[considering*considering] = []int{considering}
			}

			p.next++
		}
	}

	return p.primes[index]
}
