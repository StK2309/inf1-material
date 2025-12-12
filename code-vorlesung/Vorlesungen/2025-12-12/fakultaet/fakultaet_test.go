package main

import "fmt"

func ExampleFactorial() {
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(4))
	fmt.Println(Factorial(3))
	fmt.Println(Factorial(2))
	fmt.Println(Factorial(1))

}

// Output:
// 1
// 1
// 2
// 6
// 24
// 120
