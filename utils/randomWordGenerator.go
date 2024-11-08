package utils

import (
	"math/rand"
	"time"
)

func GenerateWord() (string, error) {
	lines := readDict()
	rand.Seed(time.Now().UnixNano())

	numLines := len(lines)
	randomLine1 := rand.Intn(numLines)
	randomLine2 := rand.Intn(numLines)

	return lines[randomLine1] + "-" + lines[randomLine2], nil

}
