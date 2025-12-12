package main

import "fmt"

func ExamplePrimes() {
	fmt.Println(Primes(10))
	//fmt.Println(Primes(1:3)) // Slicing! von Stelle 1 (hier gleich 3), bis nicht merh Stelle 3 (hier 5)

	// Output:
	// [2 3 5 7]
}
