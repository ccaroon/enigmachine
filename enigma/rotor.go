package enigma

import (
	"strings"
)

type Rotor struct {
	Id     string
	Wiring string
	// Inverse  string
	position byte
}

var (
	presetRotors [6]Rotor = [6]Rotor{
		NewRotor("0", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 'A'),
		NewRotor("I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ", 'A'),
		NewRotor("II", "AJDKSIRUXBLHWTMCQGZNPYFVOE", 'A'),
		NewRotor("III", "BDFHJLCPRTXVZNYEIWGAKMUSQO", 'A'),
		NewRotor("IV", "ESOVPZJAYQUIRHXLNFTGKDCMWB", 'A'),
		NewRotor("V", "VZBRGITYUPSDNHLXAWMJQOFECK", 'A'),
	}
)

func NewRotor(id string, wiring string, topLetter byte) Rotor {
	rotor := Rotor{
		Id:     id,
		Wiring: wiring,
	}

	rotor.SetTopLetter(topLetter)
	// rotor.Initialize()

	return rotor
}

func GetRotor(Id string) Rotor {
	var rotor Rotor

	switch Id {
	case "0":
		rotor = presetRotors[0]
	case "I":
		rotor = presetRotors[1]
	case "II":
		rotor = presetRotors[2]
	case "III":
		rotor = presetRotors[3]
	case "IV":
		rotor = presetRotors[4]
	case "V":
		rotor = presetRotors[5]
	}

	return rotor
}

// func (rotor *Rotor) Initialize() {
// 	// Get Inverse Wiring
// 	inverse := make([]byte, 26)
// 	for idx, letter := range rotor.Wiring {
// 		newIdx := LetterToIdx(byte(letter))
// 		newLtr := IdxToLetter(byte(idx))
// 		inverse[newIdx] = newLtr
// 	}
// 	rotor.Inverse = string(inverse)
// }

func (rotor *Rotor) SetTopLetter(letter byte) {
	rotor.position = LetterToIdx(letter)
}

// Right to Left
func (rotor *Rotor) Forward(letter byte) byte {
	pos := (LetterToIdx(letter) + rotor.position) % 26
	wLetter := rotor.Wiring[pos]

	idx := ((26 - rotor.position) + LetterToIdx(wLetter)) % 26
	return IdxToLetter(idx)
}

// Left to Right
func (rotor *Rotor) Reverse(letter byte) byte {
	pos := (LetterToIdx(letter) + rotor.position) % 26
	wLetter := IdxToLetter(pos)

	idx := strings.IndexByte(rotor.Wiring, wLetter)
	return IdxToLetter(byte(idx))

}
