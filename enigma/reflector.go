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
	case "I": // "I"dentity Reflector
		reflector = NewReflector("T", ALPHABET)
	case "B":
		reflector = NewReflector("B", "YRUHQSLDPXNGOKMIEBFZCWVJAT")
	case "C":
		reflector = NewReflector("C", "FVPJIAOYEDRZXWGCTKUQSBNMHL")
	}

	return reflector
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
