package aufgabe1

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {

	longestLen := 0
	longestPos := -1

	for pos, val := range list {
		currentLen := len(val)
		if currentLen >= 3 && val[:3] == "abc" {
			if currentLen > longestLen {
				longestLen = currentLen
				longestPos = pos
			}
		}
	}
	if longestPos != -1 {
		return list[longestPos]
	}

	return ""
}

/* 	longest := ""

for _, val := range list {
	if strings.HasPrefix(val, "abc") {
		if len(val) > len(longest) {
			longest = val
		}
	}
}

return longest */
//}
