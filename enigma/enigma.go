package enigma

import (
	"fmt"
	"slices"
	"strings"
)

type Enigma struct {
	plugboard *Plugboard
	rotors    []*Rotor
	reflector *Reflector
	trace     bool
}

func NewEnigma(refId rune, rotorIds []string, pbSpec []string) (*Enigma, error) {
	// Rotors
	rotors := make([]*Rotor, len(rotorIds))
	rotorMap := make(map[string]any, len(rotorIds))
	for idx, rId := range rotorIds {
		// Check for duplicate rotors
		if _, exists := rotorMap[rId]; exists {
			return nil, fmt.Errorf("Duplicate Rotors Detected: [%s @ %d]", rId, idx)
		}
		rotorMap[rId] = struct{}{}

		rotors[idx] = GetRotor(rId)
	}

	// Plugboard
	plugboard, err := NewPlugboard(pbSpec)
	if err != nil {
		return nil, err
	}

	// Reflector
	reflector, err := GetReflector(refId)
	if err != nil {
		return nil, err
	}

	return &Enigma{
		plugboard: plugboard,
		rotors:    rotors,
		reflector: reflector,
	}, nil
}

func (enigma *Enigma) ToggleTrace() {
	enigma.trace = !enigma.trace
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
		enigma.rotors[idx].SetTopLetter(letter)
	}
}

func (enigma *Enigma) Step() {
	// TODO: Generalize to N rotors

	// Assume 3 Rotors
	// Check: Middle, then Right
	if enigma.rotors[1].AtNotch() {
		enigma.printTrace("-> Step Rotor %s\n", enigma.rotors[0].Id())
		enigma.rotors[0].Step()

		enigma.printTrace("-> Step Rotor %s\n", enigma.rotors[1].Id())
		enigma.rotors[1].Step()
	} else if enigma.rotors[2].AtNotch() {
		enigma.printTrace("-> Step Rotor %s\n", enigma.rotors[1].Id())
		enigma.rotors[1].Step()
	}

	// Always step Right
	enigma.printTrace("-> Step Rotor %s\n", enigma.rotors[2].Id())
	enigma.rotors[2].Step()
}

func (enigma *Enigma) EncipherLetter(letter rune) rune {
	var inLetter rune = letter
	var outLetter rune

	// ### FORWARD (right-to-left) ###
	// PLUGBOARD
	outLetter = enigma.plugboard.Map(inLetter)
	outIdx := LetterToIdx(outLetter)
	enigma.printTrace("PB(1): %c -> %c\n", inLetter, outLetter)

	// ROTORS
	for idx := len(enigma.rotors) - 1; idx >= 0; idx-- {
		rotor := enigma.rotors[idx]
		inIdx := outIdx
		inLetter = IdxToLetter(inIdx)
		outIdx, outLetter = rotor.Forward(inIdx)
		enigma.printTrace("Rotor%s(F): %c -> %c\n", rotor.Id(), inLetter, outLetter)
	}

	// REFLECTOR
	inLetter = IdxToLetter(outIdx)
	outLetter = enigma.reflector.Reflect(inLetter)
	outIdx = LetterToIdx(outLetter)
	enigma.printTrace("Refl%c: %c -> %c\n", enigma.reflector.Id(), inLetter, outLetter)

	// ### REVERSE (left-to-right) ###
	// ROTORS
	for idx := 0; idx < len(enigma.rotors); idx++ {
		rotor := enigma.rotors[idx]
		inIdx := outIdx
		inLetter = IdxToLetter(inIdx)
		outIdx, outLetter = rotor.Reverse(inIdx)
		enigma.printTrace("Rotor%s(R): %c -> %c\n", rotor.Id(), inLetter, outLetter)
	}

	// PLUGBOARD
	inLetter = outLetter
	outLetter = enigma.plugboard.Map(inLetter)
	enigma.printTrace("PB(2): %c -> %c\n", inLetter, outLetter)

	return outLetter
}

func (enigma *Enigma) EncipherString(input string, strict bool) EnigmaOutput {
	var output strings.Builder

	for _, letter := range input {
		inLtr := rune(strings.ToUpper(string(letter))[0])
		if strings.ContainsRune(ALPHABET, inLtr) {
			enigma.Step()
			enigma.printTrace("rotors")
			newLtr := enigma.EncipherLetter(inLtr)

			output.WriteRune(newLtr)
		} else if !strict {
			output.WriteRune(letter)
		}
	}

	return EnigmaOutput(output.String())
}

func (enigma *Enigma) printTrace(what string, args ...any) {
	if enigma.trace {
		switch what {
		case "rotors":
			lWindow := IdxToLetter(enigma.rotors[0].position)
			mWindow := IdxToLetter(enigma.rotors[1].position)
			rWindow := IdxToLetter(enigma.rotors[2].position)
			fmt.Printf("%c-%c-%c\n", lWindow, mWindow, rWindow)
		default:
			fmt.Printf(what, args...)
		}
	}
}
