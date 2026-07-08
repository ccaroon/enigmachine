package enigma

type Cable struct {
	Plug1 byte
	Plug2 byte
}

func NewCable(ltr1, ltr2 byte) *Cable {
	return &Cable{
		Plug1: ltr1,
		Plug2: ltr2,
	}
}

func (cable *Cable) Follow(inLetter byte) byte {
	var outLetter byte

	if cable.Plug1 == inLetter {
		outLetter = cable.Plug2
	} else {
		outLetter = cable.Plug1
	}

	return outLetter
}

func (cable *Cable) ConnectedTo(letter byte) bool {
	var isConnected = false
	if cable.Plug1 == letter || cable.Plug2 == letter {
		isConnected = true
	}

	return isConnected
}
