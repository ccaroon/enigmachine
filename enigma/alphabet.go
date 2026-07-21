package enigma

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LetterToIdx(letter rune) int {
	idx := letter - 'A'
	return int(idx)
}

func IdxToLetter(idx int) rune {
	letter := rune(idx)%26 + 'A'
	return letter
}
