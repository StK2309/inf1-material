package main

import "fmt"

func fizzbuzz_v1() {
	// i := 0
	// fmt.Println(i)
	// i++

	for i := 0; i < 42; i++ {

		// if i%5 !=0 && i%7 !=0 {
		// 	fmt.Println(i)
		// }else if i%5 == 0 {
		// 	fmt.Println("FIIZZ!!!")
		// }else if i%7 == 0 {
		// 	fmt.Println("BUZZ!!!")
		// }else if ... {
		// 	fmt.Println("FIZZBUZZ!!!")
		// }

		if i%5 == 0 && i%7 == 0 {
			fmt.Println("FIZZBUZZ!!!") //durch 7 und 5 teilbar, dann Ausgabe FIZZBUZZ
		} else if i%5 == 0 {
			fmt.Println("FIZZ!!!") //nur durch 5 teilbar, dann Ausgabe FIZZ
		} else if i%7 == 0 {
			fmt.Println("BUZZ!!!") //nur durch 7 teilbar, dann Ausgabe BUZZ
		} else {
			fmt.Println(i) //nicht durch 7 und 5 teibar, dann Zahl ausgeben
		}
	}
}

func main() {
	fizzbuzz_v1()
}
