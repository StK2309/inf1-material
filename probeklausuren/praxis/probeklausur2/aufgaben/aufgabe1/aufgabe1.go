package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// PrefixBelow10 erwartet eine Liste "list" von Zahlen und liefert
// die längste Teil-Liste, mit der "list" beginnt und die nur Zahlen < 10 enthält.
func PrefixBelow10(list []int) []int {
	var prefix []int
	for _, num := range list {
		if num < 10 {
			prefix = append(prefix, num)
		} else {
			break
		}
	}
	return prefix
}

/* var prefix []int
	for _, num := range list {
		if num < 10 {
			prefix = append(prefix, num)
		} else {
			break
		}
	}
	return prefix
} */
