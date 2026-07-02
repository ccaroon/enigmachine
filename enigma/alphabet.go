package enigma

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LetterToIdx(letter byte) byte {
	idx := letter - byte('A')
	return idx
}

func IdxToLetter(idx byte) byte {
	return byte(idx)%26 + byte('A')
}
