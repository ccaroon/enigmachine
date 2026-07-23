package enigma

type Rotor struct {
	id       string
	wiring   []int
	inverse  []int
	notch    int
	position int
}

// type RotorPosition struct {
// 	Index  int
// 	Letter rune
// }

func NewRotor(id string, wiring string, topLetter rune, notch rune) *Rotor {
	rotor := &Rotor{
		id:     id,
		wiring: make([]int, 26),
		notch:  LetterToIdx(notch),
	}

	for inIdx, letter := range wiring {
		rotor.wiring[inIdx] = LetterToIdx(letter)
	}

	// Create inverse wirings for reverse/left-to-right operations
	inverse := make([]int, 26)
	for idx, outPos := range rotor.wiring {
		inverse[outPos] = idx
	}
	rotor.inverse = inverse

	rotor.SetTopLetter(topLetter)

	return rotor
}

func GetRotor(id string) *Rotor {
	var rotor *Rotor

	switch id {
	case "I":
		rotor = NewRotor("I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ", 'A', 'Q')
	case "II":
		rotor = NewRotor("II", "AJDKSIRUXBLHWTMCQGZNPYFVOE", 'A', 'E')
	case "III":
		rotor = NewRotor("III", "BDFHJLCPRTXVZNYEIWGAKMUSQO", 'A', 'V')
	case "IV":
		rotor = NewRotor("IV", "ESOVPZJAYQUIRHXLNFTGKDCMWB", 'A', 'J')
	case "V":
		rotor = NewRotor("V", "VZBRGITYUPSDNHLXAWMJQOFECK", 'A', 'Z')
	}

	return rotor
}

func (rotor *Rotor) Id() string {
	return rotor.id
}

func (rotor *Rotor) Wiring() []int {
	return rotor.wiring
}

func (rotor *Rotor) SetTopLetter(letter rune) {
	rotor.position = LetterToIdx(letter)
}

func (rotor *Rotor) GetTopLetter() rune {
	return IdxToLetter(rotor.position)
}

func (rotor *Rotor) Step() {
	rotor.position = (rotor.position + 1) % 26
}

func (rotor *Rotor) AtNotch() bool {
	atNotch := false
	if rotor.position == rotor.notch {
		atNotch = true
	}

	return atNotch
}

// Forward | RightToLeft
// func (rotor *Rotor) RightToLeft(letter rune) rune {
// 	inPos := LetterToIdx(letter)
// 	outPos := rotor.Forward(inPos)

// 	return IdxToLetter(outPos)
// }

func (rotor *Rotor) Fwd(inPos int) (int, rune) {
	adjPos := (inPos + rotor.position) % 26

	result := rotor.wiring[adjPos]

	// output_letter = self.wiring[(index + self.offset)%26]
	outLtr := IdxToLetter(result)
	// output_index = (ALPHABET.index(output_letter) - self.offset)%26
	outIdx := (result - rotor.position) % 26

	return outIdx, outLtr
}

// func (rotor *Rotor) Fwd2(inPos int) RotorPosition {
// 	adjPos := (inPos + rotor.position) % 26

// 	result := rotor.wiring[adjPos]

// 	// output_letter = self.wiring[(index + self.offset)%26]
// 	outLtr := IdxToLetter(result)
// 	// output_index = (ALPHABET.index(output_letter) - self.offset)%26
// 	outIdx := (result - rotor.position) % 26

// 	return RotorPosition{outIdx, outLtr}
// }

// func (rotor *Rotor) Forward(inPos int) int {
// 	adjPos := (inPos + rotor.position) % 26

// 	return rotor.wiring[adjPos]
// }

// Reverse | LeftToRight
func (rotor *Rotor) LeftToRight(letter rune) rune {
	inPos := LetterToIdx(letter)
	outPos := rotor.Reverse(inPos)

	return IdxToLetter(outPos)
}

// III(F): [1 3 5 7 9 11 2 15 17 19 23 21 25 13 24 4 8 22 6 0 10 12 20 18 16 14]
// -------
// III(R): [19 0 6 1 15 2 18 3 16 4 20 5 21 13 25 7 24 8 23 9 22 11 17 10 14 12]
func (rotor *Rotor) Reverse(inPos int) int {
	outPos := rotor.inverse[inPos]
	adjOut := outPos + rotor.position

	if adjOut < 0 {
		// adjOut = 26 + adjOut%26
		// E.g: 26 + -1 => 25
		adjOut = 26 + adjOut
		// fmt.Printf("\nWrap(%d): %d -> %d\n", rotor.position, inPos, adjOut)
	}

	// fmt.Printf("\n%d -> %d => %d\n", inPos, outPos, adjOut)

	return adjOut + rotor.position
}
