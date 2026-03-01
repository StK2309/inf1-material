package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {
	var result []int
	for _, v := range l1 {
		if !contains(l2, v) {
			result = append(result, v)
		}
	}
	for _, v := range l2 {
		if !contains(l1, v) {
			result = append(result, v)
		}
	}
	return result
}

func contains(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

/* var result []int
	for _, v := range l1 {
		if !contains(l2, v) {
			result = append(result, v)
		}
	}
	for _, v := range l2 {
		if !contains(l1, v) {
			result = append(result, v)

		}

	}
	return result
}

func contains(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
} */
