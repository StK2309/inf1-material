package aufgabe2

import "sort"

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {

	result := []int{}

	for _, s := range list {
		if s > m && s < n {
			result = append(result, s)
		}

	}
	sort.Ints(result)
	return result
}
