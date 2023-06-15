package strandom

import (
	"math/rand"
	"strings"
	"time"
)

var combined []byte

var (
	numbers      string = "0123456789"
	lowerLetters string = "abcdefghijklmnopqrstuvwxyz"
	upperLetters string = strings.ToUpper(lowerLetters)
)

func shuffleLetters(letters []byte) []byte {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})
	return letters
}

func randomByte(length int) []byte {
	rand.Seed(time.Now().UnixNano())
	randomByte := make([]byte, length)

	for i := 0; i < length; i++ {

		shuffledNumbers := shuffleLetters([]byte(numbers))
		shuffledLowerLetters := shuffleLetters([]byte(lowerLetters))
		shuffledUpperLetters := shuffleLetters([]byte(upperLetters))
		combined = shuffleLetters(
			append(
				append(
					shuffledNumbers,
					shuffledLowerLetters...,
				),
				shuffledUpperLetters...),
		)

		randomByte[i] = combined[rand.Intn(len(combined))]
	}

	return randomByte
}

func RandomString(length int) string {
	if length == 0 {
		length = 8
	}
	return string(randomByte(length))
}
