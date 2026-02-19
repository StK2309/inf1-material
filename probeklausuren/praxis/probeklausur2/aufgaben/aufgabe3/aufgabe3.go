package aufgabe3

import "math"

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.
func CountSquares(list []int) int {
	if len(list) == 0 {
		return 0
	}
	count := 0
	if isSquare(list[0]) {
		count = 1
	}
	return count + CountSquares(list[1:])
}

// isSquare returns true if n is a perfect square.
func isSquare(n int) bool {
	r := int(math.Sqrt(float64(n)))
	return r*r == n
}
