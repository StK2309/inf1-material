package aufgabe4

// ElementProducts erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils das Produkt der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt,
// soll dieses Element ins Ergebnis übernommen werden.
func ElementProducts(l1, l2 []int) []int {
	max := len(l1)
	if len(l2) > max {
		max = len(l2)
	}
	result := make([]int, max)
	for i := 0; i < max; i++ {
		if i < len(l1) && i < len(l2) {
			result[i] = l1[i] * l2[i]
		} else if i >= len(l1) {
			result[i] = l2[i]
		} else if i >= len(l2) {
			result[i] = l1[i]
		}
	}
	return result
}
