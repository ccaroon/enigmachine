package enigma

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LetterToIdx(letter byte) int {
	idx := letter - byte('A')
	return int(idx)
}

func IdxToLetter(idx int) byte {
	return byte(idx)%26 + byte('A')
}
