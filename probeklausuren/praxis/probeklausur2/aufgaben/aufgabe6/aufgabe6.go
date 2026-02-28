package aufgabe6

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// DuplicateSinglets erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste, bei der alle Elemente,
// die in list nur einmal vorkommen, verdoppelt sind,
// also zwei Mal hintereinander stehen.
// Elemente, die schon in list mehrfach vorkommen, sollen wie sie sind
// ins Ergebnis übertragen werden.
func DuplicateSinglets(list []int) []int {
	result := []int{}
	// TODO
	counts := make(map[int]int)
	for _, num := range list {
		counts[num]++
	}
	for _, num := range list {
		if counts[num] == 1 {
			result = append(result, num, num)
		} else {
			result = append(result, num)
		}
	}
	return result
}
