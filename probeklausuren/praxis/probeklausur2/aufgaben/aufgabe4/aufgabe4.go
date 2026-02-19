package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	result := []int{}
	max := len(l1)
	if len(l2) > max {
		max = len(l2)
	}

	for i := 0; i < max; i++ {
		var v1, v2 int = 0, 0 // Default values for empty lists
		if i < len(l1) {
			v1 = l1[i]
		} else if len(l1) > 0 {
			v1 = l1[len(l1)-1] // Use last element of l1 if i exceeds its length
		}

		if i < len(l2) {
			v2 = l2[i]
		} else if len(l2) > 0 {
			v2 = l2[len(l2)-1] // Use last element of l2 if i exceeds its length
		}

		result = append(result, v1+v2)
	}

	return result
}
