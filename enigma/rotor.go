package enigma

type Rotor struct {
	id       string
	wiring   string
	inverse  string
	notch    byte
	position byte
}

func NewRotor(id string, wiring string, topLetter byte, notch byte) *Rotor {
	rotor := &Rotor{
		id:     id,
		wiring: wiring,
		notch:  notch,
	}

	// Create inverse wirings for "reverse" operations
	inverse := make([]byte, 26)
	for idx, letter := range rotor.wiring {
		newIdx := LetterToIdx(byte(letter))
		newLtr := IdxToLetter(byte(idx))
		inverse[newIdx] = newLtr
	}
	rotor.inverse = string(inverse)

	rotor.SetTopLetter(topLetter)

	return rotor
}

func GetRotor(id string) *Rotor {
	var rotor *Rotor

	switch id {
	case "0":
		// Identity Rotor
		rotor = NewRotor("0", ALPHABET, 'A', 'Z')
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
	// Rotor Wiring adjusted for position/offset
	adjWiring := rotor.wiring[rotor.position:] + rotor.wiring[0:rotor.position]

	return adjWiring
}

func (rotor *Rotor) InverseWiring() string {
	// Rotor Inverse Wiring adjusted for position/offset
	adjWiring := rotor.inverse[rotor.position:] + rotor.wiring[0:rotor.position]

	return adjWiring
}

func (rotor *Rotor) SetTopLetter(letter byte) {
	rotor.position = LetterToIdx(letter)
}

func (rotor *Rotor) GetTopLetter() byte {
	return IdxToLetter(rotor.position)
}

func (rotor *Rotor) Step() {
	rotor.position = (rotor.position + 1) % 26
}

func (rotor *Rotor) AtNotch() bool {
	atNotch := false
	if IdxToLetter(rotor.position) == rotor.notch {
		atNotch = true
	}

	return atNotch
}

// Right to Left
// ------
// IO   -   ABCDEFGHIJKLMNOPQRSTUVWXYZ
// III  -  BDFHJLCPRTXVZNYEIWGAKMUSQO
// II   -   AJDKSIRUXBLHWTMCQGZNPYFVOE
// ------
func (rotor *Rotor) Forward(letter byte) byte {
	idx := (LetterToIdx(letter) + rotor.position) % 26
	return rotor.wiring[idx] - rotor.position
}

// Left to Right
// --------------------
// func (rotor *Rotor) Reverse(letter byte) byte {
// 	lIdx := strings.IndexByte(rotor.wiring, letter)
// 	outPos := (byte(lIdx) + (26 - rotor.position)) % 26

// 	outLetter := IdxToLetter(outPos)

//		return outLetter
//	}
//
// --------------------
func (rotor *Rotor) Reverse(letter byte) byte {
	x := (LetterToIdx(letter) + rotor.position) % 26
	y := LetterToIdx(rotor.inverse[x])
	z := IdxToLetter((y - rotor.position) % 26)

	return z
}

// EOF
