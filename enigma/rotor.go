package enigma

import (
	"fmt"
)

type Rotor struct {
	Id       string
	Wiring   string
	Inverse  string
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
	rotor.Initialize()

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

func (rotor *Rotor) Initialize() {
	// Get Inverse Wiring
	inverse := make([]byte, 26)
	for idx, letter := range rotor.Wiring {
		newIdx := LetterToIdx(byte(letter))
		newLtr := IdxToLetter(byte(idx))
		inverse[newIdx] = newLtr
	}
	rotor.Inverse = string(inverse)
}

func (rotor *Rotor) SetTopLetter(letter byte) {
	rotor.position = LetterToIdx(letter)
}

// Right to Left
func (rotor *Rotor) Forward(letter byte) byte {
	// (* [map_r_to_l wiring top_letter input_pos] is the left-hand output position
	//   - at which current would appear when current enters at right-hand input
	//   - position [input_pos] to a rotor whose wiring specification is given by
	//   - [wiring].  The orientation of the rotor is given by [top_letter],
	//   - which is the top letter appearing to the operator in the rotor's
	//   - present orientation.
	//   - requires:
	//   - - [wiring] is a valid wiring specification.
	//   - - [top_letter] is in 'A'..'Z'
	//   - - [input_pos] is in 0..25
	//     *)
	//
	// val map_r_to_l : string -> char -> int -> int
	// ----
	// "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	idx := (LetterToIdx(letter) + rotor.position) % 26

	return rotor.Wiring[idx]
}

// Left to Right
func (rotor *Rotor) Reverse(letter byte) byte {
	return 'A'
}

func (rotor *Rotor) Debug() {
	fmt.Printf("W: %s\n", rotor.Wiring)
	fmt.Printf("I: %s\n", rotor.Inverse)
}
