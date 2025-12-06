package game

import "fmt"

func Run() {
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess) {
			fmt.Println("Yeaaaahhhh, Gewonnen!!!")
			return
		} else {
			fmt.Println("Falsche Antwort, versuche es noch ein Mal.")
		}
	}
	fmt.Println("Game Over!!!")
}
