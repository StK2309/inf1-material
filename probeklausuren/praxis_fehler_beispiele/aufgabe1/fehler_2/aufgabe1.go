package aufgabe1

import "strings"

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {
	if len(list) == 0 {
		return ""
	}
	longest := ""
	for _, str := range list {
		if strings.HasPrefix(str, "abc") {
			if len(str) > len(longest) {
				longest = str
			}
		}
	}
	return longest

	/* return findLongestABC(list, "")
	}
	func findLongestABC(list []string, longest string) string {
		if len(list) == 0 {
			return longest
		}
		if len(list[0]) >= 3 && list[0][:3] == "abc" {
			if len(list[0]) > len(longest) {
				longest = list[0]
			}
		}
		return findLongestABC(list[1:], longest) */
}
