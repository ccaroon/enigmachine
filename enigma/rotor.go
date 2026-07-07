package enigma

import (
	"strings"
)

type Rotor struct {
	id       string
	wiring   string
	position byte
}

func NewRotor(id string, wiring string, topLetter byte) *Rotor {
	rotor := &Rotor{
		id:     id,
		wiring: wiring,
	}

	rotor.SetTopLetter(topLetter)

	return rotor
}

func GetRotor(id string) *Rotor {
	var rotor *Rotor

	switch id {
	case "0":
		rotor = NewRotor("0", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 'A')
	case "I":
		rotor = NewRotor("I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ", 'A')
	case "II":
		rotor = NewRotor("II", "AJDKSIRUXBLHWTMCQGZNPYFVOE", 'A')
	case "III":
		rotor = NewRotor("III", "BDFHJLCPRTXVZNYEIWGAKMUSQO", 'A')
	case "IV":
		rotor = NewRotor("IV", "ESOVPZJAYQUIRHXLNFTGKDCMWB", 'A')
	case "V":
		rotor = NewRotor("V", "VZBRGITYUPSDNHLXAWMJQOFECK", 'A')
	}

	return rotor
}

func (rotor *Rotor) Index(idx byte) byte {
	return rotor.wiring[idx]
}

func (rotor *Rotor) SetTopLetter(letter byte) {
	rotor.position = LetterToIdx(letter)
}

// Right to Left
func (rotor *Rotor) Forward(letter byte) byte {
	lIdx := LetterToIdx(letter)
	inPos := (lIdx + (26 - rotor.position)) % 26

	return rotor.wiring[inPos]
}

// Left to Right
func (rotor *Rotor) Reverse(letter byte) byte {
	lIdx := strings.IndexByte(rotor.wiring, letter)
	outPos := (byte(lIdx) + rotor.position) % 26

	return IdxToLetter(outPos)
}
