package enigma

type Plugboard struct {
	connections []Cable
}

func NewPlugboard(connSpec []rune) *Plugboard {
	var specLen = len(connSpec)

	// Should be an even number of letters
	// If not, then we have an unconnected, dangling cable...ignore it.
	if specLen%2 != 0 {
		specLen -= 1
		connSpec = connSpec[0:specLen]
	}

	// TODO: Error handling
	// * can't plug more than 1 cable into any one letter socket

	var numCables = specLen / 2
	var plugboard Plugboard = Plugboard{
		connections: make([]Cable, 0, numCables),
	}

	for i := 0; i < specLen; i += 2 {
		cable := Cable{Plug1: connSpec[i], Plug2: connSpec[i+1]}
		plugboard.connections = append(plugboard.connections, cable)
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
