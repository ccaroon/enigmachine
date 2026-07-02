package enigma

import "fmt"

type Rotor struct {
	Id       string
	Wiring   string
	Inverse  string
	position byte
}

var (
	presetRotors [6]Rotor = [6]Rotor{
		Rotor{
			Id:       "0",
			Wiring:   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			position: 0,
		},
		Rotor{
			Id:       "I",
			Wiring:   "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
			position: 0,
		},
		Rotor{
			Id:       "II",
			Wiring:   "AJDKSIRUXBLHWTMCQGZNPYFVOE",
			position: 0,
		},
		Rotor{
			Id:       "III",
			Wiring:   "BDFHJLCPRTXVZNYEIWGAKMUSQO",
			position: 0,
		},
		Rotor{
			Id:       "IV",
			Wiring:   "ESOVPZJAYQUIRHXLNFTGKDCMWB",
			position: 0,
		},
		Rotor{
			Id:       "V",
			Wiring:   "VZBRGITYUPSDNHLXAWMJQOFECK",
			position: 0,
		},
	}
)

func NewRotor(id string, wiring string) Rotor {
	rotor := Rotor{
		Id:     id,
		Wiring: wiring,
	}

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

	rotor.Initialize()

	return rotor
}

func (rotor *Rotor) Initialize() {
	// Get Inverse Wiring
	inverse := make([]byte, 26)
	for idx, letter := range rotor.Wiring {
		// fmt.Printf("%d -> %s\n", idx, string(letter))

		newIdx := LetterToIdx(byte(letter))
		newLtr := IdxToLetter(idx)
		// fmt.Printf("%d -> %s | %d -> %s\n", idx, string(letter), newIdx, string(newLtr))
		inverse[newIdx] = newLtr
	}

	// fmt.Println(inverse)

	rotor.Inverse = string(inverse)
}

func (rotor *Rotor) Debug() {
	fmt.Printf("W: %s\n", rotor.Wiring)
	fmt.Printf("I: %s\n", rotor.Inverse)
}
