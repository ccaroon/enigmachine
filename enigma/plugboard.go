package enigma

type Plugboard struct {
	connections []Cable
}

func NewPlugboard(connSpec []string) *Plugboard {
	// TODO: Error handling
	// * can't plug more than 1 cable into any one letter socket

	var numCables = len(connSpec)
	var plugboard Plugboard = Plugboard{
		connections: make([]Cable, 0, numCables),
	}

	for _, ltrPair := range connSpec {
		// if less than 2 letters, then ignore
		// if more than 2 letter, then ignore all after first two
		if len(ltrPair) >= 2 {
			cable := Cable{Plug1: rune(ltrPair[0]), Plug2: rune(ltrPair[1])}
			plugboard.connections = append(plugboard.connections, cable)
		}
	}

	return &plugboard
}

func (pb *Plugboard) NumCables() int {
	return len(pb.connections)
}

func (pb *Plugboard) GetCable(idx int) *Cable {
	var cable *Cable = nil

	if idx < len(pb.connections) {
		cable = &pb.connections[idx]
	}

	return cable
}

func (pb *Plugboard) FindCable(letter rune) *Cable {
	var foundCable *Cable

	for _, cable := range pb.connections {
		if cable.ConnectedTo(letter) {
			foundCable = &cable
			break
		}
	}

	return foundCable
}

func (pb *Plugboard) Map(letter rune) rune {
	var outLetter rune

	cable := pb.FindCable(letter)
	if cable != nil {
		outLetter = cable.Follow(letter)
	} else {
		outLetter = letter
	}

	return outLetter
}
