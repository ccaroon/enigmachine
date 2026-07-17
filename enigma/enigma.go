package enigma

import (
	"fmt"
	"slices"
)

type Enigma struct {
	plugboard *Plugboard
	rotors    []*Rotor
	reflector *Reflector
}

func NewEnigma(refId string, rotorIds []string, pbSpec []byte) *Enigma {
	rotors := make([]*Rotor, len(rotorIds))
	for idx, rId := range rotorIds {
		rotors[idx] = GetRotor(rId)
	}

	return &Enigma{
		plugboard: NewPlugboard(pbSpec),
		rotors:    rotors,
		reflector: GetReflector(refId),
	}
}

func (enigma *Enigma) GetRotor(id string) *Rotor {
	var rotor *Rotor

	idFunc := func(rotor *Rotor) bool {
		return rotor.Id() == id
	}

	idx := slices.IndexFunc(enigma.rotors, idFunc)
	if idx >= 0 {
		rotor = enigma.rotors[idx]
	}

	return rotor
}

func (enigma *Enigma) GetRotorByIdx(idx int) *Rotor {
	return enigma.rotors[idx]
}

func (enigma *Enigma) ConfigureRotors(startLetters string) {
	for idx, letter := range startLetters {
		enigma.rotors[idx].SetTopLetter(byte(letter))
	}
}

// func (enigma *Enigma) Step2() {
// 	// Just before every letter is enciphered,
// 	// if the top letter of any rotor *except the leftmost* is its turnover,
// 	// then that rotor and the rotor to its left step.

// }

func (enigma *Enigma) Step() {
	// TODO: Generalize to N rotors

	// Assume 3 Rotors
	// Check: Middle, then Right
	if enigma.rotors[1].AtNotch() {
		enigma.rotors[0].Step()
		enigma.rotors[1].Step()
	} else if enigma.rotors[2].AtNotch() {
		enigma.rotors[1].Step()
	}

	// Always step Right
	enigma.rotors[2].Step()
}

func (enigma *Enigma) EncipherLetter(inLetter byte) byte {
	var outLetter byte

	// ### FORWARD (right-to-left) ###
	// PLUGBOARD
	outLetter = enigma.plugboard.Map(inLetter)

	// ROTORS
	for idx := len(enigma.rotors) - 1; idx >= 0; idx-- {
		rotor := enigma.rotors[idx]
		outLetter = rotor.Forward(outLetter)
	}
	// REFLECTOR
	fmt.Println(outLetter)
	outLetter = enigma.reflector.Reflect(outLetter)

	// ### REVERSE (left-to-right) ###
	// ROTORS
	for idx := 0; idx < len(enigma.rotors); idx++ {
		rotor := enigma.rotors[idx]
		outLetter = rotor.Reverse(outLetter)
	}
	// PLUGBOARD
	outLetter = enigma.plugboard.Map(outLetter)

	return outLetter
}

func (enigma *Enigma) EncipherString(input string) string {
	var output string

	for _, letter := range input {
		enigma.Step()
		newLtr := enigma.EncipherLetter(byte(letter))
		// fmt.Println(newLtr)
		output += string(newLtr)
	}

	return output
}
