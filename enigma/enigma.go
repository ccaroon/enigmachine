package enigma

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

func (enigma *Enigma) GetRotor(idx int) *Rotor {
	return enigma.rotors[idx]
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
	outLetter = enigma.reflector.Reflect(outLetter)

	// ### REVERSE (left-to-right) ###
	// // ROTORS
	// for idx := 0; idx < len(enigma.rotors); idx++ {
	// 	rotor := enigma.rotors[idx]
	// 	outLetter = rotor.Reverse(outLetter)
	// }
	// // PLUGBOARD
	// outLetter = enigma.plugboard.Map(outLetter)

	return outLetter
}
