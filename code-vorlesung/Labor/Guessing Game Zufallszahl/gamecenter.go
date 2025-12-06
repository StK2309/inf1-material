package gamecenter

import "fmt"

func Go() {
	for n := 0; n < 3; n++ {
		guess := Read()
		if NumberCorrect(guess) {
			fmt.Println("Gewonnen!!!")
			return
		} else {
			fmt.Println("Falsche Antwort, versuche es noch ein Mal.")
		}
	}
	fmt.Println("Gameover!!!")
}
