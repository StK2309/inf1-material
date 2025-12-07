package gamecenter

import "fmt"

func Read() int {
	var n int
	fmt.Print("Rate eine Zahl.")
	fmt.Scan(&n)
	return n
}
