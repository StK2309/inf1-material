package main

import "fmt"

func GetInput() int {
	fmt.Println("Gib eine Zahl zwischen 1 und 10 aus")
	var n int
	fmt.Scan(&n)

	if n > 0 && n <= 10 {
		retun n 
	}
	fmt.Printf("%v ist eine gute Zahl.", n)
	return GetInput()
}

	fmt.Println("Ungültige Eingabe, versuche es noch einmal")

func main() {
	n := GetInput()
	fmt.Printf("%v, ist eine gute Zahl.", n)
}
