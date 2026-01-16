package digits

import (
	"fmt"
	"slices"
)

// BinDigits erwartet eine Zahl n und liefert eine Liste von Ziffern
func BinDigits(n int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % 2
		result = append(result, last_digit) // result verkehrt herum
		// result = append([]int{last_digit}, result...) // Direkt umgekehrt anhängen
		n /= 2 // n = n / 2
	}

	slices.Reverse(result)

	return result
}

//

func HexDigits(n int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % 16
		result = append(result, last_digit) //result verkehrt herum
		// result = append([]int{last_digit}, result...) // Direkt umgekehrt anhängen
		n /= 2 // n = n / 2
	}

	slices.Reverse(result)

	return result
}

func Digits(n, base int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % base
		result = append(result, last_digit) // result verkehrt herum
		n /= 2                              // n = n / 2
	}

	slices.Reverse(result)

	return result
}

// Sum erwartet eine Liste vin Zahlen und berechnet deren Summe.

func Sum(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}

// ParityBit erwartet eine Anzahl n und liefert:
// 1: Fallls die Anzahl der Einsen in der Binärdarstellung von n ungerade sit.
// 0: Falls die Anzahl der Einsen in der Binärdarstellung von n gerade ist
func ParityBit(n int) int {
	// Sum(Digits(n, 2)) % 2
	return DigitSum(n, 2) % 2
}

// DigitSum berechnet de Quersumme von n (zu Basis 10)
func DigitSum(n, base int) int {
	return Sum(Digits(n, base))
}

func ExampleBinDigits() {
	fmt.Println(BinDigits(42, 2))
	fmt.Println(Digits(42, 16))
	fmt.Println(Digits(42, 10))
	fmt.Println(Digits(42, 8))

	fmt.Println(ParityBit(42))
	fmt.Println(ParityBit(42))

	// Output:
	// [1 0 1 0 1 0]
	// [2 10]
	// [4 2]
	// [5 2]
}
