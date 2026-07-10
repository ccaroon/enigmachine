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
		// Identity Rotor
		rotor = NewRotor("0", ALPHABET, 'A')
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

func (rotor *Rotor) Id() string {
	return rotor.id
}

func (rotor *Rotor) Wiring() string {
	// Rotor Wiring adjusted for position/offset
	adjWiring := rotor.wiring[rotor.position:] + rotor.wiring[0:rotor.position]

	return adjWiring
}

// func (rotor *Rotor) Index(idx byte) byte {
// 	return rotor.wiring[idx]
// }

func (rotor *Rotor) SetTopLetter(letter byte) {
	rotor.position = LetterToIdx(letter)
}

func (rotor *Rotor) GetTopLetter() byte {
	return IdxToLetter(rotor.position)
}

// Right to Left
func (rotor *Rotor) Forward(letter byte) byte {
	lIdx := LetterToIdx(letter)
	inPos := (lIdx + (26 + rotor.position)) % 26

	return rotor.wiring[inPos]
}

// Left to Right
//
// -	             ABCDEFGHIJKLMNOPQRSTUVWXYZ
// -	   AJDKSIRUXBLHWTMCQGZNPYFVOE
func (rotor *Rotor) Reverse(letter byte) byte {
	lIdx := strings.IndexByte(rotor.wiring, letter)
	outPos := (byte(lIdx) + (26 - rotor.position)) % 26

	// fmt.Printf("\n%c) (%d + %d) %% 26 => %d\n", letter, lIdx, (26 - rotor.position), outPos)

	outLetter := IdxToLetter(outPos)
	// fmt.Printf("%d -> %c\n", outPos, outLetter)

	return outLetter
}

// ---
