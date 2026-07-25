package enigma

type Rotor struct {
	id       string
	wiring   string
	inverse  string
	notch    int
	position int
}

func NewRotor(id string, wiring string, topLetter rune, notch rune) *Rotor {
	rotor := &Rotor{
		id:     id,
		wiring: wiring,
		notch:  LetterToIdx(notch),
	}

	inverse := make([]rune, 26)
	for idx, letter := range wiring {
		inverse[LetterToIdx(letter)] = IdxToLetter(idx)
	}
	rotor.inverse = string(inverse)

	// fmt.Printf("%s -> inverse -> %s\n", rotor.id, rotor.inverse)

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

func (rotor *Rotor) Forward(inPos int) (int, rune) {
	adjPos := (inPos + rotor.position) % 26

	outLtr := rune(rotor.wiring[adjPos])

	outIdx := (LetterToIdx(outLtr) - rotor.position) % 26
	if outIdx < 0 {
		outIdx = 26 + outIdx
	}

	return outIdx, outLtr
}

func (rotor *Rotor) Reverse(inPos int) (int, rune) {
	adjPos := (inPos + rotor.position) % 26

	outIdx := (LetterToIdx(rune(rotor.inverse[adjPos])) - rotor.position) % 26
	if outIdx < 0 {
		outIdx = 26 + outIdx
	}

	return outIdx, IdxToLetter(outIdx)
}
