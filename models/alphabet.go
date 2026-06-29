package models

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LetterToIdx(letter byte) byte {
	return letter - byte('A')
}

func IdxToLetter(idx byte) byte {
	return idx%26 + byte('A')
}
