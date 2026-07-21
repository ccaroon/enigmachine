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

func (enigma *Enigma) EncipherLetter(letter byte) byte {
	var inLetter byte = letter
	var outLetter byte

	// ### FORWARD (right-to-left) ###
	// PLUGBOARD
	outLetter = enigma.plugboard.Map(inLetter)
	fmt.Printf("\nPB(1): %c -> %c\n", inLetter, outLetter)

	// ROTORS
	for idx := len(enigma.rotors) - 1; idx >= 0; idx-- {
		rotor := enigma.rotors[idx]
		inLetter = outLetter
		outLetter = rotor.Forward(inLetter)
		fmt.Printf("Rotor%s(F): %c -> %c\n", rotor.Id(), inLetter, outLetter)
	}
	// REFLECTOR
	inLetter = outLetter
	outLetter = enigma.reflector.Reflect(inLetter)
	fmt.Printf("Refl%s: %c -> %c\n", enigma.reflector.Id(), inLetter, outLetter)

	// ### REVERSE (left-to-right) ###
	// ROTORS
	for idx := 0; idx < len(enigma.rotors); idx++ {
		rotor := enigma.rotors[idx]
		inLetter = outLetter
		outLetter = rotor.Reverse(inLetter)
		fmt.Printf("Rotor%s(R): %c -> %c\n", rotor.Id(), inLetter, outLetter)
	}
	// PLUGBOARD
	inLetter = outLetter
	outLetter = enigma.plugboard.Map(inLetter)
	fmt.Printf("PB(2): %c -> %c\n", inLetter, outLetter)

	return outLetter
}

func (enigma *Enigma) EncipherString(input string) string {
	var output string

	for _, letter := range input {
		enigma.Step()
		newLtr := enigma.EncipherLetter(byte(letter))
		fmt.Printf("%c => %c\n", letter, newLtr)
		output += string(newLtr)
	}

	return output
}
