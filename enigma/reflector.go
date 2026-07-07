package enigma

type Reflector struct {
	id     string
	wiring string
}

func NewReflector(id, wiring string) *Reflector {
	return &Reflector{
		id:     id,
		wiring: wiring,
	}
}

func GetReflector(id string) *Reflector {
	var reflector *Reflector

	switch id {
	case "B":
		reflector = NewReflector("B", "YRUHQSLDPXNGOKMIEBFZCWVJAT")
	case "C":
		reflector = NewReflector("C", "FVPJIAOYEDRZXWGCTKUQSBNMHL")
	}

	return reflector
}

func (ref *Reflector) Index(idx byte) byte {
	return ref.wiring[idx]
}

func (ref *Reflector) Reflect(letter byte) byte {
	inPos := LetterToIdx(letter)

	return ref.wiring[inPos]
}
