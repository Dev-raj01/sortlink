package utils

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"time"
)

func GenerateWord() (string, error) {
	file, err := os.Open("./utils/words_alpha.txt")
	if err != nil {
		return "", errors.New("Error opening file: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Set the seed for random number generation
	rand.Seed(time.Now().UnixNano())

	// Choose two random line numbers
	numLines := len(lines)
	randomLine1 := rand.Intn(numLines)
	randomLine2 := rand.Intn(numLines)

	return lines[randomLine1] + "-" + lines[randomLine2], nil

}
