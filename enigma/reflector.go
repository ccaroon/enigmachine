package enigma

import "fmt"

type Reflector struct {
	id     rune
	wiring string
}

func newReflector(id rune, wiring string) *Reflector {
	return &Reflector{
		id:     id,
		wiring: wiring,
	}
}

func GetReflector(id rune) (*Reflector, error) {
	var reflector *Reflector

	switch id {
	case 'A':
		reflector = newReflector('A', "EJMZALYXVBWFCRQUONTSPIKHGD")
	case 'B':
		reflector = newReflector('B', "YRUHQSLDPXNGOKMIEBFZCWVJAT")
	case 'C':
		reflector = newReflector('C', "FVPJIAOYEDRZXWGCTKUQSBNMHL")
	default:
		return nil, fmt.Errorf("Unsupported Reflector: '%c'", id)
	}

	return reflector, nil
}

func (ref *Reflector) Id() rune {
	return ref.id
}

func (ref *Reflector) Index(idx int) rune {
	return rune(ref.wiring[idx])
}

func (ref *Reflector) Reflect(letter rune) rune {
	inPos := LetterToIdx(letter)

	return rune(ref.wiring[inPos])
}
