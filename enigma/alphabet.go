package enigma

import (
	"math/rand/v2"
	"strings"
)

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LetterToIdx(letter rune) int {
	idx := letter - 'A'
	return int(idx)
}

func IdxToLetter(idx int) rune {
	letter := rune(idx)%26 + 'A'
	return letter
}

func RandomLetters(count int) string {
	var letters []string = make([]string, count)

	for i := range count {
		idx := rand.IntN(len(ALPHABET))
		letters[i] = string(ALPHABET[idx])
	}

	return strings.Join(letters, "")
}
