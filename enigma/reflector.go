package enigma

import "fmt"

type Reflector struct {
	id     string
	wiring string
}

func newReflector(id, wiring string) *Reflector {
	return &Reflector{
		id:     id,
		wiring: wiring,
	}
}

func GetReflector(id string) (*Reflector, error) {
	var reflector *Reflector

	switch id {
	case "B":
		reflector = newReflector("B", "YRUHQSLDPXNGOKMIEBFZCWVJAT")
	case "C":
		reflector = newReflector("C", "FVPJIAOYEDRZXWGCTKUQSBNMHL")
	default:
		return nil, fmt.Errorf("Unsupported Reflector: '%s'", id)
	}

	return reflector, nil
}

func (ref *Reflector) Id() string {
	return ref.id
}

func (ref *Reflector) Index(idx int) rune {
	return rune(ref.wiring[idx])
}

func (ref *Reflector) Reflect(letter rune) rune {
	inPos := LetterToIdx(letter)

	return rune(ref.wiring[inPos])
}
