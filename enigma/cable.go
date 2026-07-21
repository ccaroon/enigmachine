package enigma

type Cable struct {
	Plug1 rune
	Plug2 rune
}

func NewCable(ltr1, ltr2 rune) *Cable {
	return &Cable{
		Plug1: ltr1,
		Plug2: ltr2,
	}
}

func (cable *Cable) Follow(inLetter rune) rune {
	var outLetter rune

	if cable.Plug1 == inLetter {
		outLetter = cable.Plug2
	} else {
		outLetter = cable.Plug1
	}

	return outLetter
}

func (cable *Cable) ConnectedTo(letter rune) bool {
	var isConnected = false
	if cable.Plug1 == letter || cable.Plug2 == letter {
		isConnected = true
	}

	return isConnected
}
