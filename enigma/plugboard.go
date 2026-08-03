package enigma

import "fmt"

type Plugboard struct {
	connections []Cable
}

func NewPlugboard(connSpec []string) (*Plugboard, error) {
	var numCables = len(connSpec)
	var plugboard Plugboard = Plugboard{
		connections: make([]Cable, 0, numCables),
	}

	for _, ltrPair := range connSpec {
		// if less than 2 letters, then ignore
		// if more than 2 letter, then ignore all after first two
		if len(ltrPair) >= 2 {
			ltr1 := rune(ltrPair[0])
			ltr2 := rune(ltrPair[1])

			// Check for duplicate plug settings
			for _, existingCable := range plugboard.connections {
				if existingCable.ConnectedTo(ltr1) || existingCable.ConnectedTo(ltr2) {
					return nil, fmt.Errorf("Duplicate Plugboard Setting: [%s] [%c%c]", ltrPair, existingCable.Plug1, existingCable.Plug2)
				}
			}

			cable := Cable{Plug1: ltr1, Plug2: ltr2}
			plugboard.connections = append(plugboard.connections, cable)
		}
	}

	return &plugboard, nil
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
