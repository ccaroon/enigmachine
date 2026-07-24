package enigma

type Rotor struct {
	id       string
	wiring   string
	inverse  string
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
		wiring: wiring,
		notch:  LetterToIdx(notch),
	}

	// fmt.Println("\nWiring: ", wiring)

	inverse := make([]rune, 26)
	for idx, letter := range wiring {
		inverse[LetterToIdx(letter)] = IdxToLetter(idx)
	}
	rotor.inverse = string(inverse)
	// fmt.Println("\nInverse: ", rotor.inverse)

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

func (rotor *Rotor) Wiring() string {
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

func (rotor *Rotor) Forward(inPos int) (int, rune) {
	adjPos := (inPos + rotor.position) % 26

	// output_letter = self.wiring[(index + self.offset)%26]
	// outLtr := rune(rotor.wiring[adjPos])

	// output_index = (ALPHABET.index(output_letter) - self.offset)%26
	// outIdx := (LetterToIdx(outLtr) - rotor.position) % 26
	outIdx := (LetterToIdx(rune(rotor.wiring[adjPos])) - rotor.position) % 26
	if outIdx < 0 {
		outIdx = 26 + outIdx
	}

	return outIdx, IdxToLetter(outIdx)
}

// III(F): [1 3 5 7 9 11 2 15 17 19 23 21 25 13 24 4 8 22 6 0 10 12 20 18 16 14]
// -------
// III(R): [19 0 6 1 15 2 18 3 16 4 20 5 21 13 25 7 24 8 23 9 22 11 17 10 14 12]
func (rotor *Rotor) Reverse(inPos int) (int, rune) {
	adjPos := (inPos + rotor.position) % 26

	// output_letter = self.inverse[(index + self.offset)%26]
	// outLtr := rune(rotor.inverse[adjPos])

	// output_index = (ALPHABET.index(output_letter) - self.offset)%26
	// craig := LetterToIdx(outLtr) - rotor.position
	// if craig < 0 {
	// 	fmt.Printf("\nDEBUG => NEGATIVE IDX: %d | %d\n", craig, craig%26)
	// }
	// outIdx := (LetterToIdx(outLtr) - rotor.position) % 26
	outIdx := (LetterToIdx(rune(rotor.inverse[adjPos])) - rotor.position) % 26
	if outIdx < 0 {
		outIdx = 26 + outIdx
	}

	// fmt.Printf("\n%d: %d -> [%d|%d]\n", inPos, adjPos, outLtr, outIdx)

	return outIdx, IdxToLetter(outIdx)
}
