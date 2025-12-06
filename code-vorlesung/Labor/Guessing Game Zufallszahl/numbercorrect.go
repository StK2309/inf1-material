package gamecenter

import (
	"math/rand"
	"time"
)

var secret int

func init() {
	rand.Seed(time.Now().UnixNano())
	secret = rand.Intn(100) + 1
}

func NumberCorrect(g int) bool {
	return g == secret
}
