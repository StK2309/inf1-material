package aufgabe3

// CountX erwartet eine Liste von Zahlen sowie eine Zahl x.
// CountX liefert die Anzahl der Vorkommen von x in der Liste.
func CountX(list []int, x int) int {
	count := 0
	for _, num := range list {
		if num == x {
			count++
		}
	}
	return count
}
