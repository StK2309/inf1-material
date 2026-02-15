package aufgabe2

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {
	result := []int{}

	for i := 0; i < len(list); i++ {
		x := list[i]
		if x > m && x < n {
			result = append(result, x)
		}
	}
	return result
}
